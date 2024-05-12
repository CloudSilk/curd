package model

import (
	"encoding/json"
	"errors"
	"sort"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/model"
	commonmodel "github.com/CloudSilk/pkg/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	model.Model
	TenantID           string      `gorm:"size:36;index"`
	ProjectID          string      `gorm:"index;size:36"`
	Name               string      `json:"name" gorm:"size:100;comment:服务名称" `
	DisplayName        string      `json:"displayName" gorm:"size:100;comment:服务显示名称" `
	Package            string      `json:"package" gorm:"size:50"`
	Params             string      `json:"params" gorm:"size:500;comment:自定义参数"`
	CodeFiles          []*CodeFile `json:"codeFiles"`
	ServiceFunctionals []*ServiceFunctional
}

func (s *Service) Sort() {
	sort.Slice(s.CodeFiles, func(i, j int) bool {
		return s.CodeFiles[i].Dir < s.CodeFiles[j].Dir
	})
}

type ServiceFunctional struct {
	commonmodel.Model
	//服务ID
	ServiceID string `json:"serviceID" gorm:"size:36;index" `
	//功能模版ID
	FunctionalTemplateID string `json:"functionalTemplateID" gorm:"size:36;index" `
	//元数据ID
	MetadataID string `json:"metadataID" gorm:"size:36;index" `
	//自定义参数
	Params string `json:"params" gorm:"size:500" `
	//生成配置文件
	GenConfig bool `json:"genConfig"`
	Enable    bool `json:"enable"`
}

type CodeFile struct {
	model.Model
	ServiceID       string `json:"serviceID" gorm:"index" copier:"-"`
	Name            string `json:"name" gorm:"size:100;comment:文件名称"`
	Dir             string `json:"dir" gorm:"size:100;comment:文件存放的目录" `
	Package         string `json:"package" gorm:"size:50;comment:包名"`
	MetadataID      string `json:"metadataID" gorm:"size:36;comment:元数据ID"`
	Start           string `json:"start" gorm:"size:36;comment:开头模板ID"`
	End             string `json:"end" gorm:"size:36;comment:结尾模板ID"`
	Body            string `json:"body" gorm:"size:500;comment:一组模板ID"`
	Params          string `json:"params" gorm:"size:500;comment:自定义参数"`
	Code            string `json:"code" gorm:"-"`
	TemplateContent string `json:"templateContent" gorm:"-"`
	Enable          bool   `json:"enable"`
}

func CreateService(m *Service) error {
	duplication, err := dbClient.CreateWithCheckDuplication(m, " name=? and tenant_id =? and project_id = ?", m.Name, m.TenantID, m.ProjectID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同Service")
	}
	return nil
}

func AddCodeFiles(codeFiles []*CodeFile) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		for _, codeFile := range codeFiles {
			err := tx.Omit("created_at").Save(codeFile).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func QueryService(req *apipb.QueryServiceRequest, resp *apipb.QueryServiceResponse, preload bool) {
	db := dbClient.DB().Model(&Service{})
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}

	OrderStr := "`name`"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}
	var err error
	var list []*Service
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = ServicesToPB(list)
	}
	resp.Total = resp.Records
}

func DeleteService(id string) (err error) {
	return dbClient.DB().Unscoped().Delete(&Service{}, "id=?", id).Error
}

func GetServiceByID(id string, all bool) (*Service, error) {
	m := &Service{}
	db := dbClient.DB().Preload(clause.Associations)
	db = db.Preload("CodeFiles")
	err := db.Where("id = ?", id).First(m).Error
	return m, err
}

func CopyService(id string) error {
	from, err := GetServiceByID(id, false)
	if err != nil {
		return err
	}
	to := &Service{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreateService(to)
}

func GetAllService(req *apipb.QueryServiceRequest) (list []*Service, err error) {
	db := dbClient.DB()
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}
	err = db.Find(&list).Error
	return
}

