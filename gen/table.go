package gen

import (
	"fmt"
	"strconv"

	curdmodel "github.com/CloudSilk/curd/model"
	curdpb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/db"
	"gorm.io/gorm"
)

type ColumnInfo struct {
	TableSchema            string `gorm:"column(table_schema)"`
	TableName              string `gorm:"column(table_name)"`
	ColumnName             string `gorm:"column(column_name)"`
	ColumnDefault          string `gorm:"column(column_default)"`
	IsNullable             string `gorm:"column(is_nullable)"`
	DataType               string `gorm:"column(data_type)"`
	ColumnKey              string `gorm:"column(column_key)"`
	ColumnComment          string `gorm:"column(column_comment)"`
	CharacterMaximumLength string `gorm:"column(character_maximum_length)"`
}

type TableInfo struct {
	TableName    string `gorm:"column(table_name)"`
	TableComment string `gorm:"column(table_comment)"`
}

func InitMetadata(db *gorm.DB, tableSchema string, isPostgres bool) {
	var tables []*TableInfo
	sqlStr := `select table_name as table_name,table_comment as table_comment from information_schema.tables  where table_schema = ?`
	if isPostgres {
		sqlStr = `select table_name as table_name from information_schema.tables  where table_catalog = ? and table_schema='public'`
	}
	//
	err := db.Raw(sqlStr, tableSchema).Find(&tables).Error
	if err != nil {
		panic(err)
	}
	// fmt.Println(tables)
	var mds []*curdpb.MetadataInfo
	for _, table := range tables {
		md := CreateMetadataByTable(db, "CompaAI", tableSchema, table.TableName, isPostgres)
		if md != nil {
			mds = append(mds, md)
		}
	}
	fmt.Println(JsonMarshal(mds))
}

func InitMetadataFromSqlServer(db db.DBClientInterface) {
	var tables []*TableInfo
	err := db.DB().Raw(`Select Name as table_name FROM SysObjects Where XType='U' orDER BY Name`).Find(&tables).Error
	if err != nil {
		panic(err)
	}
	// fmt.Println(tables)
	var mds []*curdpb.MetadataInfo
	for _, table := range tables {
		if table.TableName == "Role" || table.TableName == "User" || table.TableName == "MigrationHistory" ||
			table.TableName == "Menu" || table.TableName == "OnlineUser" || table.TableName == "Power" || table.TableName == "RolePower" || table.TableName == "RoleUser" {
			continue
		}
		md := CreateMetadataByTableFromSqlServer(db.DB(), "MES", "XZMES", table.TableName)
		if md != nil {
			mds = append(mds, md)
		}
	}
	// result := JsonMarshal(mds)
	// os.WriteFile("md.json", []byte(result), fs.ModePerm)
}

