package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	apipb "github.com/CloudSilk/form/api"
    model "github.com/CloudSilk/form/model"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "codeup.aliyun.com/atali/proto/utils/middleware"
	"gorm.io/gorm"
)

// Add{{$.Metadata.Name}} godoc
// @Summary 新增
// @Description 新增
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.{{$.Metadata.Name}}Info true "Add {{$.Metadata.Name}}"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/add [post]
func Add{{$.Metadata.Name}}(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.{{$.Metadata.Name}}Info{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建{{$.Metadata.DisplayName}}请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	req.TenantID = ucm.GetTenantID(c)
	err = model.Create{{$.Metadata.Name}}(model.PBTo{{$.Metadata.Name}}(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Update{{$.Metadata.Name}} godoc
// @Summary 更新
// @Description 更新
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.{{$.Metadata.Name}}Info true "Update {{$.Metadata.Name}}"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/update [put]
func Update{{$.Metadata.Name}}(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.{{$.Metadata.Name}}Info{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新{{$.Metadata.DisplayName}}请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.Update{{$.Metadata.Name}}(model.PBTo{{$.Metadata.Name}}(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Delete{{$.Metadata.Name}} godoc
// @Summary 删除
// @Description 删除
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete {{$.Metadata.Name}}"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/delete [delete]
func Delete{{$.Metadata.Name}}(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除{{$.Metadata.DisplayName}}请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.Delete{{$.Metadata.Name}}(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Query{{$.Metadata.Name}} godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param pageName query string false "{{$.Metadata.DisplayName}}名称"
// @Param group query string false "分组"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.Query{{$.Metadata.Name}}Response
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/query [get]
func Query{{$.Metadata.Name}}(c *gin.Context) {
	req := &apipb.Query{{$.Metadata.Name}}Request{}
	resp := &apipb.Query{{$.Metadata.Name}}Response{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	model.Query{{$.Metadata.Name}}(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAll{{$.Metadata.Name}} godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAll{{$.Metadata.Name}}Response
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/all [get]
func GetAll{{$.Metadata.Name}}(c *gin.Context) {
	resp := &apipb.Query{{$.Metadata.Name}}Response{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAll{{$.Metadata.Name}}s()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.{{$.Metadata.Name}}sToPB(list)
	resp.Records = int64(len(list))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

// Get{{$.Metadata.Name}}Detail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.Get{{$.Metadata.Name}}DetailResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/detail [get]
func Get{{$.Metadata.Name}}Detail(c *gin.Context) {
	resp := &apipb.Get{{$.Metadata.Name}}DetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := model.Get{{$.Metadata.Name}}ById(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.{{$.Metadata.Name}}ToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// Copy{{$.Metadata.Name}} godoc
// @Summary 复制
// @Description 复制
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy {{$.Metadata.Name}}"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/copy [post]
func Copy{{$.Metadata.Name}}(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code =apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}

	err = model.Copy{{$.Metadata.Name}}(req.Id)
	if err != nil {
		resp.Code =apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Export{{$.Metadata.Name}} godoc
// @Summary 导出
// @Description 导出
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param pageName query string false "{{$.Metadata.DisplayName}}名称"
// @Param group query string false "分组"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/export [get]
func Export{{$.Metadata.Name}}(c *gin.Context) {
	req := &apipb.Query{{$.Metadata.Name}}Request{}
	resp := &apipb.Query{{$.Metadata.Name}}Response{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	req.PageIndex = 1
	req.PageSize = 1000
	model.Query{{$.Metadata.Name}}(req, resp, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename={{$.Metadata.Name}}.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

// Import{{$.Metadata.Name}}
// @Summary 导入
// @Description 导入
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param files formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/import [post]
func Import{{$.Metadata.Name}}(c *gin.Context) {
	resp := &apipb.Query{{$.Metadata.Name}}Response{
		Code: apipb.Code_Success,
	}
	//从{{$.Metadata.DisplayName}}中读取文件
	file, fileHeader, err := c.Request.FormFile("files")
	if err != nil {
		fmt.Println(err)
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	//defer 结束时关闭文件
	defer file.Close()
	fmt.Println("filename: " + fileHeader.Filename)
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var list []*apipb.{{$.Metadata.Name}}Info
	err = json.Unmarshal(buf, &list)
	if err != nil {
		fmt.Println(err)
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	successCount := 0
	failCount := 0
	for _, f := range list {
		err = model.Update{{$.Metadata.Name}}All(model.PBTo{{$.Metadata.Name}}(f))
		if err == gorm.ErrRecordNotFound {
			err = model.Create{{$.Metadata.Name}}(model.PBTo{{$.Metadata.Name}}(f))
		}
		if err != nil {
			failCount++
			fmt.Println(err)
		} else {
			successCount++
		}
	}
	resp.Message = fmt.Sprintf("导入成功数量:%d,导入失败数量:%d", successCount, failCount)
	c.JSON(http.StatusOK, resp)
}

// Enable{{$.Metadata.Name}} godoc
// @Summary 禁用/启用
// @Description 禁用/启用
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.EnableRequest true "Enable/Disable {{$.Metadata.Name}}"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/enable [post]
func Enable{{$.Metadata.Name}}(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.EnableRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,启用/禁用{{$.Metadata.DisplayName}}请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code =apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.Enable{{$.Metadata.Name}}(req.Id, req.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Get{{$.Metadata.Name}}Tree godoc
// @Summary 查询所有-Tree
// @Description 查询所有-Tree
// @Tags {{$.Metadata.DisplayName}}管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAll{{$.Metadata.Name}}Response
// @Router /api/{{LowerSnakeCase $.Metadata.Name}}/true [get]
func Get{{$.Metadata.Name}}Tree(c *gin.Context) {
	resp := &apipb.Query{{$.Metadata.Name}}Response{
		Code: apipb.Code_Success,
	}
	var err error
	list, err := {{$.Metadata.Package}}.Get{{$.Metadata.Name}}Tree()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.{{$.Metadata.Name}}sToPB(list)
	c.JSON(http.StatusOK, resp)
}


func Register{{$.Metadata.Name}}Router(r *gin.Engine) {
	g := r.Group("/api/{{ToLower $.Metadata.Package}}/{{ToLower $.Metadata.Name}}")

	g.POST("add", Add{{$.Metadata.Name}})
	g.PUT("update", Update{{$.Metadata.Name}})
	g.GET("query", Query{{$.Metadata.Name}})
	g.DELETE("delete", Delete{{$.Metadata.Name}})
	g.GET("all", GetAll{{$.Metadata.Name}})
    g.GET("tree", Get{{$.Metadata.Name}}Tree)
	g.GET("detail", Get{{$.Metadata.Name}}Detail)
	g.POST("copy", Copy{{$.Metadata.Name}})
        g.POST("enable", Enable{{$.Metadata.Name}})
	g.GET("export", Export{{$.Metadata.Name}})
	g.POST("import", Import{{$.Metadata.Name}})
}