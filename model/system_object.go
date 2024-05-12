package model

import (
	"errors"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateSystemObject(m *SystemObject) (string, error) {
	duplication, err := dbClient.CreateWithCheckDuplication(m, " name=? ", m.Name)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同系统定义的对象")
	}
	return m.ID, nil
}

func UpdateSystemObject(m *SystemObject) error {
	duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(dbClient.DB(), m, false, []string{"created_at"}, "id != ? and  name=? ", m.ID, m.Name)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同系统定义的对象")
	}

	return nil
}

func QuerySystemObject(req *apipb.QuerySystemObjectRequest, resp *apipb.QuerySystemObjectResponse, preload bool) {
	db := dbClient.DB().Model(&SystemObject{})
	if req.Enable > 0 {
		db = db.Where("enable = ?", req.Enable == 1)
	}

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		db = db.Where("language = ?", req.Language)
	}

	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "`name`")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*SystemObject
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = SystemObjectsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSystemObjects(req *apipb.GetAllSystemObjectRequest) (list []*SystemObject, err error) {
	db := dbClient.DB().Model(&SystemObject{})
	db = db.Where("enable = ?", req.Enable)
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		db = db.Where("language = ?", req.Language)
	}

	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}
	err = db.Find(&list).Error
	return
}

func GetSystemObjectByID(id string) (*SystemObject, error) {
	m := &SystemObject{}
	err := dbClient.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetSystemObjectByIDs(ids []string) ([]*SystemObject, error) {
	var m []*SystemObject
	err := dbClient.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

//删除子表

func DeleteSystemObject(id string) (err error) {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Unscoped().Delete(&SystemObject{}, "id=?", id).Error
	})
}

func CopySystemObject(id string) (string, error) {
	from, err := GetSystemObjectByID(id)
	if err != nil {
		return "", err
	}
	to := &SystemObject{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return "", err
	}
	to.Name += " Copy"
	return CreateSystemObject(to)
}

func EnableSystemObject(id string, enable bool) error {
	err := dbClient.DB().Model(&SystemObject{}).Where("id=?", id).Update("enable", enable).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateSystemObjectAll(m *SystemObject) error {
	duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(dbClient.DB(), m, true, []string{"created_at"}, "id != ? and  name=? ", m.ID, m.Name)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同系统定义的对象")
	}

	return nil
}
