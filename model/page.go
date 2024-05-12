package model

import (
	"errors"
	"fmt"
	"sort"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/curd/service"
	"github.com/CloudSilk/pkg/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Page struct {
	model.TenantModel
	ProjectID          string       `gorm:"index;size:36"`
	Name               string       `json:"name" gorm:"size:100;uniqueindex:page_uidx1" `
	Enable             bool         `json:"enable" gorm:"comment:是否启用"`
	MetadataID         string       `json:"metadataID" gorm:"" `
	Metadata           *Metadata    `json:"metadata" gorm:"" copier:"-"`
	PageSize           uint32       `json:"pageSize" gorm:"comment:每页数量" `
	Editable           string       `json:"editable" gorm:"size:100" `
	ShowIndex          bool         `json:"showIndex" gorm:"comment:是否显示序号"`
	ShowSelection      bool         `json:"showSelection" gorm:"显示批量操作"`
	ToolBar            *PageToolBar `json:"toolBar" gorm:"" `
	Fields             []*PageField `json:"fields" gorm:"" `
	Path               string       `json:"path" gorm:"size:200;comment:接口前缀,例如/api/core/auth/user"`
	Title              string       `json:"title" gorm:"size:100"`
	Description        string       `json:"description" gorm:"size:500"`
	SearchDefaultValue string       `json:"searchDefaultValue" gorm:"size:500"`
	AddDefaultValue    string       `json:"addDefaultValue"`
	EditFormID         string       `json:"editFormID"`
	SearchFormID       string       `json:"searchFormID"`
	AddFormID          string       `json:"addFormID"`
	ViewFormID         string       `json:"viewFormID"`
	Type               int32        `json:"type" gorm:"index;comment:1-表格 2-树形"`
	SubmitBefore       string       `json:"submitBefore" gorm:"comment:提交前执行"`
	SubmitAfter        string       `json:"submitAfter" gorm:"comment:提交成功后执行"`
	LoadDetailBefore   string       `json:"loadDetailBefore" gorm:"comment:加载明细前执行"`
	LoadDetailAfter    string       `json:"loadDetailAfter" gorm:"comment:加载明细成功后执行"`
	QueryBefore        string       `json:"queryBefore" gorm:"comment:查询前执行"`
	QueryAfter         string       `json:"queryAfter" gorm:"comment:查询成功后执行"`
	LabelField         string       `json:"labelField" gorm:"size:100"`
	ValueField         string       `json:"valueField" gorm:"size:100"`

	ScrollX int32 `json:"scrollX" gorm:"comment:横向滚动条"`

	ListKeyField         string `json:"listKeyField" gorm:"size:100"`
	ListAvatarField      string `json:"listAvatarField" gorm:"size:100"`
	ListTitleField       string `json:"listTitleField" gorm:"size:100"`
	ListDescriptionField string `json:"listDescriptionField" gorm:"size:100"`
	ListContentField     string `json:"listContentField" gorm:"size:100"`
	ListLoadType         int32  `json:"listLoadType"`

	CardAvatarField      string        `json:"cardAvatarField" gorm:"size:100"`
	CardTitleField       string        `json:"cardTitleField" gorm:"size:100"`
	CardDescriptionField string        `json:"cardDescriptionField" gorm:"size:100"`
	CardContentField     string        `json:"cardContentField" gorm:"size:100"`
	CardLoadType         int32         `json:"cardLoadType"`
	CardImageField       string        `json:"cardImageField" gorm:"size:100"`
	Buttons              []*PageButton `json:"buttons"`

	ProListGhost           bool   `json:"proListGhost"`
	ProListCardActionProps string `json:"proListCardActionProps" gorm:"size:50"`

	ProListShowTitle      bool   `json:"proListShowTitle"`
	ProListTitleDataIndex string `json:"proListTitleDataIndex" gorm:"size:100"`
	ProListTitleValueType string `json:"proListTitleValueType" gorm:"size:50"`
	ProListTitleRender    string `json:"proListTitleRender" gorm:"size:500"`

	ProListShowSubTitle      bool   `json:"proListShowSubTitle"`
	ProListSubTitleDataIndex string `json:"proListSubTitleDataIndex" gorm:"size:100"`
	ProListSubTitleValueType string `json:"proListSubTitleValueType" gorm:"size:50"`
	ProListSubTitleRender    string `json:"proListSubTitleRender" gorm:"size:500"`

	ProListShowMetaType  bool   `json:"proListShowMetaType"`
	ProListTypeDataIndex string `json:"proListTypeDataIndex" gorm:"size:100"`
	ProListTypeValueType string `json:"proListTypeValueType" gorm:"size:50"`
	ProListTypeRender    string `json:"proListTypeRender" gorm:"size:500"`

	ProListShowAvatar      bool   `json:"proListShowAvatar"`
	ProListAvatarDataIndex string `json:"proListAvatarDataIndex" gorm:"size:100"`
	ProListAvatarValueType string `json:"proListAvatarValueType" gorm:"size:50"`
	ProListAvatarRender    string `json:"proListAvatarRender" gorm:"size:500"`

	ProListShowContent      bool   `json:"proListShowContent"`
	ProListContentDataIndex string `json:"proListContentDataIndex" gorm:"size:100"`
	ProListContentValueType string `json:"proListContentValueType" gorm:"size:50"`
	ProListContentRender    string `json:"proListContentRender" gorm:"size:500"`

	ProListShowMetaExtra  bool   `json:"proListShowMetaExtra"`
	ProListExtraDataIndex string `json:"proListExtraDataIndex" gorm:"size:100"`
	ProListExtraValueType string `json:"proListExtraValueType" gorm:"size:50"`
	ProListExtraRender    string `json:"proListExtraRender" gorm:"size:500"`

	ProListShowActions    string `json:"proListShowActions" gorm:"size:20;comment:hover | always"`
	ProListShowType       int32  `json:"proListShowType" gorm:"comment:1-List 2-Card"`
	ProListShowExtra      string `json:"proListShowExtra" gorm:"size:20;comment:hover | always"`
	ProListItemClick      string `json:"proListItemClick" gorm:"size:500"`
	ProListItemMouseEnter string `json:"proListItemMouseEnter" gorm:"size:500"`

	ListGridTypeGutter int32  `json:"listGridTypeGutter"`
	ListGridTypeColumn int32  `json:"listGridTypeColumn"`
	ListItemLayout     string `json:"listItemLayout" gorm:"size:20"`
	ListExpandable     bool   `json:"listExpandable"`

	IsMust   bool   `json:"isMust" gorm:"index;comment:系统必须要有的数据"`
	IsChild  bool   `json:"isChild" gorm:"index;comment:子列表，应用于左右分栏"`
	Children string `json:"children" gorm:"size:200;comment:子列表的PageName,多个使用逗号隔开，应用于左右分栏等"`
	Bordered bool   `json:"bordered" gorm:"comment:是否显示边框"`
}

