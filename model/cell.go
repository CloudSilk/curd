package model

import (
	"errors"
	"fmt"
	"sort"

	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/curd/service"
	commonmodel "github.com/CloudSilk/pkg/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateCell(m *Cell) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		count, err := statisticCellCount(tx, m.TenantID, m.ProjectID)
		if err != nil {
			return err
		}

		expired, projectCellCount, err := service.GetProjectCellCount(m.ProjectID)
		if err != nil {
			return err
		}
		if expired {
			return fmt.Errorf("账号使用期限已过，你可以联系管理员!")
		}

		if projectCellCount > 0 && projectCellCount <= int32(count) {
			return fmt.Errorf("只能创建 %d 个Cell", projectCellCount)
		}
		duplication, err := dbClient.CreateWithCheckDuplicationWithDB(tx, m, " name=? and project_id=? and tenant_id = ?", m.Name, m.ProjectID, m.TenantID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同Cell")
		}
		return nil
	})
}

//删除子表

func DeleteMarkup(tx *gorm.DB, old, m *Cell) error {
	var deleteIDs []string
	for _, oldObj := range old.Markup {
		flag := false
		for _, newObj := range m.Markup {
			if newObj.ID == oldObj.ID {
				flag = true

			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&CellMarkup{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteAttrs(tx *gorm.DB, old, m *Cell) error {
	var deleteIDs []string
	for _, oldObj := range old.Attrs {
		flag := false
		for _, newObj := range m.Attrs {
			if newObj.ID == oldObj.ID {
				flag = true

			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&CellAttrs{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteConnectings(tx *gorm.DB, old, m *Cell) error {
	var deleteIDs []string
	for _, oldObj := range old.Connectings {
		flag := false
		for _, newObj := range m.Connectings {
			if newObj.ID == oldObj.ID {
				flag = true

			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&CellConnecting{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCell(m *Cell) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldCell := &Cell{}
		//TODO 需要TenantID作为限制条件，避免出现漏洞
		err := tx.Preload("Markup").Preload("Attrs").Preload(clause.Associations).Where("id = ?", m.ID).First(oldCell).Error
		if err != nil {
			return err
		}
		err = DeleteMarkup(tx, oldCell, m)
		if err != nil {
			return err
		}

		err = DeleteConnectings(tx, oldCell, m)
		if err != nil {
			return err
		}

		err = DeleteAttrs(tx, oldCell, m)
		if err != nil {
			return err
		}

		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  and  name=? and project_id=? and tenant_id = ?", m.ID, m.Name, m.ProjectID, m.TenantID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同Cell")
		}

		return nil
	})
}

func DeleteCell(id string) (err error) {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Unscoped().Delete(&CellAttrs{}, "cell_id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&CellMarkup{}, "cell_id = ?", id).Error
		if err != nil {
			return err
		}
		return tx.Unscoped().Delete(&Cell{}, "id=?", id).Error
	})
}

func QueryCell(req *apipb.QueryCellRequest, resp *apipb.QueryCellResponse, preload bool) {
	db := dbClient.DB().Model(&Cell{})
	if preload {
		db = db.Preload("Attrs").Preload("Connectings").Preload(clause.Associations)
	}
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}

	if req.System != "" {
		db = db.Where("`system` = ?", req.System)
	}

	if req.View != "" {
		db = db.Where("`view` = ?", req.View)
	}

	if req.Shape != "" {
		db = db.Where("shape = ?", req.Shape)
	}
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}
	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}

	if req.Group != "" {
		db = db.Where("`group` = ?", req.Group)
	}
	if req.IsEdge > 0 {
		db = db.Where("is_edge=?", req.IsEdge == 1)
	}

	// orderStr := "`system`,`index`"
	orderStr := "`updated_at` desc"
	if req.OrderField != "" {
		if req.Desc {
			orderStr = req.OrderField + " desc"
		} else {
			orderStr = req.OrderField
		}
	}
	var err error
	var list []*Cell

	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list, nil)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = CellsToPB(list)
	}
	resp.Total = resp.Records
}

func GetCellByID(id string) (*Cell, error) {
	m := &Cell{}
	err := dbClient.DB().Preload("Markup").Preload("Attrs").Preload("Connectings").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

type Cell struct {
	commonmodel.Model
	TenantID          string `gorm:"index;size:36"`
	ProjectID         string `gorm:"index;size:36"`
	Name              string `json:"name" gorm:"size:0;index" `
	System            string `json:"system" gorm:"size:100;index;comment:例如BPM、IoT、HR" `
	Group             string `json:"group" gorm:"size:100;index;comment:例如BPM中分组，开始，结束，任务，事件等" `
	Shape             string `json:"shape" gorm:"size:100;index;comment:继承自X6的基础形状，例如Circle、Rect等" `
	IdPrefix          string `json:"idPrefix" gorm:"size:100" `
	PropertyForm      string `json:"propertyForm" gorm:"size:36" `
	View              string `json:"view" gorm:"size:500;comment:渲染节点/边的视图" `
	DefaultLabel      string `json:"defaultLabel" gorm:"size:100" `
	Icon              string `json:"icon" gorm:"comment:可以是url，也可以是path data，也可以是默认的icon里面获取" `
	IconSource        int32  `json:"iconSource" gorm:"comment:1-URL 2-PathData 3-IconComponent" `
	IsEdge            bool   `json:"isEdge" gorm:"index" `
	Common            bool   `json:"common" gorm:"index" `
	Resizing          bool   `json:"resizing" gorm:"index;comment:定义组件可以调整大小" `
	MustTarget        bool   `json:"mustTarget" gorm:"index;comment:必须是连线的结束" `
	MustSource        bool   `json:"mustSource" gorm:"index;comment:必须是连线的开始" `
	Width             int32  `json:"width" gorm:"" `
	Height            int32  `json:"height" gorm:"" `
	Width2            int32  `json:"width2" gorm:"" `
	Height2           int32  `json:"height2" gorm:"" `
	Parent            bool   `gorm:"index"`
	Index             int32
	DefaultEdge       bool          `gorm:"comment:例如默认Edge"`
	Other             string        `gorm:"comment:其他属性,json格式"`
	DefaultLabelAttrs string        `gorm:"comment:默认Label其他属性,json格式"`
	Ports             string        `gorm:"comment:链接桩属性,json格式"`
	FormDefaultValue  string        `gorm:"comment:表单默认值,json格式"`
	Attrs             []*CellAttrs  `json:"attrs" gorm:""`
	Markup            []*CellMarkup `json:"markup" gorm:""`
	Connectings       []*CellConnecting
	IsMust            bool `json:"isMust" gorm:"index;comment:系统必须要有的数据"`
}

func (c *Cell) Sort() {
	sort.Slice(c.Markup, func(i, j int) bool {
		return c.Markup[i].Index < c.Markup[j].Index
	})
	sort.Slice(c.Attrs, func(i, j int) bool {
		return c.Attrs[i].Name < c.Attrs[j].Name
	})
}

type CellMarkup struct {
	commonmodel.Model
	TextContent    string `json:"textContent" gorm:"size:200" `
	CellID         string `json:"cellID" gorm:"size:36;index" `
	Other          string `json:"other" gorm:"comment:json格式" `
	Children       string `json:"children" gorm:"size:500" `
	ClassName      string `json:"className" gorm:"size:100" `
	Selector       string `json:"selector" gorm:"size:100" `
	GroupSelector  string `json:"groupSelector" gorm:"size:100" `
	TagName        string `json:"tagName" gorm:"size:100" `
	Attrs          string `json:"attrs" gorm:"size:200" `
	Style          string `json:"style" `
	Index          int32
	IsDefaultLabel bool
}
type CellAttrs struct {
	commonmodel.Model
	Stroke             string `json:"stroke" gorm:"size:50"`
	SelectedStroke     string `json:"selectedStroke" gorm:"size:50"`
	FontSize           int32  `json:"fontSize" gorm:"" `
	Other              string `json:"other" gorm:"comment:json格式"`
	TextAnchor         string `json:"textAnchor" gorm:"size:10" `
	Name               string `json:"name" gorm:"size:100;index" `
	Fill               string `json:"fill" gorm:"size:50" `
	SelectedFill       string `json:"selectedFill" gorm:"size:50" `
	Ref                string `json:"ref" gorm:"size:100" `
	Magnet             bool   `json:"magnet" gorm:"" `
	TextVerticalAnchor string `json:"textVerticalAnchor" gorm:"size:10" `
	CellID             string `json:"cellID" gorm:"size:36;index" `
	IsDefaultLabel     bool
	LinkHref           string `gorm:"size:36"`
}

type CellConnecting struct {
	commonmodel.Model
	CellID      string `json:"cellID" gorm:"size:36;index" `
	AnotherCell string `json:"anotherCell" gorm:"size:36;null" `
	Edge        string `json:"edge" gorm:"size:36;null" `
	Direct      int32  `gorm:"comment:连线方向,1-开始 2-结束 3-任意"`
}

func GetAllCells(req *apipb.GetAllCellRequest) (list []*Cell, err error) {
	db := dbClient.DB()
	if req.System != "" {
		db = db.Where("`system` = ?", req.System)
	}
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}
	if req.ProjectID != "" {
		db = db.Where("project_id = ?", req.ProjectID)
	}
	err = db.Preload("Markup").Preload("Attrs").Preload("Connectings").Preload(clause.Associations).Order("`index`").Find(&list).Error
	return
}

func CopyCell(id string) error {
	from, err := GetCellByID(id)
	if err != nil {
		return err
	}
	to := &Cell{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreateCell(to)
}

func EnableCell(id string, enable bool) error {
	err := dbClient.DB().Model(&Cell{}).Where("id=?", id).Update("enable", enable).Error
	if err != nil {
		return err
	}
	return nil
}

func PBToCells(in []*apipb.CellInfo) []*Cell {
	var result []*Cell
	for _, c := range in {
		result = append(result, PBToCell(c))
	}
	return result
}

func PBToCell(in *apipb.CellInfo) *Cell {
	return &Cell{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		ProjectID:         in.ProjectID,
		MustSource:        in.MustSource,
		Height:            in.Height,
		Name:              in.Name,
		System:            in.System,
		DefaultLabel:      in.DefaultLabel,
		View:              in.View,
		Icon:              in.Icon,
		PropertyForm:      in.PropertyForm,
		IdPrefix:          in.IdPrefix,
		IconSource:        in.IconSource,
		Shape:             in.Shape,
		IsEdge:            in.IsEdge,
		Common:            in.Common,
		Resizing:          in.Resizing,
		Group:             in.Group,
		MustTarget:        in.MustTarget,
		Width:             in.Width,
		TenantID:          in.TenantID,
		Height2:           in.Height2,
		Width2:            in.Width2,
		Parent:            in.Parent,
		Index:             in.Index,
		DefaultEdge:       in.DefaultEdge,
		Other:             in.Other,
		DefaultLabelAttrs: in.DefaultLabelAttrs,
		Ports:             in.Ports,
		Attrs:             PBToCellAttrses(in.Attrs),
		Markup:            PBToCellMarkups(in.Markup),
		Connectings:       PBToCellConnectings(in.Connectings),
		IsMust:            in.IsMust,
		FormDefaultValue:  in.FormDefaultValue,
	}
}

func CellsToPB(in []*Cell) []*apipb.CellInfo {
	var list []*apipb.CellInfo
	for _, f := range in {
		list = append(list, CellToPB(f))
	}
	return list
}

func CellToPB(in *Cell) *apipb.CellInfo {
	in.Sort()
	return &apipb.CellInfo{
		Id:                in.ID,
		ProjectID:         in.ProjectID,
		MustSource:        in.MustSource,
		Height:            in.Height,
		Name:              in.Name,
		System:            in.System,
		DefaultLabel:      in.DefaultLabel,
		View:              in.View,
		Icon:              in.Icon,
		PropertyForm:      in.PropertyForm,
		IdPrefix:          in.IdPrefix,
		IconSource:        in.IconSource,
		Shape:             in.Shape,
		IsEdge:            in.IsEdge,
		Common:            in.Common,
		Resizing:          in.Resizing,
		Group:             in.Group,
		MustTarget:        in.MustTarget,
		Width:             in.Width,
		TenantID:          in.TenantID,
		Height2:           in.Height2,
		Width2:            in.Width2,
		Parent:            in.Parent,
		Index:             in.Index,
		DefaultEdge:       in.DefaultEdge,
		Other:             in.Other,
		DefaultLabelAttrs: in.DefaultLabelAttrs,
		Ports:             in.Ports,
		Markup:            CellMarkupsToPB(in.Markup),
		Attrs:             CellAttrsesToPB(in.Attrs),
		Connectings:       CellConnectingsToPB(in.Connectings),
		IsMust:            in.IsMust,
		FormDefaultValue:  in.FormDefaultValue,
	}
}

func PBToCellMarkups(in []*apipb.CellMarkup) []*CellMarkup {
	var result []*CellMarkup
	for _, c := range in {
		result = append(result, PBToCellMarkup(c))
	}
	return result
}

func PBToCellMarkup(in *apipb.CellMarkup) *CellMarkup {
	return &CellMarkup{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		TextContent:    in.TextContent,
		CellID:         in.CellID,
		Other:          in.Other,
		Children:       in.Children,
		ClassName:      in.ClassName,
		Selector:       in.Selector,
		GroupSelector:  in.GroupSelector,
		TagName:        in.TagName,
		Attrs:          in.Attrs,
		Style:          in.Style,
		Index:          in.Index,
		IsDefaultLabel: in.IsDefaultLabel,
	}
}

func CellMarkupsToPB(in []*CellMarkup) []*apipb.CellMarkup {
	var list []*apipb.CellMarkup
	for _, f := range in {
		list = append(list, CellMarkupToPB(f))
	}
	return list
}

func CellMarkupToPB(in *CellMarkup) *apipb.CellMarkup {
	return &apipb.CellMarkup{
		Id:             in.ID,
		TextContent:    in.TextContent,
		CellID:         in.CellID,
		Other:          in.Other,
		Children:       in.Children,
		ClassName:      in.ClassName,
		Selector:       in.Selector,
		GroupSelector:  in.GroupSelector,
		TagName:        in.TagName,
		Attrs:          in.Attrs,
		Style:          in.Style,
		Index:          in.Index,
		IsDefaultLabel: in.IsDefaultLabel,
	}
}

func PBToCellAttrses(in []*apipb.CellAttrs) []*CellAttrs {
	var result []*CellAttrs
	for _, c := range in {
		result = append(result, PBToCellAttrs(c))
	}
	return result
}

func PBToCellAttrs(in *apipb.CellAttrs) *CellAttrs {
	return &CellAttrs{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Stroke:             in.Stroke,
		FontSize:           in.FontSize,
		Other:              in.Other,
		TextAnchor:         in.TextAnchor,
		Name:               in.Name,
		Fill:               in.Fill,
		Ref:                in.Ref,
		Magnet:             in.Magnet,
		TextVerticalAnchor: in.TextVerticalAnchor,
		CellID:             in.CellID,
		SelectedStroke:     in.SelectedStroke,
		SelectedFill:       in.SelectedFill,
		IsDefaultLabel:     in.IsDefaultLabel,
		LinkHref:           in.LinkHref,
	}
}

func CellAttrsesToPB(in []*CellAttrs) []*apipb.CellAttrs {
	var list []*apipb.CellAttrs
	for _, f := range in {
		list = append(list, CellAttrsToPB(f))
	}
	return list
}

func CellAttrsToPB(in *CellAttrs) *apipb.CellAttrs {
	return &apipb.CellAttrs{
		Id:                 in.ID,
		Stroke:             in.Stroke,
		FontSize:           in.FontSize,
		Other:              in.Other,
		TextAnchor:         in.TextAnchor,
		Name:               in.Name,
		Fill:               in.Fill,
		Ref:                in.Ref,
		Magnet:             in.Magnet,
		TextVerticalAnchor: in.TextVerticalAnchor,
		CellID:             in.CellID,
		SelectedStroke:     in.SelectedStroke,
		SelectedFill:       in.SelectedFill,
		IsDefaultLabel:     in.IsDefaultLabel,
		LinkHref:           in.LinkHref,
	}
}

func PBToCellConnectings(in []*apipb.CellConnecting) []*CellConnecting {
	var result []*CellConnecting
	for _, c := range in {
		result = append(result, PBToCellConnecting(c))
	}
	return result
}

func PBToCellConnecting(in *apipb.CellConnecting) *CellConnecting {
	return &CellConnecting{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		CellID:      in.CellID,
		AnotherCell: in.AnotherCell,
		Edge:        in.Edge,
		Direct:      in.Direct,
	}
}

func CellConnectingsToPB(in []*CellConnecting) []*apipb.CellConnecting {
	var list []*apipb.CellConnecting
	for _, f := range in {
		list = append(list, CellConnectingToPB(f))
	}
	return list
}

func CellConnectingToPB(in *CellConnecting) *apipb.CellConnecting {
	return &apipb.CellConnecting{
		Id:          in.ID,
		CellID:      in.CellID,
		AnotherCell: in.AnotherCell,
		Edge:        in.Edge,
		Direct:      in.Direct,
	}
}

func statisticCellCount(db *gorm.DB, tenantID, projectID string) (int64, error) {
	db = db.Model(&Cell{})
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
