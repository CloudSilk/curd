package model

import (
	"errors"
	"fmt"
	"time"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/curd/service"
	"github.com/CloudSilk/pkg/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Form struct {
	model.TenantModel
	ProjectID string         `gorm:"index;size:36"`
	Name      string         `json:"name" gorm:"index;size:200"`
	PageName  string         `json:"pageName" gorm:"index;size:200"`
	Group     string         `json:"group" gorm:"index;size:200"`
	Schema    string         `json:"schema"`
	Type      string         `json:"type" gorm:"index;size:100;comment:例如Cell、CURD、AIoT、Page、BPM"`
	Public    bool           `gorm:"comment:是否是公共表单"`
	Subform   bool           `gorm:"index;子表单"`
	Versions  []*FormVersion `json:"versions" copier:"-"`
	IsMust    bool           `json:"isMust" gorm:"index;comment:系统必须要有的数据"`
}

type FormVersion struct {
	model.Model
	Version     string `json:"version" gorm:"size:50;index"`
	FormID      string `json:"formID" gorm:"index"`
	Schema      string `json:"schema"`
	Description string `json:"description" gorm:"size:200"`
}

func CreateForm(m *Form) error {
	count, err := statisticFormCount(dbClient.DB(), m.TenantID, m.ProjectID)
	if err != nil {
		return err
	}

	expired, projectFormCount, err := service.GetProjectFormCount(m.ProjectID)
	if err != nil {
		return err
	}
	if expired {
		return errors.New("账号使用期限已过，你可以联系管理员!")
	}

	if projectFormCount > 0 && projectFormCount <= int32(count) {
		return fmt.Errorf("只能创建 %d 个表单", projectFormCount)
	}
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {

		duplication, err := dbClient.CreateWithCheckDuplicationWithDB(tx, m, "name = ? and page_name = ? and project_id=? and tenant_id = ?", m.Name, m.PageName, m.ProjectID, m.TenantID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同表单")
		}
		return nil
	})
}

func PublishForm(req *FormVersion) error {
	if req.Version == "" {
		req.Version = time.Now().Format("20060102150405")
	}
	duplication, err := dbClient.CreateWithCheckDuplication(req, "version = ? and form_id=?", req.Version, req.FormID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同表单版本号")
	}
	return nil
}

func SwitchFormVersion(versionID string) error {
	version := &FormVersion{}
	err := dbClient.DB().Where("id = ?", versionID).First(&version).Error
	if err != nil {
		return err
	}
	return dbClient.DB().Model(&Form{}).Where("id=?", version.FormID).Update("schema", version.Schema).Error
}

func DeleteForm(id string) (err error) {
	return dbClient.DB().Delete(&Form{}, "id=?", id).Error
}

func QueryForm(req *apipb.QueryFormRequest, resp *apipb.QueryFormResponse, preload bool) {
	db := dbClient.DB().Model(&Form{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.PageName != "" {
		db = db.Where("page_name LIKE ?", "%"+req.PageName+"%")
	}

	if req.Group != "" {
		db = db.Where("`group` LIKE ?", "%"+req.Group+"%")
	}
	if len(req.Ids) > 0 {
		db = db.Where("id in ?", req.Ids)
	}
	if req.Subform > 0 {
		db = db.Where("subform=?", req.Subform == 1)
	}
	if req.Public > 0 {
		db = db.Where("`public`=?", req.Public == 1)
	}
	if req.Type != "" {
		db = db.Where("`type`=?", req.Type)
	}
	if req.TenantID != "" {
		db = db.Where("`tenant_id`=?", req.TenantID)
	}
	if !preload {
		db = db.Omit("schema")
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}

	if req.InclusiveBasic {
		db = db.Or("is_must = ?", req.InclusiveBasic)
	}

	OrderStr := "`updated_at` desc"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}

	var err error
	var forms []Form
	if preload {
		resp.Records, resp.Pages, err = dbClient.PageQueryWithPreload(db, req.PageSize, req.PageIndex, OrderStr, []string{"Versions"}, &forms)
	} else {
		resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &forms, nil)
	}
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = FormsToPB(forms)
		resp.Total = resp.Records
	}
}

func GetAllForms(req *apipb.GetAllFormRequest) (fs []Form, err error) {
	db := dbClient.DB()
	if req.Subform > 0 {
		db = db.Where("subform=?", req.Subform == 1)
	}
	if req.Public > 0 {
		db = db.Where("`public`=?", req.Public == 1)
	}
	if req.Type != "" {
		db = db.Where("`type`=?", req.Type)
	}
	if req.TenantID != "" {
		db = db.Where("`tenant_id`=?", req.TenantID)
	}
	if req.ProjectID != "" {
		db = db.Where("`project_id`=?", req.ProjectID)
	}

	err = db.Find(&fs).Error
	return
}

func GetFormById(id string, containerVersions bool) (f Form, err error) {
	db := dbClient.DB()
	if containerVersions {
		db = db.Preload("Versions")
	}
	err = db.Where("id = ?", id).First(&f).Error
	return f, err
}

func GetFormVersionById(id string) (version FormVersion, err error) {
	err = dbClient.DB().Where("id = ?", id).First(&version).Error
	return
}

func UpdateForm(f *Form) error {
	duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(dbClient.DB(), f, false, []string{"schema", "Versions", "created_at"}, "id <> ? and name = ? and page_name = ? and project_id=? and tenant_id = ?", f.ID, f.Name, f.PageName, f.ProjectID, f.TenantID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同表单")
	}

	return nil
}

func UpdateFormAll(f *Form) error {
	duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(dbClient.DB(), f, true, []string{"created_at"}, "id <> ? and name = ? and page_name = ? and project_id=?", f.ID, f.Name, f.PageName, f.ProjectID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同表单")
	}

	return nil
}

func UpdateFormSchema(f *Form) error {
	return dbClient.DB().Model(f).Where("id", f.ID).Update("schema", f.Schema).Error
}

func CopyForm(id string) error {
	from, err := GetFormById(id, false)
	if err != nil {
		return err
	}
	to := &Form{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreateForm(to)
}

func statisticFormCount(db *gorm.DB, tenantID, projectID string) (int64, error) {
	db = db.Model(&Form{})
	if tenantID != "" {
		db = db.Where("tenant_id = ?", tenantID)
	}
	if projectID != "" {
		db = db.Where("project_id = ?", projectID)
	}
	var count int64
	err := db.Count(&count).Error
	return count, err
}
