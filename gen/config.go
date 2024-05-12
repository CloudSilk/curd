package gen

import (
	"fmt"
	stdtpl "html/template"
	"strings"

	curdmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	ucpb "github.com/CloudSilk/usercenter/proto"
	"github.com/google/uuid"
)

const (
	apiConfigTpl = ``
)

type Config struct {
	APIs []ucpb.APIInfo
	Menu []ucpb.MenuInfo
}

func GenMenuConfig(service *curdmodel.Service, metadata *curdmodel.Metadata) Config {
	apiPrefix := strings.ToLower(service.Name)
	if service.Package != "" {
		apiPrefix = fmt.Sprintf("%s/%s", strings.ToLower(service.Package), strings.ToLower(service.Name))
	}
	addAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/add", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "POST",
		Description: "新增",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}
	updateAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/update", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "PUT",
		Description: "更新",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	deleteAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/delete", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "DELETE",
		Description: "删除",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	queryAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/query", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "GET",
		Description: "分页查询",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	enableAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/enable", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "POST",
		Description: "启用/禁用",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	allAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/all", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "GET",
		Description: "查询所有",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	detailAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/detail", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "GET",
		Description: "查询明细",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	copyAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/copy", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "POST",
		Description: "复制",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	exportAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/export", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "GET",
		Description: "导出",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	importAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/import", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "POST",
		Description: "导入",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	treeAPI := ucpb.APIInfo{
		Id:          uuid.New().String(),
		Path:        fmt.Sprintf("/api/%s/%s/tree", apiPrefix, strings.ToLower(metadata.Name)),
		Group:       metadata.Name,
		Method:      "GET",
		Description: "查询(Tree)",
		Enable:      true,
		CheckAuth:   true,
		CheckLogin:  true,
		TenantID:    service.TenantID,
		ProjectID:   service.ProjectID,
	}

	menu := ucpb.MenuInfo{
		Id:        uuid.New().String(),
		Path:      fmt.Sprintf("/curd/page/manager/%s?pageName=%s", CamelName(metadata.Name), CamelName(metadata.Name)),
		Name:      metadata.DisplayName,
		Title:     metadata.DisplayName,
		TenantID:  service.TenantID,
		ProjectID: service.ProjectID,
	}
	addFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Add%s", CamelName(metadata.Name)),
		Title:  "新增",
	}
	addFunc.MenuFuncApis = append(addFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: addFunc.Id,
			ApiID:      addAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &addFunc)

	updateFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Update%s", CamelName(metadata.Name)),
		Title:  "更新",
	}
	updateFunc.MenuFuncApis = append(updateFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: updateFunc.Id,
			ApiID:      updateAPI.Id,
		},
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: updateFunc.Id,
			ApiID:      detailAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &updateFunc)

	deleteFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Delete%s", CamelName(metadata.Name)),
		Title:  "删除",
	}
	deleteFunc.MenuFuncApis = append(deleteFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: deleteFunc.Id,
			ApiID:      deleteAPI.Id,
		},
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: deleteFunc.Id,
			ApiID:      detailAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &deleteFunc)

	queryFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Query%s", CamelName(metadata.Name)),
		Title:  "查询",
	}
	queryFunc.MenuFuncApis = append(queryFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: queryFunc.Id,
			ApiID:      queryAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &queryFunc)

	allFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("GetAll%s", CamelName(metadata.Name)),
		Title:  "查询所有",
	}
	allFunc.MenuFuncApis = append(allFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: allFunc.Id,
			ApiID:      allAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &allFunc)

	detailFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("GetDetail%s", CamelName(metadata.Name)),
		Title:  "查询明细",
	}
	detailFunc.MenuFuncApis = append(detailFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: addFunc.Id,
			ApiID:      detailAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &detailFunc)

	copyFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Copy%s", CamelName(metadata.Name)),
		Title:  "复制",
	}
	copyFunc.MenuFuncApis = append(copyFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: copyFunc.Id,
			ApiID:      copyAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &copyFunc)

	enableFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Enable%s", CamelName(metadata.Name)),
		Title:  "启用/禁用",
	}
	enableFunc.MenuFuncApis = append(enableFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: enableFunc.Id,
			ApiID:      enableAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &enableFunc)

	exportFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Export%s", CamelName(metadata.Name)),
		Title:  "导出",
	}
	exportFunc.MenuFuncApis = append(exportFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: exportFunc.Id,
			ApiID:      exportAPI.Id,
		},
	)
	menu.MenuFuncs = append(menu.MenuFuncs, &exportFunc)

	importFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Import%s", CamelName(metadata.Name)),
		Title:  "导入",
	}
	importFunc.MenuFuncApis = append(importFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: importFunc.Id,
			ApiID:      importAPI.Id,
		},
	)

	treeFunc := ucpb.MenuFunc{
		Id:     uuid.New().String(),
		MenuID: menu.Id,
		Name:   fmt.Sprintf("Get%sTree", CamelName(metadata.Name)),
		Title:  "查询(Tree)",
	}
	treeFunc.MenuFuncApis = append(treeFunc.MenuFuncApis,
		&ucpb.MenuFuncApi{
			Id:         uuid.New().String(),
			MenuFuncID: treeFunc.Id,
			ApiID:      treeAPI.Id,
		},
	)

	menu.MenuFuncs = append(menu.MenuFuncs, &importFunc)
	config := Config{}
	config.Menu = append(config.Menu, menu)
	config.APIs = append(config.APIs, addAPI, updateAPI, deleteAPI, queryAPI, enableAPI, allAPI, detailAPI, copyAPI, exportAPI, importAPI, treeAPI)
	return config
}

