package model

import (
	"errors"

	apipb "github.com/CloudSilk/curd/proto"
	commonmodel "github.com/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateFunctionalTemplate(m *FunctionalTemplate) (string, error) {
	duplication, err := dbClient.CreateWithCheckDuplication(m, " name=? ", m.Name)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同功能模版")
	}
	return m.ID, nil
}

//删除子表

func UpdateFunctionalTemplate(m *FunctionalTemplate) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldFunctionalTemplate := &FunctionalTemplate{}
		err := tx.Preload(clause.Associations).Where("id = ?", m.ID).First(oldFunctionalTemplate).Error
		if err != nil {
			return err
		}
		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  and  name=? ", m.ID, m.Name)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同功能模版")
		}

		return nil
	})
}

func DeleteFunctionalTemplate(id string) (err error) {
	return dbClient.DB().Delete(&FunctionalTemplate{}, "id=?", id).Error
}

func QueryFunctionalTemplate(req *apipb.QueryFunctionalTemplateRequest, resp *apipb.QueryFunctionalTemplateResponse, preload bool) {
	db := dbClient.DB().Model(&FunctionalTemplate{})
	if req.Language != "" {
		db = db.Where("language = ?", req.Language)
	}

	if req.Group != "" {
		db = db.Where("group = ?", req.Group)
	}

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "`name`")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*FunctionalTemplate
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = FunctionalTemplatesToPB(list)
	}
	resp.Total = resp.Records
}

func GetFunctionalTemplateByID(id string) (*FunctionalTemplate, error) {
	m := &FunctionalTemplate{}
	err := dbClient.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetFunctionalTemplateByIDs(ids []string) ([]*FunctionalTemplate, error) {
	var m []*FunctionalTemplate
	err := dbClient.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func GetAllFunctionalTemplates() (list []*FunctionalTemplate, err error) {
	err = dbClient.DB().Find(&list).Error
	return
}

func CopyFunctionalTemplate(id string) (string, error) {
	from, err := GetFunctionalTemplateByID(id)
	if err != nil {
		return "", err
	}
	to := &FunctionalTemplate{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return "", err
	}
	to.Name += " Copy"
	return CreateFunctionalTemplate(to)
}

func EnableFunctionalTemplate(id string, enable bool) error {
	err := dbClient.DB().Model(&FunctionalTemplate{}).Where("id=?", id).Update("enable", enable).Error
	if err != nil {
		return err
	}
	return nil
}

type FunctionalTemplate struct {
	commonmodel.Model
	//描述
	Description string `json:"description" gorm:"size:500" `
	//自定义参数
	Params string `json:"params" gorm:"size:500" `
	//语音
	Language string `json:"language" gorm:"size:50;index" `
	//分组
	Group string `json:"group" gorm:"size:50;index" `
	//名称
	Name string `json:"name" gorm:"size:100;index" `
	//文件模版ID
	FileTemplateIDs string `json:"fileTemplateIDs" gorm:"size:500" `
}

func PBToFunctionalTemplates(in []*apipb.FunctionalTemplateInfo) []*FunctionalTemplate {
	var result []*FunctionalTemplate
	for _, c := range in {
		result = append(result, PBToFunctionalTemplate(c))
	}
	return result
}

func PBToFunctionalTemplate(in *apipb.FunctionalTemplateInfo) *FunctionalTemplate {
	if in == nil {
		return nil
	}
	return &FunctionalTemplate{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Description:     in.Description,
		Params:          in.Params,
		Language:        in.Language,
		Group:           in.Group,
		Name:            in.Name,
		FileTemplateIDs: ObjectToJsonString(in.FileTemplateIDs),
	}
}

func FunctionalTemplatesToPB(in []*FunctionalTemplate) []*apipb.FunctionalTemplateInfo {
	var list []*apipb.FunctionalTemplateInfo
	for _, f := range in {
		list = append(list, FunctionalTemplateToPB(f))
	}
	return list
}

func FunctionalTemplateToPB(in *FunctionalTemplate) *apipb.FunctionalTemplateInfo {
	if in == nil {
		return nil
	}
	return &apipb.FunctionalTemplateInfo{
		Id:              in.ID,
		Description:     in.Description,
		Params:          in.Params,
		Language:        in.Language,
		Group:           in.Group,
		Name:            in.Name,
		FileTemplateIDs: JsonToStringArray(in.FileTemplateIDs),
	}
}
