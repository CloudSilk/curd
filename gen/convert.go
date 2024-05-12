package gen

import (
	"encoding/json"

	curdmodel "github.com/CloudSilk/curd/model"
)

func ConvertTSType(f *curdmodel.MetadataField) string {
	if f.RefMetadata != "" {
		md, _ := GetMetadataById(f.RefMetadata)
		if f.IsArray {
			return CamelName(md.Name) + "[]"
		}
		return CamelName(md.Name)
	}
	switch f.Type {
	case "bigint", "int":
		return "number"
	case "varchar", "longtext":
		return "string"
	case "datetime":
		return "string"
	case "tinyint", "bool":
		return "boolean"
	default:
		return f.Type
	}
}

func JsonMarshal(obj interface{}) string {
	buf, _ := json.Marshal(obj)
	return string(buf)
}