func genPageConfig(service *curdmodel.Service, metadata *curdmodel.Metadata) stdtpl.HTML {
	prefix := strings.ToLower(service.Name)
	if service.Package != "" {
		prefix = fmt.Sprintf("%s/%s", strings.ToLower(service.Package), strings.ToLower(service.Name))
	}
	page := &apipb.PageInfo{
		Id:            uuid.NewString(),
		Name:          metadata.Name,
		Enable:        true,
		MetadataID:    metadata.ID,
		PageSize:      15,
		ShowSelection: true,
		Path:          fmt.Sprintf("%s/%s", prefix, strings.ToLower(metadata.Name)),
		Title:         metadata.DisplayName,
		Description:   metadata.DisplayName,
		Type:          1,
		TenantID:      service.TenantID,
		ProjectID:     service.ProjectID,
	}
	page.ToolBar = &apipb.PageToolBar{
		Id:           uuid.NewString(),
		PageID:       page.Id,
		FullScreen:   true,
		Reload:       true,
		Setting:      true,
		ShowAdd:      true,
		ShowExport:   true,
		ShowImport:   true,
		ShowBatchDel: true,
	}
	page.Buttons = append(page.Buttons,
		&apipb.PageButton{
			Id:           uuid.NewString(),
			PageID:       page.Id,
			Key:          "editable",
			Label:        "编辑",
			Script:       "curdPage.showEditDialog(record)",
			Enable:       true,
			ShowType:     "Dialog",
			Index:        1,
			ShowPosition: 1,
		},
		&apipb.PageButton{
			Id:           uuid.NewString(),
			PageID:       page.Id,
			Key:          "view",
			Label:        "查看",
			Script:       "",
			Enable:       true,
			ShowType:     "Page",
			Index:        2,
			ShowPosition: 1,
		},
		&apipb.PageButton{
			Id:           uuid.NewString(),
			PageID:       page.Id,
			Key:          "copy",
			Label:        "复制",
			Script:       "",
			Enable:       true,
			Expanded:     true,
			ShowType:     "",
			Index:        3,
			ShowPosition: 1,
		},
		&apipb.PageButton{
			Id:           uuid.NewString(),
			PageID:       page.Id,
			Key:          "enable",
			Label:        "启用",
			Script:       "",
			Enable:       true,
			Expanded:     true,
			ShowType:     "",
			Index:        4,
			ShowPosition: 1,
		},
		&apipb.PageButton{
			Id:           uuid.NewString(),
			PageID:       page.Id,
			Key:          "disable",
			Label:        "禁用",
			Script:       "",
			Enable:       true,
			Expanded:     true,
			ShowType:     "",
			Index:        5,
			ShowPosition: 1,
		},
		&apipb.PageButton{
			Id:           uuid.NewString(),
			PageID:       page.Id,
			Key:          "delete",
			Label:        "删除",
			Script:       "",
			Enable:       true,
			Expanded:     true,
			ShowType:     "",
			Index:        6,
			ShowPosition: 1,
		},
	)
	index := int32(1)
	for _, field := range metadata.MetadataFields {
		if !field.ShowInTable {
			continue
		}
		page.Fields = append(page.Fields, &apipb.PageField{
			Id:          uuid.NewString(),
			PageID:      page.Id,
			Name:        field.Name,
			Title:       field.DisplayName,
			Sort:        index,
			ShowInTable: true,
			DataType:    "Text",
		})
		index++
	}
	return stdtpl.HTML(JsonMarshal([]*apipb.PageInfo{page}))
}

func GenPageConfig(ctx *GenContext) stdtpl.HTML {
	return genPageConfig(ctx.Service, ctx.Metadata)
}
