package model

import (
	apipb "github.com/CloudSilk/curd/proto"
	commonmodel "github.com/CloudSilk/pkg/model"
)

func PBToPage(in *apipb.PageInfo) *Page {
	return &Page{
		TenantModel: commonmodel.TenantModel{
			Model: commonmodel.Model{
				ID: in.Id,
			},
			TenantID: in.TenantID,
		},
		ProjectID:          in.ProjectID,
		Name:               in.Name,
		Enable:             in.Enable,
		MetadataID:         in.MetadataID,
		PageSize:           in.PageSize,
		Editable:           in.Editable,
		ShowIndex:          in.ShowIndex,
		ShowSelection:      in.ShowSelection,
		Path:               in.Path,
		Title:              in.Title,
		Description:        in.Description,
		SearchDefaultValue: in.SearchDefaultValue,
		AddDefaultValue:    in.AddDefaultValue,
		EditFormID:         in.EditFormID,
		SearchFormID:       in.SearchFormID,
		AddFormID:          in.AddFormID,
		ViewFormID:         in.ViewFormID,
		Type:               in.Type,
		SubmitBefore:       in.SubmitBefore,
		SubmitAfter:        in.SubmitAfter,
		LoadDetailBefore:   in.LoadDetailBefore,
		LoadDetailAfter:    in.LoadDetailAfter,
		QueryBefore:        in.QueryBefore,
		QueryAfter:         in.QueryAfter,
		LabelField:         in.LabelField,
		ValueField:         in.ValueField,

		ScrollX:              in.ScrollX,
		ListKeyField:         in.ListKeyField,
		ListAvatarField:      in.ListAvatarField,
		ListTitleField:       in.ListTitleField,
		ListDescriptionField: in.ListDescriptionField,
		ListContentField:     in.ListContentField,
		ListLoadType:         in.ListLoadType,
		CardAvatarField:      in.CardAvatarField,
		CardTitleField:       in.CardTitleField,
		CardDescriptionField: in.CardDescriptionField,
		CardContentField:     in.CardContentField,
		CardLoadType:         in.CardLoadType,
		CardImageField:       in.CardImageField,
		ToolBar:              PBToPageToolBar(in.ToolBar),
		Buttons:              PageButtonsToPB(in.Buttons),
		Fields:               PBToPageFields(in.Fields),

		ProListGhost:           in.ProListGhost,
		ProListCardActionProps: in.ProListCardActionProps,

		ProListShowTitle:      in.ProListShowTitle,
		ProListTitleDataIndex: in.ProListTitleDataIndex,
		ProListTitleValueType: in.ProListTitleValueType,
		ProListTitleRender:    in.ProListTitleRender,

		ProListShowSubTitle:      in.ProListShowSubTitle,
		ProListSubTitleDataIndex: in.ProListSubTitleDataIndex,
		ProListSubTitleValueType: in.ProListSubTitleValueType,
		ProListSubTitleRender:    in.ProListSubTitleRender,

		ProListShowMetaType:  in.ProListShowMetaType,
		ProListTypeDataIndex: in.ProListTypeDataIndex,
		ProListTypeValueType: in.ProListTypeValueType,
		ProListTypeRender:    in.ProListTypeRender,

		ProListShowAvatar:      in.ProListShowAvatar,
		ProListAvatarDataIndex: in.ProListAvatarDataIndex,
		ProListAvatarValueType: in.ProListAvatarValueType,
		ProListAvatarRender:    in.ProListAvatarRender,

		ProListShowContent:      in.ProListShowContent,
		ProListContentDataIndex: in.ProListContentDataIndex,
		ProListContentValueType: in.ProListContentValueType,
		ProListContentRender:    in.ProListContentRender,

		ProListShowMetaExtra:  in.ProListShowMetaExtra,
		ProListExtraDataIndex: in.ProListExtraDataIndex,
		ProListExtraValueType: in.ProListExtraValueType,
		ProListExtraRender:    in.ProListExtraRender,

		ProListShowActions:    in.ProListShowActions,
		ProListShowType:       in.ProListShowType,
		ProListShowExtra:      in.ProListShowExtra,
		ProListItemClick:      in.ProListItemClick,
		ProListItemMouseEnter: in.ProListItemMouseEnter,

		ListGridTypeGutter: in.ListGridTypeGutter,
		ListGridTypeColumn: in.ListGridTypeColumn,
		ListItemLayout:     in.ListItemLayout,
		ListExpandable:     in.ListExpandable,

		IsMust:   in.IsMust,
		IsChild:  in.IsChild,
		Children: in.Pages,
		Bordered: in.Bordered,
	}
}