func DeleteCodeFiles(tx *gorm.DB, old, m *Service) error {
	var deleteIDs []string
	for _, oldObj := range old.CodeFiles {
		flag := false
		for _, newObj := range m.CodeFiles {
			if newObj.ID == oldObj.ID {
				flag = true

			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&CodeFile{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteServiceFunctionals(tx *gorm.DB, old, m *Service) error {
	var deleteIDs []string
	for _, oldObj := range old.ServiceFunctionals {
		flag := false
		for _, newObj := range m.ServiceFunctionals {
			if newObj.ID == oldObj.ID {
				flag = true

			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&ServiceFunctional{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateService(m *Service) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldService := &Service{}
		err := tx.Preload(clause.Associations).Where("id = ?", m.ID).First(oldService).Error
		if err != nil {
			return err
		}
		err = DeleteCodeFiles(tx, oldService, m)
		if err != nil {
			return err
		}

		err = DeleteServiceFunctionals(tx, oldService, m)
		if err != nil {
			return err
		}

		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  and  name=? and tenant_id =? and project_id = ?", m.ID, m.Name, m.TenantID, m.ProjectID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同Service")
		}

		return nil
	})
}

func PBToServices(in []*apipb.ServiceInfo) []*Service {
	var result []*Service
	for _, c := range in {
		result = append(result, PBToService(c))
	}
	return result
}

func PBToService(in *apipb.ServiceInfo) *Service {
	return &Service{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Name:               in.Name,
		CodeFiles:          PBToCodeFiles(in.CodeFiles),
		DisplayName:        in.DisplayName,
		Package:            in.Package,
		Params:             in.Params,
		TenantID:           in.TenantID,
		ProjectID:          in.ProjectID,
		ServiceFunctionals: PBToServiceFunctionals(in.ServiceFunctionals),
	}
}

func ServicesToPB(in []*Service) []*apipb.ServiceInfo {
	var list []*apipb.ServiceInfo
	for _, f := range in {
		list = append(list, ServiceToPB(f))
	}
	return list
}

func ServiceToPB(in *Service) *apipb.ServiceInfo {
	return &apipb.ServiceInfo{
		Id:                 in.ID,
		Name:               in.Name,
		CodeFiles:          CodeFilesToPB(in.CodeFiles),
		DisplayName:        in.DisplayName,
		Package:            in.Package,
		Params:             in.Params,
		TenantID:           in.TenantID,
		ProjectID:          in.ProjectID,
		ServiceFunctionals: ServiceFunctionalsToPB(in.ServiceFunctionals),
	}
}

func PBToCodeFiles(in []*apipb.CodeFileInfo) []*CodeFile {
	var result []*CodeFile
	for _, c := range in {
		result = append(result, PBToCodeFile(c))
	}
	return result
}

func PBToCodeFile(in *apipb.CodeFileInfo) *CodeFile {
	return &CodeFile{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		ServiceID:  in.ServiceID,
		Name:       in.Name,
		Dir:        in.Dir,
		Start:      in.Start,
		End:        in.End,
		Body:       ObjectToJsonString(in.Body),
		Package:    in.Package,
		MetadataID: in.MetadataID,
		Params:     in.Params,
		Enable:     in.Enable,
	}
}

func CodeFilesToPB(in []*CodeFile) []*apipb.CodeFileInfo {
	var list []*apipb.CodeFileInfo
	for _, f := range in {
		list = append(list, CodeFileToPB(f))
	}
	return list
}

func CodeFileToPB(in *CodeFile) *apipb.CodeFileInfo {
	return &apipb.CodeFileInfo{
		Id:              in.ID,
		ServiceID:       in.ServiceID,
		Name:            in.Name,
		Dir:             in.Dir,
		Start:           in.Start,
		End:             in.End,
		Body:            JsonToStringArray(in.Body),
		Package:         in.Package,
		MetadataID:      in.MetadataID,
		Params:          in.Params,
		Code:            in.Code,
		TemplateContent: in.TemplateContent,
		Enable:          in.Enable,
	}
}

func PBToServiceFunctionals(in []*apipb.ServiceFunctionalInfo) []*ServiceFunctional {
	var result []*ServiceFunctional
	for _, c := range in {
		result = append(result, PBToServiceFunctional(c))
	}
	return result
}

func PBToServiceFunctional(in *apipb.ServiceFunctionalInfo) *ServiceFunctional {
	if in == nil {
		return nil
	}
	return &ServiceFunctional{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		ServiceID:            in.ServiceID,
		FunctionalTemplateID: in.FunctionalTemplateID,
		MetadataID:           in.MetadataID,
		Params:               in.Params,
		GenConfig:            in.GenConfig,
		Enable:               in.Enable,
	}
}

func ServiceFunctionalsToPB(in []*ServiceFunctional) []*apipb.ServiceFunctionalInfo {
	var list []*apipb.ServiceFunctionalInfo
	for _, f := range in {
		list = append(list, ServiceFunctionalToPB(f))
	}
	return list
}

func ServiceFunctionalToPB(in *ServiceFunctional) *apipb.ServiceFunctionalInfo {
	if in == nil {
		return nil
	}
	return &apipb.ServiceFunctionalInfo{
		Id:                   in.ID,
		ServiceID:            in.ServiceID,
		FunctionalTemplateID: in.FunctionalTemplateID,
		MetadataID:           in.MetadataID,
		Params:               in.Params,
		GenConfig:            in.GenConfig,
		Enable:               in.Enable,
	}
}

func ObjectToJsonString(obj interface{}) string {
	buf, _ := json.Marshal(obj)
	return string(buf)
}

func JsonToStringArray(str string) []string {
	var result []string
	json.Unmarshal([]byte(str), &result)
	return result
}
