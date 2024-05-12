package model

import (
	"errors"
	"sort"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Metadata struct {
	model.Model
	TenantID       string           `gorm:"size:36;index"`
	ProjectID      string           `gorm:"size:36;index"`
	Name           string           `json:"name" gorm:"size:100;index"`
	DisplayName    string           `json:"displayName" gorm:"size:100;comment:显示名字"`
	Level          int32            `json:"level"`
	ParentID       string           `json:"parentID"`
	Description    string           `json:"description" gorm:"size:200;"`
	Package        string           `json:"package" gorm:"size:100"`
	MetadataFields []*MetadataField `json:"metadataFields"`
	Children       []*Metadata      `json:"children" gorm:"-"`
	UniqueFields   string           `json:"-" gorm:"-"`
	Fields         string           `json:"-" gorm:"-"`
	Preloads       []string         `json:"-" gorm:"-"`
	System         string           `json:"system" gorm:"index;size:100"`
	IsMust         bool             `json:"isMust" gorm:"index;comment:系统必须要有的数据"`
}

func (md *Metadata) Sort() {
	sort.Slice(md.MetadataFields, func(i, j int) bool {
		return md.MetadataFields[i].Order < md.MetadataFields[j].Order
	})
}

func (md *Metadata) FieldSort() {
	for i, field := range md.MetadataFields {
		if field.Order == 0 {
			field.Order = int32(i + 1)
		}
	}
}

type MetadataField struct {
	model.Model
	MetadataID   string `json:"metadataID"`
	ProjectID    string `gorm:"index;size:36"`
	Name         string `json:"name" gorm:"size:100;"`
	Type         string `json:"type" gorm:"size:50;comment:基础数据类型"`
	Length       int32  `json:"length" gorm:"comment:字段长度"`
	NotNull      bool   `json:"notNull" gorm:"comment:是否可以为空"`
	Comment      string `json:"comment" gorm:"size:200;"`
	IsArray      bool   `json:"isArray" gorm:"comment:数组"`
	RefMetadata  string `json:"refMetadata" gorm:"size:100;comment:引用其他元数据"`
	DisplayName  string `json:"displayName" gorm:"size:100;comment:显示名字"`
	ShowInTable  bool   `json:"showInTable" gorm:"comment:是否在列表中显示"`
	ShowInEdit   bool   `json:"showInEdit" gorm:"comment:是否在编辑中显示"`
	Component    string `json:"component" gorm:"size:100;comment:显示组件类型"`
	Unique       bool   `json:"unique" gorm:"comment:唯一索引"`
	Index        bool   `json:"index" gorm:"comment:索引"`
	DefaultValue string `json:"defaultValue" gorm:"size:100;comment:默认值"`
	ShowInQuery  bool   `json:"showInQuery" gorm:"comment:查询条件"`
	Order        int32  `json:"order" gorm:"comment:显示顺序"`
	Like         bool   `json:"like" gorm:"comment:like查询"`
	Copier       bool   `json:"copier" gorm:"comment:是否不拷贝"`
	DotNotGen    bool
	PBToStruct   string `json:"pbToStruct" gorm:"size:100;"`
	StructToPB   string `json:"structToPB" gorm:"size:100;"`
}

func CreateMetadata(md *Metadata) error {
	if md.ParentID != "" {
		parent, err := GetMetadataById(md.ParentID)
		if err != nil {
			return err
		}
		md.Level = parent.Level + 1
	}
	md.FieldSort()
	duplication, err := dbClient.CreateWithCheckDuplication(md, "`system`=? and name = ? and project_id=? and tenant_id=?", md.System, md.Name, md.ProjectID, md.TenantID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同元数据")
	}
	return nil
}

func DeleteMetadata(id string) (err error) {
	return dbClient.DB().Delete(&Metadata{}, "id=?", id).Error
}

func QueryMetadata(req *apipb.QueryMetadataRequest, resp *apipb.QueryMetadataResponse, preload bool) {
	db := dbClient.DB().Model(&Metadata{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.ParentID != "1000000000" && req.ParentID != "" {
		db = db.Where("parent_id=?", req.ParentID)
	}

	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}
	if len(req.Ids) > 0 {
		db = db.Where("id in ?", req.Ids)
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
	var metadatas []*Metadata
	if preload {
		resp.Records, resp.Pages, err = dbClient.PageQueryWithPreload(db, req.PageSize, req.PageIndex, OrderStr, []string{"MetadataFields", clause.Associations}, &metadatas)
	} else {
		resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &metadatas)
	}
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = MetadatasToPB(metadatas)
	}
	resp.Total = resp.Records
}

func GetAllMetadatas(req *apipb.QueryMetadataRequest) (mds []*Metadata, err error) {
	db := dbClient.DB()
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}
	err = db.Preload("MetadataFields").Find(&mds).Error
	return
}

