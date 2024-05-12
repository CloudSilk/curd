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

func CreateFileTemplate(m *FileTemplate) (string, error) {
	duplication, err := dbClient.CreateWithCheckDuplication(m, "`language`=? and `group`=? and name =? ", m.Language, m.Group, m.Name)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同文件模版")
	}
	return m.ID, nil
}

//删除子表

func UpdateFileTemplate(m *FileTemplate) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldFileTemplate := &FileTemplate{}
		err := tx.Preload(clause.Associations).Where("id = ?", m.ID).First(oldFileTemplate).Error
		if err != nil {
			return err
		}
		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  and `language`=? and `group`=? and name =? ", m.ID, m.Language, m.Group, m.Name)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同文件模版")
		}

		return nil
	})
}

func DeleteFileTemplate(id string) (err error) {
	return dbClient.DB().Delete(&FileTemplate{}, "id=?", id).Error
}

func QueryFileTemplate(req *apipb.QueryFileTemplateRequest, resp *apipb.QueryFileTemplateResponse, preload bool) {
	db := dbClient.DB().Model(&FileTemplate{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Language != "" {
		db = db.Where("language = ?", req.Language)
	}

	if req.Group != "" {
		db = db.Where("`group` = ?", req.Group)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "`language`,`group`,`name`")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*FileTemplate
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = FileTemplatesToPB(list)
	}
	resp.Total = resp.Records
}

func GetFileTemplateByID(id string) (*FileTemplate, error) {
	m := &FileTemplate{}
	err := dbClient.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetFileTemplateByIDs(ids []string) ([]*FileTemplate, error) {
	var m []*FileTemplate
	err := dbClient.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func GetAllFileTemplates() (list []*FileTemplate, err error) {
	err = dbClient.DB().Order("`language`,`group`,`name`").Find(&list).Error
	return
}

func CopyFileTemplate(id string) (string, error) {
	from, err := GetFileTemplateByID(id)
	if err != nil {
		return "", err
	}
	to := &FileTemplate{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return "", err
	}
	to.Name += " Copy"
	return CreateFileTemplate(to)
}

type FileTemplate struct {
	commonmodel.Model
	//文件模版名称
	Name string `json:"name" gorm:"size:100;index;uniqueindex:FileTemplate_uidx1" `
	//文件存放的目录
	Dir string `json:"dir" gorm:"size:100" `
	//包名
	Package string `json:"package" gorm:"size:100" `
	//开头模板ID
	Start string `json:"start" gorm:"size:36"`
	//结尾模板ID
	End string `json:"end" gorm:"size:36"`
	//一组模板ID
	Body string `json:"body" gorm:"size:500"`
	//自定义参数
	Params   string `json:"params" gorm:"size:500" `
	Language string `json:"language" gorm:"size:50;index;uniqueindex:FileTemplate_uidx1"`
	Group    string `json:"group" gorm:"size:50;index;uniqueindex:FileTemplate_uidx1"`
	//文件名称后缀
	FileNameSuffix string `json:"fileNameSuffix" gorm:"size:100"`
}

func PBToFileTemplates(in []*apipb.FileTemplateInfo) []*FileTemplate {
	var result []*FileTemplate
	for _, c := range in {
		result = append(result, PBToFileTemplate(c))
	}
	return result
}

func PBToFileTemplate(in *apipb.FileTemplateInfo) *FileTemplate {
	if in == nil {
		return nil
	}
	return &FileTemplate{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Name:           in.Name,
		Dir:            in.Dir,
		Package:        in.Package,
		Start:          in.Start,
		End:            in.End,
		Body:           ObjectToJsonString(in.Body),
		Params:         in.Params,
		Language:       in.Language,
		Group:          in.Group,
		FileNameSuffix: in.FileNameSuffix,
	}
}

func FileTemplatesToPB(in []*FileTemplate) []*apipb.FileTemplateInfo {
	var list []*apipb.FileTemplateInfo
	for _, f := range in {
		list = append(list, FileTemplateToPB(f))
	}
	return list
}

func FileTemplateToPB(in *FileTemplate) *apipb.FileTemplateInfo {
	if in == nil {
		return nil
	}
	return &apipb.FileTemplateInfo{
		Id:             in.ID,
		Name:           in.Name,
		Dir:            in.Dir,
		Package:        in.Package,
		Start:          in.Start,
		End:            in.End,
		Body:           JsonToStringArray(in.Body),
		Params:         in.Params,
		Language:       in.Language,
		Group:          in.Group,
		FileNameSuffix: in.FileNameSuffix,
	}
}