type PageField struct {
	model.Model
	PageID         string `json:"pageID" gorm:"" copier:"-"`
	Name           string `json:"name" gorm:"size:100;comment:字段名" copier:"-"`
	Title          string `json:"title" gorm:"size:100;comment:显示名称" copier:"-"`
	Copyable       bool   `json:"copyable" gorm:"comment:显示复制按钮" copier:"-"`
	Ellipsis       bool   `json:"ellipsis" gorm:"comment:是否自动缩略" copier:"-"`
	RowKey         bool   `json:"rowKey" gorm:"comment:Row Key" copier:"-"`
	Sort           int32  `json:"sort"`
	EnableSort     bool   `json:"enableSort" gorm:"comment:启用排序" copier:"-"`
	Fixed          string `json:"fixed" gorm:"comment:固定列;size:20" copier:"-"`
	ShowInTable    bool   `json:"showInTable"`
	ValueEnum      string `json:"valueEnum" gorm:"size:500;comment:枚举值转换"`
	Component      string `json:"component" gorm:"size:200;comment:组件"`
	ComponentProps string `json:"componentProps" gorm:"size:500"`
	DataType       string `json:"dataType" gorm:"size:100;comment:数据类型"`
	LabelField     string `json:"labelField" gorm:"size:100"`
	ValueField     string `json:"valueField" gorm:"size:100"`
	Align          string `json:"align" gorm:"size:20"`
	Width          string `json:"width" gorm:"size:20"`
}

