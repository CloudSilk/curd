package model

import (
	apipb "github.com/CloudSilk/curd/proto"
	commonmodel "github.com/CloudSilk/pkg/model"
)

func PBToSystemObjects(in []*apipb.SystemObjectInfo) []*SystemObject {
	var result []*SystemObject
	for _, c := range in {
		result = append(result, PBToSystemObject(c))
	}
	return result
}

func PBToSystemObject(in *apipb.SystemObjectInfo) *SystemObject {
	if in == nil {
		return nil
	}
	return &SystemObject{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Enable:      in.Enable,
		Name:        in.Name,
		DisplayName: in.DisplayName,
		Language:    in.Language,
		Type:        in.Type,
		Code:        in.Code,
		Description: in.Description,
	}
}

func SystemObjectsToPB(in []*SystemObject) []*apipb.SystemObjectInfo {
	var list []*apipb.SystemObjectInfo
	for _, f := range in {
		list = append(list, SystemObjectToPB(f))
	}
	return list
}

func SystemObjectToPB(in *SystemObject) *apipb.SystemObjectInfo {
	if in == nil {
		return nil
	}
	return &apipb.SystemObjectInfo{
		Id:          in.ID,
		Enable:      in.Enable,
		Name:        in.Name,
		DisplayName: in.DisplayName,
		Language:    in.Language,
		Type:        in.Type,
		Code:        in.Code,
		Description: in.Description,
	}
}

type SystemObject struct {
	commonmodel.Model
	//是否启用
	Enable bool `json:"enable" gorm:"index" `
	//名称
	Name string `json:"name" gorm:"size:200;index" `
	//显示名称
	DisplayName string `json:"displayName" gorm:"size:100" `
	//编程语言
	Language string `json:"language" gorm:"size:50;index" `
	//类型 1-类 2-结构体 3-方法 4-代码
	Type int32 `json:"type" gorm:"index;comment:1-类 2-结构体 3-方法 4-代码" `
	//代码
	Code string `json:"code" gorm:"size:2000" `
	//备注
	Description string `json:"description" gorm:"size:500" `
}