func PageButtonsToPB(versions []*apipb.PageButton) []*PageButton {
	var list []*PageButton
	for _, btn := range versions {
		list = append(list, &PageButton{
			ID:           btn.Id,
			PageID:       btn.PageID,
			Key:          btn.Key,
			Label:        btn.Label,
			Expanded:     btn.Expanded,
			ShowType:     btn.ShowType,
			Href:         btn.Href,
			HrefFunc:     btn.HrefFunc,
			HiddenScript: btn.HiddenScript,
			Script:       btn.Script,
			Index:        btn.Index,
			Enable:       btn.Enable,
			Permission:   btn.Permission,
			ShowPosition: btn.ShowPosition,
			FormID:       btn.FormID,
		})
	}
	return list
}

func PBToPageButtons(btns []*PageButton) []*apipb.PageButton {
	var list []*apipb.PageButton
	for _, btn := range btns {
		list = append(list, &apipb.PageButton{
			Id:           btn.ID,
			PageID:       btn.PageID,
			Key:          btn.Key,
			Label:        btn.Label,
			Expanded:     btn.Expanded,
			ShowType:     btn.ShowType,
			Href:         btn.Href,
			HrefFunc:     btn.HrefFunc,
			HiddenScript: btn.HiddenScript,
			Script:       btn.Script,
			Index:        btn.Index,
			Enable:       btn.Enable,
			Permission:   btn.Permission,
			ShowPosition: btn.ShowPosition,
			FormID:       btn.FormID,
		})
	}
	return list
}

func PageToPB(in *Page) *apipb.PageInfo {
	return &apipb.PageInfo{
		TenantID:             in.TenantID,
		ProjectID:            in.ProjectID,
		Id:                   in.ID,
		Name:                 in.Name,
		Enable:               in.Enable,
		MetadataID:           in.MetadataID,
		PageSize:             in.PageSize,
		Editable:             in.Editable,
		ShowIndex:            in.ShowIndex,
		ShowSelection:        in.ShowSelection,
		Path:                 in.Path,
		Title:                in.Title,
		Description:          in.Description,
		SearchDefaultValue:   in.SearchDefaultValue,
		AddDefaultValue:      in.AddDefaultValue,
		EditFormID:           in.EditFormID,
		SearchFormID:         in.SearchFormID,
		AddFormID:            in.AddFormID,
		ViewFormID:           in.ViewFormID,
		Type:                 in.Type,
		SubmitBefore:         in.SubmitBefore,
		SubmitAfter:          in.SubmitAfter,
		LoadDetailBefore:     in.LoadDetailBefore,
		LoadDetailAfter:      in.LoadDetailAfter,
		QueryBefore:          in.QueryBefore,
		QueryAfter:           in.QueryAfter,
		LabelField:           in.LabelField,
		ValueField:           in.ValueField,
		ListKeyField:         in.ListKeyField,
		ListAvatarField:      in.ListAvatarField,
		ListTitleField:       in.ListTitleField,
		ListDescriptionField: in.ListDescriptionField,
		ListContentField:     in.ListContentField,
		ListLoadType:         in.ListLoadType,
		CardAvatarField:      in.CardAvatarField,
		CardTitleField:       in.CardTitleField,
		CardDescriptionField: in.CardDescriptionField,
		CardContentField:     in.CardContentField,
		CardLoadType:         in.CardLoadType,
		CardImageField:       in.CardImageField,
		Buttons:              PBToPageButtons(in.Buttons),
		ToolBar:              PageToolBarToPB(in.ToolBar),
		Fields:               PageFieldsToPB(in.Fields),
		CreatedAt:            in.CreatedAt.Unix(),
		UpdatedAt:            in.UpdatedAt.Unix(),

		ScrollX: in.ScrollX,

		ProListGhost:           in.ProListGhost,
		ProListCardActionProps: in.ProListCardActionProps,

		ProListShowTitle:      in.ProListShowTitle,
		ProListTitleDataIndex: in.ProListTitleDataIndex,
		ProListTitleValueType: in.ProListTitleValueType,
		ProListTitleRender:    in.ProListTitleRender,

		ProListShowSubTitle:      in.ProListShowSubTitle,
		ProListSubTitleDataIndex: in.ProListSubTitleDataIndex,
		ProListSubTitleValueType: in.ProListSubTitleValueType,
		ProListSubTitleRender:    in.ProListSubTitleRender,

		ProListShowMetaType:  in.ProListShowMetaType,
		ProListTypeDataIndex: in.ProListTypeDataIndex,
		ProListTypeValueType: in.ProListTypeValueType,
		ProListTypeRender:    in.ProListTypeRender,

		ProListShowAvatar:      in.ProListShowAvatar,
		ProListAvatarDataIndex: in.ProListAvatarDataIndex,
		ProListAvatarValueType: in.ProListAvatarValueType,
		ProListAvatarRender:    in.ProListAvatarRender,

		ProListShowContent:      in.ProListShowContent,
		ProListContentDataIndex: in.ProListContentDataIndex,
		ProListContentValueType: in.ProListContentValueType,
		ProListContentRender:    in.ProListContentRender,

		ProListShowMetaExtra:  in.ProListShowMetaExtra,
		ProListExtraDataIndex: in.ProListExtraDataIndex,
		ProListExtraValueType: in.ProListExtraValueType,
		ProListExtraRender:    in.ProListExtraRender,

		ProListShowActions:    in.ProListShowActions,
		ProListShowType:       in.ProListShowType,
		ProListShowExtra:      in.ProListShowExtra,
		ProListItemClick:      in.ProListItemClick,
		ProListItemMouseEnter: in.ProListItemMouseEnter,

		ListGridTypeGutter: in.ListGridTypeGutter,
		ListGridTypeColumn: in.ListGridTypeColumn,
		ListItemLayout:     in.ListItemLayout,
		ListExpandable:     in.ListExpandable,

		IsMust:   in.IsMust,
		IsChild:  in.IsChild,
		Pages:    in.Children,
		Bordered: in.Bordered,
	}
}