func CreateMetadataByTableFromSqlServer(sourceDB *gorm.DB, system, tableSchema, tableName string) *curdpb.MetadataInfo {
	var columns []*ColumnInfo
	err := sourceDB.Raw(`SELECT d.name as table_name,a.name AS column_name,a.length AS character_maximum_length
    , CASE
        WHEN (
            SELECT COUNT(*)
            FROM sysobjects
            WHERE name IN (
                    SELECT name
                    FROM sysindexes
                    WHERE id = a.id
                        AND indid IN (
                            SELECT indid
                            FROM sysindexkeys
                            WHERE id = a.id
                                AND colid IN (
                                    SELECT colid
                                    FROM syscolumns
                                    WHERE id = a.id
                                        AND name = a.name
                                )
                        )
                )
                AND xtype = 'PK'
        ) > 0 THEN 'YES'
        ELSE 'NO'
    END AS column_key, b.name AS data_type
    , CASE
        WHEN a.isnullable = 0 THEN 'YES'
        ELSE 'NO'
    END AS is_nullable
    , isnull(g.[value], '') AS column_comment,e.text as column_default
FROM syscolumns a
    LEFT JOIN systypes b ON a.xtype = b.xusertype
    INNER JOIN sysobjects d
    ON a.id = d.id
        AND d.xtype = 'U'
        AND d.name <> 'dtproperties'
    LEFT JOIN syscomments e on a.cdefault=e.id 
 
 left join sys.extended_properties g 
 on a.id=g.major_id AND a.colid= g.minor_id  
 where d.name=?
 order by a.id,a.colorder`, tableName).Scan(&columns).Error
	if err != nil {
		panic(err)
	}
	md := &curdpb.MetadataInfo{
		Name:        RemoveLastChar(CamelName(tableName)),
		DisplayName: RemoveLastChar(CamelName(tableName)),
		System:      system,
		ParentID:    "7172e7b2-c8fb-4d4d-8122-8a548939f15e",
		ProjectID:   "5a5a0292-deb0-4eb3-8627-b30c59c7a5da",
		TenantID:    "73e2a57a-109b-4a9e-8801-5826316cbe8f",
		Package:     "model",
		Level:       2,
	}
	if len(columns) == 0 {
		return nil
	}
	for i, col := range columns {
		// fmt.Println(CamelName2(col.ColumnName))
		length, _ := strconv.Atoi(col.CharacterMaximumLength)
		md.MetadataFields = append(md.MetadataFields, &curdpb.MetadataField{
			Name:        CamelName2(col.ColumnName),
			Type:        col.DataType,
			Length:      int32(length),
			NotNull:     col.IsNullable == "NO",
			Comment:     col.ColumnComment,
			DisplayName: CamelName(col.ColumnName),
			Order:       int32(i) + 1,
		})
	}

	err = curdmodel.CreateMetadata(curdmodel.PBToMetadata(md))
	if err != nil {
		fmt.Println(err)

	}
	return md
}

func CreateMetadataByTable(sourceDB *gorm.DB, system, tableSchema, tableName string, isPostgres bool) *curdpb.MetadataInfo {
	var columns []*ColumnInfo
	sqlStr := `select table_schema as table_schema,table_name as table_name,column_name as column_name,column_default as column_default,is_nullable as is_nullable,data_type as data_type,column_key as column_key,column_comment as column_comment,character_maximum_length as character_maximum_length from information_schema.columns where table_schema =? and table_name = ?`
	if isPostgres {
		sqlStr = `select table_schema as table_schema,table_name as table_name,column_name as column_name,column_default as column_default,is_nullable as is_nullable,data_type as data_type,character_maximum_length as character_maximum_length from information_schema.columns where table_schema='public' and table_name = ?`
	}
	var err error
	if isPostgres {
		err = sourceDB.Raw(sqlStr, tableName).Scan(&columns).Error
	} else {
		err = sourceDB.Raw(sqlStr, tableSchema, tableName).Scan(&columns).Error
	}

	if err != nil {
		panic(err)
	}
	md := &curdpb.MetadataInfo{
		Name:        RemoveLastChar(CamelName(tableName)),
		DisplayName: RemoveLastChar(CamelName(tableName)),
		System:      system,
		ParentID:    "da9436b6-25a7-4b82-b2da-827ae75be21f",
	}
	if len(columns) == 0 {
		return nil
	}
	for i, col := range columns {
		// fmt.Println(CamelName2(col.ColumnName))
		length, _ := strconv.Atoi(col.CharacterMaximumLength)
		md.MetadataFields = append(md.MetadataFields, &curdpb.MetadataField{
			Name:        CamelName2(col.ColumnName),
			Type:        col.DataType,
			Length:      int32(length),
			NotNull:     col.IsNullable == "NO",
			Comment:     col.ColumnComment,
			DisplayName: CamelName(col.ColumnName),
			Order:       int32(i) + 1,
		})
	}
	err = curdmodel.CreateMetadata(curdmodel.PBToMetadata(md))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return md
}

//TODO生成创建表SQL
//TODO生成更新表SQL