type PageToolBar struct {
	model.Model
	PageID         string `json:"pageID" gorm:"" copier:"-"`
	FullScreen     bool   `json:"fullScreen" gorm:"comment:全屏" copier:"-"`
	Reload         bool   `json:"reload" gorm:"comment:刷新" copier:"-"`
	Setting        bool   `json:"setting" gorm:"" copier:"-"`
	Render         string `json:"render" gorm:"size:500" copier:"-"`
	ShowAdd        bool   `json:"showAdd" gorm:"comment:是否显示新增按钮"`
	AddPermission  string `json:"addPermission" gorm:"size:200;comment:新增按钮权限"`
	AddScript      string `json:"addScript"`
	ShowExport     bool   `json:"showExport" gorm:"comment:是否显示导出按钮"`
	ShowImport     bool   `json:"showImport" gorm:"comment:是否显示导入按钮"`
	RowSelection   bool   `json:"rowSelection" gorm:"comment:是否显示可选框"`
	ShowBatchDel   bool   `json:"showBatchDel" gorm:"comment:是否显示批量删除按钮"`
	BatchDelUri    string `json:"batchDelUri" gorm:"size:100;comment:批量删除API接口地址"`
	ExportUri      string `json:"exportUri" gorm:"size:100;comment:导出API接口地址"`
	ImportUri      string `json:"importUri" gorm:"size:100;comment:导入API接口地址"`
	ImportMulti    bool   `json:"importMulti" gorm:"comment:导入多个文件"`
	ImportMaxCount int32  `json:"importMaxCount" gorm:"comment:上传最大数量"`
	ImportFormID   string `json:"importFormID" gorm:"size:36;comment:导入表单"`
}

type PageButton struct {
	ID           string `json:"id" copier:"-"`
	PageID       string `json:"pageID" gorm:"" copier:"-"`
	Key          string `json:"key" gorm:"size:100"`
	Label        string `json:"label" gorm:"size:200"`
	Expanded     bool   `json:"expanded"`
	ShowType     string `json:"showType" gorm:"size:50"`
	Href         string `json:"href" gorm:"size:200"`
	HrefFunc     string `json:"hrefFunc"`
	Script       string `json:"script"`
	HiddenScript string `json:"hiddenScript"`
	Index        int32  `json:"index"`
	Enable       bool   `json:"enable" gorm:"comment:是否启用;index"`
	//按钮权限有两种控制方式
	//先判断角色权限
	//在判断数据权限
	Permission   string `json:"permission" gorm:"size:200;comment:按钮权限"`
	ShowPosition int32  `json:"showPosition" gorm:"default:0;comment:0-显示在列表行 1-在搜索框显示"`
	FormID       string `json:"formID" gorm:"size:36"`
}

func (u *PageButton) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

func SortFields(fields []*PageField) {
	for i, field := range fields {
		field.Sort = int32(i) + 1
	}
}

func SortButtons(buttons []*PageButton) {
	sort.Slice(buttons, func(i, j int) bool {
		return buttons[i].Index < buttons[j].Index
	})
}

func CreatePage(m *Page) error {
	SortFields(m.Fields)
	SortButtons(m.Buttons)
	count, err := statisticPageCount(dbClient.DB(), m.TenantID, m.ProjectID)
	if err != nil {
		return err
	}

	expired, projectPageCount, err := service.GetProjectPageCount(m.ProjectID)
	if err != nil {
		return err
	}
	if expired {
		return errors.New("账号使用期限已过，你可以联系管理员!")
	}

	if projectPageCount > 0 && projectPageCount <= int32(count) {
		return fmt.Errorf("只能创建 %d 个页面", projectPageCount)
	}
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {

		duplication, err := dbClient.CreateWithCheckDuplicationWithDB(tx, m, " name =? and project_id=?", m.Name, m.ProjectID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同页面配置")
		}
		return nil
	})
}