func PagesToPB(in []*Page) []*apipb.PageInfo {
	var list []*apipb.PageInfo
	for _, f := range in {
		list = append(list, PageToPB(f))
	}
	return list
}

func PBToPageToolBar(in *apipb.PageToolBar) *PageToolBar {
	return &PageToolBar{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		PageID:         in.PageID,
		FullScreen:     in.FullScreen,
		Reload:         in.Reload,
		Setting:        in.Setting,
		Render:         in.Render,
		ShowAdd:        in.ShowAdd,
		AddScript:      in.AddScript,
		AddPermission:  in.AddPermission,
		ShowExport:     in.ShowExport,
		ShowImport:     in.ShowImport,
		RowSelection:   in.RowSelection,
		ShowBatchDel:   in.ShowBatchDel,
		BatchDelUri:    in.BatchDelUri,
		ExportUri:      in.ExportUri,
		ImportUri:      in.ImportUri,
		ImportMulti:    in.ImportMulti,
		ImportMaxCount: in.ImportMaxCount,
		ImportFormID:   in.ImportFormID,
	}
}

func PageToolBarToPB(in *PageToolBar) *apipb.PageToolBar {
	if in == nil {
		return nil
	}
	return &apipb.PageToolBar{
		Id:             in.ID,
		PageID:         in.PageID,
		FullScreen:     in.FullScreen,
		Reload:         in.Reload,
		Setting:        in.Setting,
		Render:         in.Render,
		ShowAdd:        in.ShowAdd,
		AddScript:      in.AddScript,
		AddPermission:  in.AddPermission,
		ShowExport:     in.ShowExport,
		ShowImport:     in.ShowImport,
		RowSelection:   in.RowSelection,
		ShowBatchDel:   in.ShowBatchDel,
		BatchDelUri:    in.BatchDelUri,
		ExportUri:      in.ExportUri,
		ImportUri:      in.ImportUri,
		ImportMulti:    in.ImportMulti,
		ImportMaxCount: in.ImportMaxCount,
		ImportFormID:   in.ImportFormID,
	}
}

func PBToPageFields(in []*apipb.PageField) []*PageField {
	var list []*PageField
	for _, field := range in {
		list = append(list, &PageField{
			Model: commonmodel.Model{
				ID: field.Id,
			},
			PageID:         field.PageID,
			Name:           field.Name,
			Title:          field.Title,
			Copyable:       field.Copyable,
			Ellipsis:       field.Ellipsis,
			RowKey:         field.RowKey,
			Sort:           field.Sort,
			ShowInTable:    field.ShowInTable,
			ValueEnum:      field.ValueEnum,
			Component:      field.Component,
			ComponentProps: field.ComponentProps,
			DataType:       field.DataType,
			LabelField:     field.LabelField,
			ValueField:     field.ValueField,
			EnableSort:     field.EnableSort,
			Fixed:          field.Fixed,
			Width:          field.Width,
			Align:          field.Align,
		})
	}
	return list
}

