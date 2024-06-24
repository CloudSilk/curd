package gen

import (
	"context"
	"fmt"
	stdtpl "html/template"
	"strings"

	curdmodel "github.com/CloudSilk/curd/model"
	"github.com/CloudSilk/pkg/utils/log"
)

func GenStructField(f *curdmodel.MetadataField, m *curdmodel.Metadata) stdtpl.HTML {
	if f.Name == "UpdatedAt" || f.Name == "CreatedAt" || f.Name == "DeletedAt" || f.Name == "ID" {
		return ""
	}
	var list []string
	if f.DotNotGen {
		list = append(list, "-")
	} else if f.RefMetadata != "" {
		list = append(list, "")
	} else {
		if (f.Type == "varchar" || f.Type == "nvarchar" || f.Type == "string") && f.Type != "nvarchar(max)" {
			list = append(list, fmt.Sprintf("size:%d", f.Length))
		}
		if f.Index {
			list = append(list, "index")
		}
		if f.Comment != "" {
			list = append(list, fmt.Sprintf("comment:%s", f.Comment))
		}
		if f.Unique {
			//一索引的名称加上表名:因为sqlite里面唯一索引必须全局唯一
			list = append(list, fmt.Sprintf("uniqueindex:%s_uidx1", m.Name))
		}
	}

	notCopy := ""
	if !f.Copier {
		notCopy = `copier:"-"`
	}
	return stdtpl.HTML(fmt.Sprintf("\n\t//%s %s\n\t%s %s `json:\"%s\" gorm:\"%s\" %s`", f.DisplayName, f.Comment, CamelName(f.Name), ConvertGoType(f), LcFirst(f.Name), strings.Join(list, ";"), notCopy))
}

func GenQueryStructField(f *curdmodel.MetadataField) stdtpl.HTML {
	lcName := LcFirst(f.Name)
	return stdtpl.HTML(fmt.Sprintf("\n\t%s %s `json:\"%s\" form:\"%s\" uri:\"%s\"`", f.Name, ConvertGoType(f), lcName, lcName, lcName))
}

func GenQueryCond(f *curdmodel.MetadataField) stdtpl.HTML {
	t := ConvertGoType(f)
	switch {
	case !f.ShowInQuery:
		return stdtpl.HTML("")
	case f.Like:
		return stdtpl.HTML(fmt.Sprintf(`
		if req.%s != "" {
			db = db.Where("%s LIKE ?", "%%"+req.%s+"%%")
		}
		`, CamelName3(f.Name), LowerSnakeCase(f.Name), CamelName3(f.Name)))
	case t == "uint" || t == "int" || t == "int32" || t == "int64" || t == "float32" || t == "float64" || t == "bigint":
		return stdtpl.HTML(fmt.Sprintf(`
		if req.%s >0 {
			db = db.Where("%s = ?",req.%s)
		}
		`, CamelName3(f.Name), LowerSnakeCase(f.Name), CamelName3(f.Name)))
	case t == "bool":
		return stdtpl.HTML(fmt.Sprintf("\n\tdb = db.Where(\"%s = ?\",req.%s)", LowerSnakeCase(f.Name), CamelName3(f.Name)))
	default:
		return stdtpl.HTML(fmt.Sprintf(`
		if req.%s !="" {
			db = db.Where("%s = ?",req.%s)
		}
		`, CamelName3(f.Name), LowerSnakeCase(f.Name), CamelName3(f.Name)))
	}
}

func GenPBToStrcut(f *curdmodel.MetadataField, str string) stdtpl.HTML {
	name := strings.ToLower(f.Name)
	if name == "updatedat" || name == "createdat" || name == "deletedat" {
		return ""
	}
	if name == "id" {
		return `
		Model: commonmodel.Model{
			ID: in.Id,
		},`
	}
	name = CamelName(f.Name)
	if f.RefMetadata != "" {
		md, _ := GetMetadataById(f.RefMetadata)
		if f.IsArray {
			return stdtpl.HTML(fmt.Sprintf("\n	%s: PBTo%ss(in.%s),", name, CamelName(md.Name), name))
		}
		return stdtpl.HTML(fmt.Sprintf("\n	%s: PBTo%s(in.%s),", name, CamelName(md.Name), name))
	}

	if f.Type == "date" && !f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.ParseSqlNullDate(in.%s),", name, name))
	} else if f.Type == "date" && f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.ParseDate(in.%s),", name, name))
	}

	if f.Type == "datetime" && !f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.ParseSqlNullTime(in.%s),", name, name))
	} else if f.Type == "datetime" && f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.ParseTime(in.%s),", name, name))
	}

	if f.PBToStruct != "" {
		return stdtpl.HTML(fmt.Sprintf("\n	%s: %s(in.%s),", name, f.PBToStruct, name))
	}

	return stdtpl.HTML(fmt.Sprintf("\n	%s: in.%s,", name, name))
}