func DeleteFields(tx *gorm.DB, old, m *Page) error {
	var deleteFields []string
	for _, oldField := range old.Fields {
		flag := false
		for _, newField := range m.Fields {
			if newField.ID == oldField.ID {
				flag = true
			}
		}
		if !flag {
			deleteFields = append(deleteFields, oldField.ID)
		}
	}

	if len(deleteFields) > 0 {
		err := tx.Unscoped().Delete(&PageField{}, "id in ?", deleteFields).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteButtons(tx *gorm.DB, old, m *Page) error {
	var deleteIDs []string
	for _, oldObj := range old.Buttons {
		flag := false
		for _, newObj := range m.Buttons {
			if newObj.ID == oldObj.ID {
				flag = true
			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&PageButton{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdatePage(m *Page) error {
	SortFields(m.Fields)
	SortButtons(m.Buttons)
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldPage := &Page{}
		err := tx.Preload("Fields").Preload(clause.Associations).Where("id = ?", m.ID).First(oldPage).Error
		if err != nil {
			return err
		}

		err = DeleteFields(tx, oldPage, m)
		if err != nil {
			return err
		}

		err = DeleteButtons(tx, oldPage, m)
		if err != nil {
			return err
		}

		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id != ?  and  name =? and project_id=?", m.ID, m.Name, m.ProjectID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同页面配置")
		}

		return nil
	})
}

func QueryPage(req *apipb.QueryPageRequest, resp *apipb.QueryPageResponse, preload, sorted bool) {
	db := dbClient.DB().Model(&Page{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Enable != -1 {
		db = db.Where("enable = ?", req.Enable == 1)
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}

	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}
	if len(req.Ids) > 0 {
		db = db.Where("id in ?", req.Ids)
	}
	if req.IsChild > 0 {
		db = db.Where("is_child = ?", req.IsChild == 1)
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
	var pages []*Page
	if preload {
		resp.Records, resp.Pages, err = dbClient.PageQueryWithPreload(db, req.PageSize, req.PageIndex, OrderStr, []string{"Metadata.MetadataFields", "Fields", clause.Associations}, &pages)
	} else {
		resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &pages)
	}

	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		if sorted {
			for _, m := range pages {
				sort.Slice(m.Fields, func(i, j int) bool {
					return m.Fields[i].Sort < m.Fields[j].Sort
				})
				SortButtons(m.Buttons)
			}
		}
		resp.Data = PagesToPB(pages)
	}
	resp.Total = resp.Records
}

func GetAllPage(req *apipb.QueryPageRequest) (list []*Page, err error) {
	db := dbClient.DB()
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}

	if req.IsChild > 0 {
		db = db.Where("is_child = ?", req.IsChild == 1)
	}
	err = db.Find(&list).Error
	return
}

func GetPageByID(id string) (*Page, error) {
	m := &Page{}
	err := dbClient.DB().Preload("Metadata.MetadataFields").Preload("Fields").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	sort.Slice(m.Fields, func(i, j int) bool {
		return m.Fields[i].Sort < m.Fields[j].Sort
	})
	SortButtons(m.Buttons)
	return m, err
}

func GetPageByName(name string) (*Page, error) {
	m := &Page{}
	err := dbClient.DB().Preload("Metadata.MetadataFields").Preload("Fields").Preload(clause.Associations).Where("name = ? and enable=?", name, true).First(m).Error
	sort.Slice(m.Fields, func(i, j int) bool {
		return m.Fields[i].Sort < m.Fields[j].Sort
	})
	SortButtons(m.Buttons)
	return m, err
}

func DeletePage(id string) (err error) {
	return dbClient.DB().Delete(&Page{}, "id=?", id).Error
}

func CopyPage(id string) error {
	from, err := GetPageByID(id)
	if err != nil {
		return err
	}
	to := &Page{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreatePage(to)
}

func EnablePage(id string, enable bool) error {
	err := dbClient.DB().Model(&Page{}).Where("id=?", id).Update("enable", enable).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPageByIDs(ids []string) ([]Page, error) {
	var list []Page
	err := dbClient.DB().Preload("Metadata.MetadataFields").Preload("Fields").Preload(clause.Associations).Where("id in ?", ids).Find(&list).Error
	return list, err
}

func statisticPageCount(db *gorm.DB, tenantID, projectID string) (int64, error) {
	db = db.Model(&Page{})
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
