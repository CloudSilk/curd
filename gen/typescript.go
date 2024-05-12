package gen

import (
	"fmt"
	stdtpl "html/template"
	"strings"

	curdmodel "github.com/CloudSilk/curd/model"
)

func GenClassField(f *curdmodel.MetadataField) stdtpl.HTML {
	name := strings.ToLower(f.Name)
	if name == "updatedat" || name == "createdat" || name == "deletedat" || name == "id" {
		return ""
	}
	return stdtpl.HTML(fmt.Sprintf("\n\t//%s %s\n\t%s: %s", f.DisplayName, f.Comment, CamelName2(f.Name), ConvertTSType(f)))
}

func GenClassDefaultValue(f *curdmodel.MetadataField) stdtpl.HTML {
	name := strings.ToLower(f.Name)
	if name == "updatedat" || name == "createdat" || name == "deletedat" || name == "id" {
		return ""
	}
	if f.IsArray {
		return stdtpl.HTML(fmt.Sprintf("\n\t    %s: [],", CamelName2(f.Name)))
	}

	return stdtpl.HTML(fmt.Sprintf("\n\t     %s: %s,", CamelName2(f.Name), GetDefaultValue(f)))
}

func GetDefaultValue(f *curdmodel.MetadataField) string {
	if f.RefMetadata != "" {
		return "{} as " + f.RefMetadata
	}
	switch f.Type {
	case "varchar", "longtext", "string":
		return fmt.Sprintf("'%s'", f.DefaultValue)
	case "bigint", "int":
		if f.DefaultValue != "" {
			return f.DefaultValue
		}
		return "0"
	case "bool", "tinyint":
		if f.DefaultValue != "" {
			return f.DefaultValue
		}
		return "false"
	default:
		return f.DefaultValue
	}
}
