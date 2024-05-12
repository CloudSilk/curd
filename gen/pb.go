package gen

import (
	"fmt"
	stdtpl "html/template"
	"strings"

	curdmodel "github.com/CloudSilk/curd/model"
)

func GenPBMessageProperty(index int, f *curdmodel.MetadataField, isOnly int) stdtpl.HTML {
	name := strings.ToLower(f.Name)
	if name == "updatedat" || name == "createdat" || name == "deletedat" {
		return ""
	}
	return stdtpl.HTML(fmt.Sprintf("\n\t//%s %s\n 	%s %s=%d;", f.DisplayName, f.Comment, ConvertPBType(f, isOnly), CamelName2(f.Name), index+2))
}

func ConvertPBType(f *curdmodel.MetadataField, isOnly int) string {
	if f.RefMetadata != "" {
		md, _ := GetMetadataById(f.RefMetadata)
		name := CamelName(md.Name)
		if isOnly == 1 {
			name += "Info"
		}
		if f.IsArray {
			return "repeated " + name
		}
		return name
	}
	switch f.Type {
	case "int":
		return "int32"
	case "bigint":
		return "int64"
	case "varchar", "longtext":
		return "string"
	case "datetime", "date":
		return "string"
	case "tinyint":
		return "bool"
	case "float32":
		return "float"
	case "float64":
		return "double"
	default:
		return f.Type
	}
}

func GenPBQueryCond(index int, f *curdmodel.MetadataField) stdtpl.HTML {
	t := ConvertPBType(f, 0)
	switch {
	case !f.ShowInQuery:
		return stdtpl.HTML("")
	default:
		name := CamelName2(f.Name)
		return stdtpl.HTML(fmt.Sprintf("\n\t//%s %s\n    // @inject_tag: uri:\"%s\" form:\"%s\"\n    %s %s=%d;", f.DisplayName, f.Comment, name, name, t, name, index+4))
	}
}
