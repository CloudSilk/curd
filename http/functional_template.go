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

// AddFunctionalTemplate godoc
// @Summary 新增
// @Description 新增
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FunctionalTemplateInfo true "Add FunctionalTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/add [post]
func AddFunctionalTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FunctionalTemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建功能模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := model.CreateFunctionalTemplate(model.PBToFunctionalTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateFunctionalTemplate godoc
// @Summary 更新
// @Description 更新
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FunctionalTemplateInfo true "Update FunctionalTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/update [put]
func UpdateFunctionalTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FunctionalTemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新功能模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.UpdateFunctionalTemplate(model.PBToFunctionalTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteFunctionalTemplate godoc
// @Summary 删除
// @Description 删除
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete FunctionalTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/delete [delete]
func DeleteFunctionalTemplate(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除功能模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.DeleteFunctionalTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryFunctionalTemplate godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 功能模版管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param language query string false "语音"
// @Param group query string false "分组"
// @Param name query string false "名称"
// @Success 200 {object} apipb.QueryFunctionalTemplateResponse
// @Router /api/curd/functionaltemplate/query [get]
func QueryFunctionalTemplate(c *gin.Context) {
	req := &apipb.QueryFunctionalTemplateRequest{}
	resp := &apipb.QueryFunctionalTemplateResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	model.QueryFunctionalTemplate(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetFunctionalTemplateDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetFunctionalTemplateDetailResponse
// @Router /api/curd/functionaltemplate/detail [get]
func GetFunctionalTemplateDetail(c *gin.Context) {
	resp := &apipb.GetFunctionalTemplateDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := model.GetFunctionalTemplateByID(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.FunctionalTemplateToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllFunctionalTemplate godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllFunctionalTemplateResponse
// @Router /api/curd/functionaltemplate/all [get]
func GetAllFunctionalTemplate(c *gin.Context) {
	resp := &apipb.GetAllFunctionalTemplateResponse{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAllFunctionalTemplates()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.FunctionalTemplatesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// CopyFunctionalTemplate godoc
// @Summary 复制
// @Description 复制
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy FunctionalTemplate"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/copy [post]
func CopyFunctionalTemplate(c *gin.Context) {
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

	id, err := model.CopyFunctionalTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// EnableFunctionalTemplate godoc
// @Summary 禁用/启用
// @Description 禁用/启用
// @Tags 功能模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.EnableRequest true "Enable/Disable FunctionalTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/enable [post]
func EnableFunctionalTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.EnableRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,启用/禁用功能模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.EnableFunctionalTemplate(req.Id, req.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ImportFunctionalTemplate
// @Summary 导入
// @Description 导入
// @Tags 功能模版管理
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param files formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/import [post]
func ImportFunctionalTemplate(c *gin.Context) {
	resp := &apipb.QueryFunctionalTemplateResponse{
		Code: apipb.Code_Success,
	}
	//从功能模版中读取文件
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

	var list []*apipb.FunctionalTemplateInfo
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
		err = model.UpdateFunctionalTemplate(model.PBToFunctionalTemplate(f))
		if err == gorm.ErrRecordNotFound {
			_, err = model.CreateFunctionalTemplate(model.PBToFunctionalTemplate(f))
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

// ExportFunctionalTemplate godoc
// @Summary 导出
// @Description 导出
// @Tags 功能模版管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param language query string false "语音"
// @Param group query string false "分组"
// @Param name query string false "名称"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/functionaltemplate/export [get]
func ExportFunctionalTemplate(c *gin.Context) {
	req := &apipb.QueryFunctionalTemplateRequest{}
	resp := &apipb.QueryFunctionalTemplateResponse{
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
	model.QueryFunctionalTemplate(req, resp, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename=FunctionalTemplate.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

func GenFunctionalTemplateCode(c *gin.Context) {
	resp := &apipb.GenFunctionalTemplateCodeResponse{
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

	fileTemplate, err := model.GetFunctionalTemplateByID(tplIDStr)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	result, err := gen.GenCodeByFunctionalTemplate(fileTemplate, mdIDStr, nil)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = result

	c.JSON(http.StatusOK, resp)
}

func RegisterFunctionalTemplateRouter(r *gin.Engine) {
	g := r.Group("/api/curd/functionaltemplate")

	g.POST("add", AddFunctionalTemplate)
	g.PUT("update", UpdateFunctionalTemplate)
	g.GET("query", QueryFunctionalTemplate)
	g.DELETE("delete", DeleteFunctionalTemplate)
	g.GET("detail", GetFunctionalTemplateDetail)
	g.POST("copy", CopyFunctionalTemplate)
	g.POST("enable", EnableFunctionalTemplate)
	g.GET("all", GetAllFunctionalTemplate)
	g.GET("export", ExportFunctionalTemplate)
	g.POST("import", ImportFunctionalTemplate)
	g.GET("code", GenFunctionalTemplateCode)
}
