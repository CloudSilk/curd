package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CloudSilk/curd/gen"
	model "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddFileTemplate godoc
// @Summary 新增
// @Description 新增
// @Tags 文件模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FileTemplateInfo true "Add FileTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/filetemplate/add [post]
func AddFileTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FileTemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建文件模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := model.CreateFileTemplate(model.PBToFileTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateFileTemplate godoc
// @Summary 更新
// @Description 更新
// @Tags 文件模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FileTemplateInfo true "Update FileTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/filetemplate/update [put]
func UpdateFileTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FileTemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新文件模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.UpdateFileTemplate(model.PBToFileTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteFileTemplate godoc
// @Summary 删除
// @Description 删除
// @Tags 文件模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete FileTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/filetemplate/delete [delete]
func DeleteFileTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除文件模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.DeleteFileTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryFileTemplate godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 文件模版管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "文件名称"
// @Success 200 {object} apipb.QueryFileTemplateResponse
// @Router /api/curd/filetemplate/query [get]
func QueryFileTemplate(c *gin.Context) {
	req := &apipb.QueryFileTemplateRequest{}
	resp := &apipb.QueryFileTemplateResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	model.QueryFileTemplate(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetFileTemplateDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 文件模版管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetFileTemplateDetailResponse
// @Router /api/curd/filetemplate/detail [get]
func GetFileTemplateDetail(c *gin.Context) {
	resp := &apipb.GetFileTemplateDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := model.GetFileTemplateByID(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.FileTemplateToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllFileTemplate godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 文件模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllFileTemplateResponse
// @Router /api/curd/filetemplate/all [get]
func GetAllFileTemplate(c *gin.Context) {
	resp := &apipb.GetAllFileTemplateResponse{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAllFileTemplates()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.FileTemplatesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// CopyFileTemplate godoc
// @Summary 复制
// @Description 复制
// @Tags 文件模版管理
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy FileTemplate"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/filetemplate/copy [post]
func CopyFileTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}

	id, err := model.CopyFileTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// ImportFileTemplate
// @Summary 导入
// @Description 导入
// @Tags 文件模版管理
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param files formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/filetemplate/import [post]
func ImportFileTemplate(c *gin.Context) {
	resp := &apipb.QueryFileTemplateResponse{
		Code: apipb.Code_Success,
	}
	//从文件模版中读取文件
	file, fileHeader, err := c.Request.FormFile("files")
	if err != nil {
		fmt.Println(err)
		resp.Code = apipb.Code_BadRequest
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

	var list []*apipb.FileTemplateInfo
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
		err = model.UpdateFileTemplate(model.PBToFileTemplate(f))
		if err == gorm.ErrRecordNotFound {
			_, err = model.CreateFileTemplate(model.PBToFileTemplate(f))
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

// ExportFileTemplate godoc
// @Summary 导出
// @Description 导出
// @Tags 文件模版管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "文件名称"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/filetemplate/export [get]
func ExportFileTemplate(c *gin.Context) {
	req := &apipb.QueryFileTemplateRequest{}
	resp := &apipb.QueryFileTemplateResponse{
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
	model.QueryFileTemplate(req, resp, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename=FileTemplate.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

func GenFileTemplateCode(c *gin.Context) {
	resp := &apipb.GenFileTemplateCodeResponse{
		Code: apipb.Code_Success,
	}
	tplIDStr := c.Query("tplID")
	if tplIDStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	mdIDStr := c.Query("mdID")
	if mdIDStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	fileTemplate, err := model.GetFileTemplateByID(tplIDStr)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	templateContent, code, err := gen.GenCodeByFileTemplate(fileTemplate, mdIDStr)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = &apipb.GenFileTemplateCodeInfo{
		Template: templateContent,
		Code:     code,
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterFileTemplateRouter(r *gin.Engine) {
	g := r.Group("/api/curd/filetemplate")

	g.POST("add", AddFileTemplate)
	g.PUT("update", UpdateFileTemplate)
	g.GET("query", QueryFileTemplate)
	g.DELETE("delete", DeleteFileTemplate)
	g.GET("detail", GetFileTemplateDetail)
	g.POST("copy", CopyFileTemplate)
	g.GET("all", GetAllFileTemplate)
	g.GET("export", ExportFileTemplate)
	g.POST("import", ImportFileTemplate)
	g.GET("code", GenFileTemplateCode)
}