func GetMetadataById(id string) (*Metadata, error) {
	md := &Metadata{}
	err := dbClient.DB().Preload("MetadataFields").Where("id = ?", id).First(&md).Error
	if err == nil {
		md.Sort()
	}
	return md, err
}

func GetMetadataFieldByMDId(mdID string) ([]*MetadataField, error) {
	var fields []*MetadataField
	err := dbClient.DB().Order("`order`").Model(&MetadataField{}).Where("metadata_id = ?", mdID).Find(&fields).Error
	return fields, err
}

func GetMetadataByName(name string) (*Metadata, error) {
	md := &Metadata{}
	err := dbClient.DB().Preload("MetadataFields").Where("name = ?", name).First(md).Error
	return md, err
}

func UpdateMetadata(md *Metadata) error {
	if md.ParentID != "" {
		parent, err := GetMetadataById(md.ParentID)
		if err != nil {
			return err
		}
		md.Level = parent.Level + 1
	}
	md.FieldSort()
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldMetadata := &Metadata{}
		err := tx.Preload("MetadataFields").Preload(clause.Associations).Where("id = ?", md.ID).First(oldMetadata).Error
		if err != nil {
			return err
		}
		var deleteFile []string
		for _, oldFile := range oldMetadata.MetadataFields {
			flag := false
			for _, newFile := range md.MetadataFields {
				if newFile.ID == oldFile.ID {
					flag = true
				}
			}
			if !flag {
				deleteFile = append(deleteFile, oldFile.ID)
			}
		}
		if len(deleteFile) > 0 {
			err = tx.Unscoped().Delete(&MetadataField{}, "id in ?", deleteFile).Error
			if err != nil {
				return err
			}
		}

		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(tx, md, true, []string{"created_at"}, "id <> ? and `system`=? and  name = ? and project_id=? and tenant_id=?", md.ID, md.System, md.Name, md.ProjectID, md.TenantID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同元数据")
		}

		return nil
	})
}

func GetMetadataTree(req *apipb.QueryMetadataRequest) (list []*Metadata, total int64, err error) {
	var menuList []*Metadata
	treeMap, err := getMetadataTreeMap(req.TenantID, req.ProjectID)
	menuList = treeMap[""]
	for i := 0; i < len(menuList); i++ {
		err = getMetadataChildrenList(menuList[i], treeMap)
	}
	return menuList, total, err
}

func getMetadataTreeMap(tenantID, projectID string) (treeMap map[string][]*Metadata, err error) {
	var allMetadatas []*Metadata
	treeMap = make(map[string][]*Metadata)
	db := dbClient.DB()
	if tenantID != "" {
		db = db.Where("tenant_id = ?", tenantID)
	}

	if projectID != "" {
		db = db.Where("project_id = ?", projectID)
	}
	err = db.Order("level").Find(&allMetadatas).Error
	for _, v := range allMetadatas {
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	return treeMap, err
}

func getMetadataChildrenList(md *Metadata, treeMap map[string][]*Metadata) (err error) {
	md.Children = treeMap[md.ID]
	for i := 0; i < len(md.Children); i++ {
		err = getMetadataChildrenList(md.Children[i], treeMap)
	}
	return err
}

func CopyMetadata(id string) error {
	from, err := GetMetadataById(id)
	if err != nil {
		return err
	}
	to := &Metadata{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreateMetadata(to)
}
