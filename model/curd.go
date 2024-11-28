package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/CloudSilk/pkg/model"
	"gorm.io/gorm/schema"
)

var NamingStrategy schema.NamingStrategy

func Create(pageName string, m map[string]interface{}) error {

	page, err := GetPageByName(pageName)
	if err != nil {
		return err
	}
	md := page.Metadata
	data := make(map[string]interface{})
	for _, field := range md.MetadataFields {
		data[LowerSnakeCase(field.Name)] = m[field.Name]
	}

	var uniqueFields []string
	var fieldValues []interface{}

	for _, field := range md.MetadataFields {
		if field.Unique {
			uniqueFields = append(uniqueFields, " "+LowerSnakeCase(field.Name)+" =? ")
			fieldValues = append(fieldValues, m[field.Name])
		}
	}

	if len(uniqueFields) > 1 {
		duplication, err := dbClient.CheckDuplication(dbClient.DB().Table(NamingStrategy.TableName(page.Metadata.Name)), strings.Join(uniqueFields, " and "), fieldValues...)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同" + page.Title)
		}
	}

	var updateFields []string
	var updateValues []interface{}
	var list []string
	for _, field := range md.MetadataFields {
		if field.Name == "id" || field.Name == "ID" {
			continue
		}
		updateFields = append(updateFields, LowerSnakeCase(field.Name))
		list = append(list, "?")
		updateValues = append(updateValues, m[field.Name])
	}

	insertSql := fmt.Sprintf("insert into `%s`(%s) values(%s)", NamingStrategy.TableName(page.Metadata.Name), strings.Join(updateFields, ","), strings.Join(list, ","))
	return dbClient.DB().Exec(insertSql, updateValues...).Error
}

func Delete(pageName string, id string) (err error) {
	page, err := GetPageByName(pageName)
	if err != nil {
		return err
	}
	return dbClient.DB().Exec(fmt.Sprintf("delete from %s where id=?", NamingStrategy.TableName(page.Metadata.Name)), id).Error
}

type QueryResponse struct {
	model.CommonResponse
	Data []map[string]interface{} `json:"data"`
}

type QueryRequest struct {
	model.CommonRequest
	PageName string                 `json:"pageName" form:"pageName" uri:"pageName"`
	Data     map[string]interface{} `json:"data" form:"data" uri:"data"`
}

func Query(req *QueryRequest, resp *QueryResponse) {
	page, err := GetPageByName(req.PageName)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}

	db := dbClient.DB().Table(NamingStrategy.TableName(page.Metadata.Name))

	for key, value := range req.Data {
		db = db.Where(fmt.Sprintf("%s = ?", key), value)
	}

	OrderStr := "`id`"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}

	resp.Total, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.Current, OrderStr, &resp.Data, nil)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	result := make([]map[string]interface{}, len(resp.Data))
	for i, data := range resp.Data {
		d := make(map[string]interface{})
		for key, value := range data {
			d[CamelName2(key)] = value
		}
		result[i] = d
	}
	resp.Data = result
}

func GetAll(pageName string) (list []map[string]interface{}, err error) {
	page, err := GetPageByName(pageName)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	err = dbClient.DB().Table(NamingStrategy.TableName(page.Metadata.Name)).Find(&result).Error
	if err != nil {
		return nil, err
	}
	for _, data := range result {
		d := make(map[string]interface{})
		for key, value := range data {
			d[CamelName2(key)] = value
		}
		list = append(list, d)
	}
	return
}

func GetDetailById(pageName string, id string) (data map[string]interface{}, err error) {
	page, err := GetPageByName(pageName)
	if err != nil {
		return nil, err
	}
	data = make(map[string]interface{})
	result := make(map[string]interface{})
	err = dbClient.DB().Raw(fmt.Sprintf("select * from %s where id=?", NamingStrategy.TableName(page.Metadata.Name)), id).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	for key, value := range result {
		data[CamelName2(key)] = value
	}
	return
}

