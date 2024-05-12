package gen

import (
	"fmt"
	stdtpl "html/template"
	"strings"

	curdmodel "github.com/CloudSilk/curd/model"
)

/*
PRIMARY KEY (`id`),

	KEY `idx_api_group` (`group`),
	KEY `idx_api_method` (`method`),
	KEY `idx_api_description` (`description`),
	KEY `idx_api_enable` (`enable`),
	KEY `idx_api_check_auth` (`check_auth`),
	KEY `idx_api_deleted_at` (`deleted_at`),
	KEY `idx_api_uri` (`path`),
	KEY `idx_api_check_login` (`check_login`),
	KEY `idx_api_path` (`path`)
	UNIQUE KEY `unique_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
*/
func GenCreateSql(md *curdmodel.Metadata) stdtpl.HTML {
	var fields []string
	var indexs []string
	md.Sort()
	tableName := curdmodel.NamingStrategy.TableName(md.Name)
	for _, field := range md.MetadataFields {
		str := GenCreateColumnSql(field)
		if str != "" {
			fields = append(fields, "\n"+str)
		}

		str = GenCreateIndexsql(tableName, field)
		if str != "" {
			indexs = append(indexs, "\n"+str)
		}
	}
	uIndex := GenCreateUniqueIndexsql(md)
	if uIndex != "" {
		fields = append(fields, "\n"+uIndex)
	}
	if len(indexs) > 0 {
		fields = append(fields, indexs...)
	}
	return stdtpl.HTML(strings.Join(fields, ","))
}

func GenCreateIndexsql(tableName string, f *curdmodel.MetadataField) string {
	name := LowerSnakeCase(f.Name)
	if f.Index {
		return fmt.Sprintf("KEY `idx_%s_%s` (`%s`)", tableName, name, name)
	}
	return ""
}

func GenCreateUniqueIndexsql(md *curdmodel.Metadata) string {
	var fields []string
	for _, field := range md.MetadataFields {
		if field.Unique {
			fields = append(fields, "`"+LowerSnakeCase(field.Name)+"`")
		}
	}
	if len(fields) == 0 {
		return ""
	}
	return fmt.Sprintf("UNIQUE KEY `uidx_%s` (%s)", curdmodel.NamingStrategy.TableName(md.Name), strings.Join(fields, ","))
}

func GenCreateColumnSql(f *curdmodel.MetadataField) string {
	if f.Name == "ID" || f.Name == "id" {
		return "`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,PRIMARY KEY (`id`)"
	}
	notNull := "NULL"
	if f.NotNull {
		notNull = "NOT NULL"
	}
	defaultValue := f.DefaultValue
	if defaultValue == "" {
		defaultValue = "NULL"
	}
	name := LowerSnakeCase(f.Name)

	switch f.Type {
	case "datetime":
		return fmt.Sprintf("`%s` datetime(3) DEFAULT %s COMMENT '%s'", name, defaultValue, f.Comment)
	case "longtext":
		return fmt.Sprintf("`%s` longtext %s COMMENT '%s' DEFAULT %s", name, notNull, f.Comment, defaultValue)
	case "varchar", "string":
		size := f.Length
		if size == 0 {
			size = 100
		}
		return fmt.Sprintf("`%s` varchar(%d) %s COMMENT '%s' DEFAULT %s", name, size, notNull, f.Comment, defaultValue)
	case "tinyint", "bool":
		return fmt.Sprintf("`%s` tinyint(1) %s COMMENT '%s' DEFAULT %s", name, notNull, f.Comment, defaultValue)
	case "bigint":
		return fmt.Sprintf("`%s` bigint(20) %s COMMENT '%s' DEFAULT %s", name, notNull, f.Comment, defaultValue)
	case "int":
		return fmt.Sprintf("`%s` int(11) %s COMMENT '%s' DEFAULT %s", name, notNull, f.Comment, defaultValue)
	}
	return ""
}
