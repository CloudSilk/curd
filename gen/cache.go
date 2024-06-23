package gen

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/CloudSilk/curd/model"
	curdmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/utils/log"
)

var (
	templateCache     = sync.Map{}
	metadataCache     = sync.Map{}
	metadataNameCache = sync.Map{}
)

func LoadCache() {
	tpls, err := curdmodel.GetAllTemplates(&apipb.QueryTemplateRequest{})
	if err != nil {
		panic(err)
	}
	for _, tpl := range tpls {
		templateCache.Store(tpl.ID, tpl)
	}

	metadatas, err := curdmodel.GetAllMetadatas(&apipb.QueryMetadataRequest{})
	if err != nil {
		panic(err)
	}
	for _, md := range metadatas {
		md.Sort()
		metadataCache.Store(md.ID, md)
		metadataNameCache.Store(md.Name, md)
	}

	for _, md := range metadatas {
		InitGenCode(md)
	}
}

func GetTemplateById(tplID string) (*curdmodel.Template, error) {
	tpl, ok := templateCache.Load(tplID)
	if !ok {
		return nil, fmt.Errorf("代码模板(%s)不存在", tplID)
	}
	result := tpl.(*curdmodel.Template)
	return result, nil
}

func GetMetadataById(mdID string) (*curdmodel.Metadata, error) {
	md, ok := metadataCache.Load(mdID)
	if !ok {
		md, err := model.GetMetadataById(mdID)
		if err != nil {
			return nil, err
		}
		md.Sort()
		metadataCache.Store(md.ID, md)
		metadataNameCache.Store(md.Name, md)
		InitGenCode(md)
		return md, nil
	}
	result := md.(*curdmodel.Metadata)
	return result, nil
}

func GetMetaDataNameByID(mdID string) (string, error) {
	md, err := GetMetadataById(mdID)
	if err != nil {
		return "", err
	}
	return md.Name, nil
}

func InitGenCode(md *curdmodel.Metadata) {
	var uniqueFields []string
	var fields []string
	md.Preloads = []string{}
	md.Sort()
	for _, field := range md.MetadataFields {
		if field.Unique {
			uniqueFields = append(uniqueFields, " "+LowerSnakeCase(field.Name)+" =? ")
			fields = append(fields, "m."+CamelName(field.Name))
		}
	}

	if len(uniqueFields) == 0 {
		md.UniqueFields, md.Fields = " name=? ", "m.Name"
	} else {
		md.UniqueFields, md.Fields = strings.Join(uniqueFields, " and "), strings.Join(fields, ", ")
	}

	//嵌套查找RefMetadata是否也有RefMetadata
	for _, field := range md.MetadataFields {
		preload := RecursiveRefMetadata(md, field, "")
		if len(preload) > 0 {
			md.Preloads = append(md.Preloads, preload...)
		}
	}

}

func RecursiveRefMetadatas(parentMD *curdmodel.Metadata, fields []*curdmodel.MetadataField, prefix string) []string {
	var result []string
	for _, field := range fields {
		list := RecursiveRefMetadata(parentMD, field, CamelName(field.Name))
		for _, s := range list {
			if prefix != "" {
				result = append(result, prefix+"."+s)
			} else {
				result = append(result, s)
			}
		}
	}
	return result
}

func RecursiveRefMetadata(parentMD *curdmodel.Metadata, field *curdmodel.MetadataField, prefix string) []string {
	if field.RefMetadata != "" {
		refMetadata, err := GetMetadataById(field.RefMetadata)
		if err != nil {
			log.Errorf(context.Background(), "GetMetadataById error:%v", err)
			return []string{}
		}
		if parentMD.ID == refMetadata.ID {
			return []string{}
		}
		result := RecursiveRefMetadatas(refMetadata, refMetadata.MetadataFields, CamelName(field.Name))
		if len(result) == 0 {
			result = append(result, CamelName(field.Name))
		}
		return result
	}
	return []string{}
}

func GetMetadataByName(name string) (*curdmodel.Metadata, error) {
	md, ok := metadataNameCache.Load(name)
	if !ok {
		return nil, errors.New(fmt.Sprintf("元数据(%s)不存在", name))
	}
	result := md.(*curdmodel.Metadata)
	return result, nil
}

func AddMetadata(md *curdmodel.Metadata) {
	md.Sort()
	InitGenCode(md)
	metadataCache.Store(md.ID, md)
	metadataNameCache.Store(md.Name, md)
}

func DeleteMetadata(id string) {
	value, err := GetMetadataById(id)
	if err != nil {
		metadataCache.Delete(id)
		metadataNameCache.Delete(value.Name)
	}

}

func AddTemplate(tpl *curdmodel.Template) {
	templateCache.Store(tpl.ID, tpl)
}

func DeleteTemplate(id string) {
	templateCache.Delete(id)
}