func Update(pageName string, m map[string]interface{}) error {
	page, err := GetPageByName(pageName)
	if err != nil {
		return err
	}
	var uniqueFields []string
	var fieldValues []interface{}
	md := page.Metadata
	uniqueFields = append(uniqueFields, "id <> ?")
	fieldValues = append(fieldValues, m["ID"])
	for _, field := range md.MetadataFields {
		if field.Unique {
			uniqueFields = append(uniqueFields, " "+LowerSnakeCase(field.Name)+" =? ")
			fieldValues = append(fieldValues, m[field.Name])
		}
	}

	var updateFields []string
	var updateValues []interface{}
	for _, field := range md.MetadataFields {
		updateFields = append(updateFields, LowerSnakeCase(field.Name)+"=?")
		updateValues = append(updateValues, m[field.Name])
	}
	id := m["id"]
	if id == nil {
		id = m["ID"]
	}
	updateValues = append(updateValues, id)
	if len(uniqueFields) > 1 {
		duplication, err := dbClient.CheckDuplication(dbClient.DB().Table(NamingStrategy.TableName(page.Metadata.Name)), strings.Join(uniqueFields, " and "), fieldValues...)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同" + page.Title)
		}
	}
	return dbClient.DB().Exec(fmt.Sprintf("UPDATE %s SET %s where id=?", NamingStrategy.TableName(page.Metadata.Name), strings.Join(updateFields, ",")), updateValues...).Error
}

func GetDetailByName(pageName, name string) (map[string]interface{}, error) {
	page, err := GetPageByName(pageName)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	result := make(map[string]interface{})
	err = dbClient.DB().Raw(fmt.Sprintf("select * from %s where name=?", NamingStrategy.TableName(page.Metadata.Name)), name).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	for key, value := range result {
		data[CamelName2(key)] = value
	}
	return data, nil
}

func Copy(pageName string, id string) error {
	from, err := GetDetailById(pageName, id)
	if err != nil {
		return err
	}

	from["ID"] = 0
	if _, ok := from["name"]; ok {
		from["name"] = fmt.Sprintf("%s Copy", from["name"])
	}
	fmt.Println(from)
	return Create(pageName, from)
}

func Enable(pageName string, id string, enable bool) error {
	page, err := GetPageByName(pageName)
	if err != nil {
		return err
	}
	return dbClient.DB().Table(NamingStrategy.TableName(page.Metadata.Name)).Where("id=?", id).Update("enable", enable).Error
}

func GetTree(pageName string) (list []map[string]interface{}, total int64, err error) {
	page, err := GetPageByName(pageName)
	if err != nil {
		return nil, 0, err
	}

	var menuList []map[string]interface{}
	treeMap, err := GetTreeMap(page)
	menuList = treeMap[""]
	for i := 0; i < len(menuList); i++ {
		err = GetChildrenList(menuList[i], treeMap)
	}
	return menuList, total, err
}

func GetTreeMap(page *Page) (treeMap map[string][]map[string]interface{}, err error) {
	var all []map[string]interface{}
	treeMap = make(map[string][]map[string]interface{})
	err = dbClient.DB().Table(NamingStrategy.TableName(page.Metadata.Name)).Order("level").Find(&all).Error
	for _, v := range all {
		d := make(map[string]interface{})
		for key, value := range v {
			d[CamelName2(key)] = value
		}
		parentID, ok := d["parentID"]
		if !ok {
			return nil, fmt.Errorf("不是树形接口的数据,不存在parentID")
		}
		d["key"] = d["ID"]
		d["title"] = d["name"]
		treeMap[parentID.(string)] = append(treeMap[parentID.(string)], d)
	}
	return treeMap, err
}

func GetChildrenList(location map[string]interface{}, treeMap map[string][]map[string]interface{}) (err error) {
	id, ok := location["ID"]
	if !ok {
		return fmt.Errorf("不是树形接口的数据,不存在ID")
	}
	children, ok := treeMap[id.(string)]
	if !ok {
		return nil
	}
	location["children"] = children
	for i := 0; i < len(children); i++ {
		err = GetChildrenList(children[i], treeMap)
	}
	return err
}