func GenStrcutToPB(f *curdmodel.MetadataField, str string) stdtpl.HTML {
	name := strings.ToLower(f.Name)
	if name == "updatedat" || name == "createdat" || name == "deletedat" {
		return ""
	}
	if name == "id" {
		return "\n    Id: in.ID,"
	}
	name = CamelName(f.Name)
	if f.RefMetadata != "" {
		md, _ := GetMetadataById(f.RefMetadata)
		if f.IsArray {
			return stdtpl.HTML(fmt.Sprintf("\n	%s: %ssToPB(in.%s),", name, CamelName(md.Name), name))
		}
		return stdtpl.HTML(fmt.Sprintf("\n	%s: %sToPB(in.%s),", name, CamelName(md.Name), name))
	}

	if f.Type == "date" && !f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.FormatSqlNullDate(in.%s),", name, name))
	} else if f.Type == "date" && f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.FormatDate(in.%s),", name, name))
	}

	if f.Type == "datetime" && !f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.FormatSqlNullTime(in.%s),", name, name))
	} else if f.Type == "datetime" && f.NotNull {
		return stdtpl.HTML(fmt.Sprintf("\n	%s:utils.FormatTime(in.%s),", name, name))
	}

	if f.StructToPB != "" {
		return stdtpl.HTML(fmt.Sprintf("\n	%s: %s(in.%s),", name, f.StructToPB, name))
	}

	return stdtpl.HTML(fmt.Sprintf("\n	%s: in.%s,", name, name))
}

func ConvertGoType(f *curdmodel.MetadataField) string {
	if f.RefMetadata != "" {
		md, _ := GetMetadataById(f.RefMetadata)
		if f.IsArray {
			return "[]*" + CamelName(md.Name)
		}
		return "*" + CamelName(md.Name)
	}
	switch f.Type {
	case "bigint":
		return "int64"
	case "varchar", "longtext", "nvarchar", "nvarchar(max)":
		return "string"
	case "datetime", "date":
		if f.NotNull {
			return "time.Time"
		}
		return "sql.NullTime"
	case "tinyint", "bit":
		return "bool"
	case "int32", "int":
		return "int32"
	default:
		return f.Type
	}
}

const deleteChildrenTpl = `
func Delete{{fieldStructName}}s(tx *gorm.DB, old, m *{{name}}) error {
	var deleteIDs []string
	for _, oldObj := range old.{{fieldName}} {
		flag := false
		for _, newObj := range m.{{fieldName}} {
			if newObj.ID == oldObj.ID {
				flag = true
				{{deleteChildren}}
			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&{{fieldStructName}}{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}
`
const deleteChildrenTpl2 = `
			err := Delete{{fieldName}}s(tx, oldObj, newObj)
			if err != nil {
				return err
			}
`

func GenDeleteChildren(md *curdmodel.Metadata) stdtpl.HTML {
	result := ""
	for _, field := range md.MetadataFields {
		if field.RefMetadata != "" {
			childMD, _ := GetMetadataById(field.RefMetadata)
			if childMD.ParentID != md.ID {
				continue
			}
			fieldStructName := CamelName(childMD.Name)
			deleteChildren := GenDeleteChildren(childMD)
			tpl := strings.Replace(deleteChildrenTpl, "{{fieldName}}", CamelName(field.Name), 3)
			tpl = strings.Replace(tpl, "{{name}}", md.Name, 1)
			tpl = strings.Replace(tpl, "{{fieldStructName}}", fieldStructName, 2)
			tpl2 := GenDeleteChildren2(childMD)
			tpl = strings.Replace(tpl, "{{deleteChildren}}", string(tpl2), 1)

			result += "\n" + string(deleteChildren)
			result += "\n" + tpl
		}
	}
	return stdtpl.HTML(result)
}

func GenDeleteChildren2(md *curdmodel.Metadata) stdtpl.HTML {
	result := ""
	for _, field := range md.MetadataFields {
		if field.RefMetadata != "" {
			childMD, _ := GetMetadataById(field.RefMetadata)
			if childMD.ParentID != md.ID {
				continue
			}
			result += "\n" + strings.Replace(deleteChildrenTpl2, "{{fieldName}}", CamelName(childMD.Name), 1)
		}
	}
	return stdtpl.HTML(result)
}

