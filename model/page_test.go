package model

import "testing"

func TestAddPage(t *testing.T) {
	p := &Page{
		Name:       "PageConfig",
		Enable:     true,
		MetadataID: "59",
		PageSize:   10,
		ShowIndex:  true,
		ToolBar: &PageToolBar{
			FullScreen: true,
			Reload:     true,
			Setting:    true,
		},
		Fields: []*PageField{
			{
				Name:     "name",
				Title:    "页面名称",
				Copyable: true,
				Ellipsis: false,
				RowKey:   true,
			},
			{
				Name:  "enable",
				Title: "是否启用",
			},
		},
	}
	err := CreatePage(p)
	if err != nil {
		t.Fatal(err)
	}
}
