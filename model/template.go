package model

import (
	"errors"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/model"
	commonmodel "github.com/CloudSilk/pkg/model"
	"github.com/jinzhu/copier"
)

type Template struct {
	model.Model
	TenantID    string `gorm:"size:36;index"`
	Name        string `json:"name" gorm:"size:100;index"`
	Language    string `json:"language" gorm:"size:50;index"`
	Content     string `json:"content"`
	Description string `json:"description" gorm:"size:200;"`
	Group       string `json:"group" gorm:"size:50;index"`
}

func CreateTemplate(tpl *Template) error {
	duplication, err := dbClient.CreateWithCheckDuplication(tpl, "name = ? and `group`=? and `language`=? and tenant_id = ?", tpl.Name, tpl.Group, tpl.Language, tpl.TenantID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同代码模板")
	}
	return nil
}

func DeleteTemplate(id string) (err error) {
	return dbClient.DB().Unscoped().Delete(&Template{}, "id=?", id).Error
}

func QueryTemplate(req *apipb.QueryTemplateRequest, resp *apipb.QueryTemplateResponse) {
	db := dbClient.DB().Model(&Template{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		db = db.Where("language = ?", req.Language)
	}

	if req.Group != "" {
		db = db.Where("`group` = ?", req.Group)
	}

	if req.TenantID != "" {
		db = db.Where("`tenant_id` = ?", req.TenantID)
	}

	if req.Id != "" {
		db = db.Where("id LIKE ?", "%"+req.Id+"%")
	}

	OrderStr := "language,`group`,name"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}
	var err error
	var result []*Template
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &result, nil)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = TemplatesToPB(result)
	}
	resp.Total = resp.Records
}

func GetAllTemplates(req *apipb.QueryTemplateRequest) (tpls []*Template, err error) {
	db := dbClient.DB()
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}
	err = db.Order("language,`group`,name").Find(&tpls).Error
	return
}

func GetTemplateById(id string) (tpl Template, err error) {
	err = dbClient.DB().Where("id = ?", id).First(&tpl).Error
	return
}

func UpdateTemplate(tpl *Template) error {
	duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(dbClient.DB(), tpl, false, []string{"created_at", "tenant_id"}, "id <> ? and name = ? and `group`=? and `language`=? and tenant_id = ?", tpl.ID, tpl.Name, tpl.Group, tpl.Language, tpl.TenantID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同代码模板")
	}

	return nil
}

func CopyTemplate(id string) error {
	from, err := GetTemplateById(id)
	if err != nil {
		return err
	}
	to := &Template{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreateTemplate(to)
}

func PBToTemplates(in []*apipb.TemplateInfo) []*Template {
	var result []*Template
	for _, c := range in {
		result = append(result, PBToTemplate(c))
	}
	return result
}

func PBToTemplate(in *apipb.TemplateInfo) *Template {
	return &Template{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		TenantID:    in.TenantID,
		Name:        in.Name,
		Language:    in.Language,
		Content:     in.Content,
		Description: in.Description,
		Group:       in.Group,
	}
}

func TemplatesToPB(in []*Template) []*apipb.TemplateInfo {
	var list []*apipb.TemplateInfo
	for _, f := range in {
		list = append(list, TemplateToPB(f))
	}
	return list
}

func TemplateToPB(in *Template) *apipb.TemplateInfo {
	return &apipb.TemplateInfo{
		Id:          in.ID,
		TenantID:    in.TenantID,
		Name:        in.Name,
		Language:    in.Language,
		Content:     in.Content,
		Description: in.Description,
		Group:       in.Group,
	}
}