func PageFieldsToPB(in []*PageField) []*apipb.PageField {
	var list []*apipb.PageField
	for _, field := range in {
		list = append(list, &apipb.PageField{
			Id:             field.ID,
			PageID:         field.PageID,
			Name:           field.Name,
			Title:          field.Title,
			Copyable:       field.Copyable,
			Ellipsis:       field.Ellipsis,
			RowKey:         field.RowKey,
			Sort:           field.Sort,
			ShowInTable:    field.ShowInTable,
			ValueEnum:      field.ValueEnum,
			Component:      field.Component,
			ComponentProps: field.ComponentProps,
			DataType:       field.DataType,
			LabelField:     field.LabelField,
			ValueField:     field.ValueField,
			EnableSort:     field.EnableSort,
			Fixed:          field.Fixed,
			Width:          field.Width,
			Align:          field.Align,
		})
	}
	return list
}

func PBToMetadata(in *apipb.MetadataInfo) *Metadata {
	return &Metadata{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Name:           in.Name,
		DisplayName:    in.DisplayName,
		Level:          in.Level,
		ParentID:       in.ParentID,
		Description:    in.Description,
		Package:        in.Package,
		TenantID:       in.TenantID,
		ProjectID:      in.ProjectID,
		MetadataFields: PBToMetadataFields(in.MetadataFields),
		IsMust:         in.IsMust,
	}
}

func MetadataToPB(in *Metadata) *apipb.MetadataInfo {
	return &apipb.MetadataInfo{
		Id:             in.ID,
		TenantID:       in.TenantID,
		ProjectID:      in.ProjectID,
		Name:           in.Name,
		DisplayName:    in.DisplayName,
		Level:          in.Level,
		ParentID:       in.ParentID,
		Description:    in.Description,
		Package:        in.Package,
		MetadataFields: MetadataFieldsToPB(in.MetadataFields),
		Children:       MetadatasToPB(in.Children),
		IsMust:         in.IsMust,
	}
}

func MetadatasToPB(in []*Metadata) []*apipb.MetadataInfo {
	var list []*apipb.MetadataInfo
	for _, f := range in {
		list = append(list, MetadataToPB(f))
	}
	return list
}

func PBToMetadataFields(in []*apipb.MetadataField) []*MetadataField {
	var list []*MetadataField
	for _, field := range in {
		list = append(list, &MetadataField{
			Model: commonmodel.Model{
				ID: field.Id,
			},
			MetadataID:   field.MetadataID,
			Name:         field.Name,
			Type:         field.Type,
			Length:       field.Length,
			NotNull:      field.NotNull,
			Comment:      field.Comment,
			IsArray:      field.IsArray,
			RefMetadata:  field.RefMetadata,
			DisplayName:  field.DisplayName,
			ShowInTable:  field.ShowInTable,
			ShowInEdit:   field.ShowInEdit,
			Component:    field.Component,
			Unique:       field.Unique,
			Index:        field.Index,
			DefaultValue: field.DefaultValue,
			ShowInQuery:  field.ShowInQuery,
			Order:        field.Order,
			Like:         field.Like,
			Copier:       field.Copier,
			DotNotGen:    field.DotNotGen,
			PBToStruct:   field.PbToStruct,
			StructToPB:   field.StructToPB,
		})
	}
	return list
}

func MetadataFieldsToPB(in []*MetadataField) []*apipb.MetadataField {
	var list []*apipb.MetadataField
	for _, field := range in {
		list = append(list, &apipb.MetadataField{
			Id:           field.ID,
			MetadataID:   field.MetadataID,
			Name:         field.Name,
			Type:         field.Type,
			Length:       field.Length,
			NotNull:      field.NotNull,
			Comment:      field.Comment,
			IsArray:      field.IsArray,
			RefMetadata:  field.RefMetadata,
			DisplayName:  field.DisplayName,
			ShowInTable:  field.ShowInTable,
			ShowInEdit:   field.ShowInEdit,
			Component:    field.Component,
			Unique:       field.Unique,
			Index:        field.Index,
			DefaultValue: field.DefaultValue,
			ShowInQuery:  field.ShowInQuery,
			Order:        field.Order,
			Like:         field.Like,
			Copier:       field.Copier,
			DotNotGen:    field.DotNotGen,
			PbToStruct:   field.PBToStruct,
			StructToPB:   field.StructToPB,
		})
	}
	return list
}