const deleteChildrenByParentIDTpl = `
func Delete{{fieldName}}ByParent(tx *gorm.DB, m *{{name}}) error {
	var deleteIDs []string
	for _, newObj := range m.{{fieldName}} {
		deleteIDs = append(deleteIDs, newObj.ID)
		{{deleteChildren}}
	}

	if len(deleteIDs) > 0 {
		err := tx.Unscoped().Delete(&{{fieldStructName}}{}, "id in ?", deleteIDs).Error
		if err != nil {
			return err
		}
	}
	return nil
}
`
const deleteChildrenByParentIDTpl2 = `
			err := Delete{{fieldName}}sByParent(tx, newObj)
			if err != nil {
				return err
			}
`

func GenDeleteChildrenByParentID(md *curdmodel.Metadata) stdtpl.HTML {
	result := ""
	for _, field := range md.MetadataFields {
		if field.RefMetadata != "" {
			childMD, _ := GetMetadataById(field.RefMetadata)
			if childMD.ParentID != md.ID {
				continue
			}
			fieldStructName := CamelName(childMD.Name)
			deleteChildren := GenDeleteChildrenByParentID(childMD)
			tpl := strings.Replace(deleteChildrenByParentIDTpl, "{{fieldName}}", CamelName(field.Name), 3)
			tpl = strings.Replace(tpl, "{{name}}", md.Name, 1)
			tpl = strings.Replace(tpl, "{{fieldStructName}}", fieldStructName, 1)

			tpl2 := GenDeleteChildrenByParentID2(childMD)
			tpl = strings.Replace(tpl, "{{deleteChildren}}", string(tpl2), 1)

			result += "\n" + string(deleteChildren)
			result += "\n" + tpl
		}
	}
	return stdtpl.HTML(result)
}

func GenDeleteChildrenByParentID2(md *curdmodel.Metadata) stdtpl.HTML {
	result := ""
	for _, field := range md.MetadataFields {
		if field.RefMetadata != "" {
			childMD, _ := GetMetadataById(field.RefMetadata)
			if childMD.ParentID != md.ID {
				continue
			}
			result += "\n" + strings.Replace(deleteChildrenByParentIDTpl2, "{{fieldName}}", CamelName(childMD.Name), 1)
		}
	}
	return stdtpl.HTML(result)
}

func GenSwagQueryParam(f *curdmodel.MetadataField) stdtpl.HTML {
	if !f.ShowInQuery {
		return ""
	}
	t := "string"
	switch f.Type {
	case "bigint":
		t = "int64"
	case "varchar", "longtext", "string":
		t = "string"
	case "datetime":
		t = "string"
	case "tinyint":
		t = "bool"
	case "int32", "int":
		t = "int"
	}
	return stdtpl.HTML(fmt.Sprintf("\n// @Param %s query %s false \"%s\"", f.Name, t, f.DisplayName))
}

func RecursiveGetRefMetadatas(md *curdmodel.Metadata, mustChildren bool) []*curdmodel.Metadata {
	refMetadatas := make(map[string]*curdmodel.Metadata)
	err := recursiveGetRefMetadatas(md, md.MetadataFields, refMetadatas, mustChildren)
	if err != nil {
		return []*curdmodel.Metadata{}
	}
	var result []*curdmodel.Metadata
	for _, refMetadata := range refMetadatas {
		result = append(result, refMetadata)
	}
	return result
}

func recursiveGetRefMetadatas(parentMD *curdmodel.Metadata, fields []*curdmodel.MetadataField, result map[string]*curdmodel.Metadata, mustChildren bool) error {
	for _, field := range fields {
		if field.RefMetadata == "" {
			continue
		}

		refMetadata, err := GetMetadataById(field.RefMetadata)
		if err != nil {
			log.Errorf(context.Background(), "GetMetadataById error:%v", err)
			return err
		}
		if parentMD.ID == refMetadata.ID {
			continue
		}
		if mustChildren && refMetadata.ParentID != parentMD.ID {
			continue
		}
		result[refMetadata.ID] = refMetadata
		err = recursiveGetRefMetadatas(refMetadata, refMetadata.MetadataFields, result, mustChildren)
		if err != nil {
			return err
		}
	}
	return nil
}

func RecursiveGetRefMetadatas2(md *curdmodel.Metadata) []*curdmodel.Metadata {
	var result []*curdmodel.Metadata
	for _, field := range md.MetadataFields {
		if field.RefMetadata == "" {
			continue
		}
		if md.ID == field.RefMetadata {
			continue
		}
		refMetadata, err := GetMetadataById(field.RefMetadata)
		if err != nil {
			log.Errorf(context.Background(), "GetMetadataById error:%v", err)
			return []*curdmodel.Metadata{}
		}
		result = append(result, refMetadata)
	}
	return result
}
