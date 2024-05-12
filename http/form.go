package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	form "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddForm godoc
// @Summary 新增表单
// @Description 新增表单
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FormInfo true "Add Form"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/add [post]
func AddForm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FormInfo{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建表单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//只有平台租户才能创建其他租户的表单
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = form.CreateForm(form.PBToForm(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateForm godoc
// @Summary 更新表单管理
// @Description 更新表单管理
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FormInfo true "Update Form"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/update [put]
func UpdateForm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FormInfo{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建客户请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能更新其他租户的表单
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = form.UpdateForm(form.PBToForm(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteForm godoc
// @Summary 删除表单管理
// @Description 删除表单管理
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.DelRequest true "Delete Form"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/delete [delete]
func DeleteForm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除表单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = form.DeleteForm(req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryForm godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 表单管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param formIndex query int false "从1开始"
// @Param formSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param pageName query string false "表单名称"
// @Param group query string false "分组"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.QueryFormResponse
// @Router /api/form/query [get]
func QueryForm(c *gin.Context) {
	req := &apipb.QueryFormRequest{}
	resp := &apipb.QueryFormResponse{
		Code: model.Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的表单
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}

	form.QueryForm(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllForm godoc
// @Summary 查询所有表单管理
// @Description 查询所有表单管理
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllFormResponse
// @Router /api/form/all [get]
func GetAllForm(c *gin.Context) {
	resp := &apipb.QueryFormResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.GetAllFormRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if os.Getenv("FORM_DISABLE_AUTH") != "true" {
		//只有平台租户才能查询其他租户的表单
		tenantID := ucm.GetTenantID(c)
		if tenantID != constants.PlatformTenantID {
			req.TenantID = tenantID
		}
	}

	forms, err := form.GetAllForms(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = form.FormsToPB(forms)
	resp.Records = int64(len(forms))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

// GetFormDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetFormDetailResponse
// @Router /api/form/detail [get]
func GetFormDetail(c *gin.Context) {
	resp := &apipb.GetFormDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := form.GetFormById(idStr, c.Query("containerVersions") == "true")
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = form.FormToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateFormSchema godoc
// @Summary 更新表单Schema
// @Description 更新表单Schema
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FormInfo true "Update Form"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/schema [put]
func UpdateFormSchema(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FormInfo{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新表单Schema请求参数无效:%v", transID, err)
		return
	}
	err = form.UpdateFormSchema(form.PBToForm(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// PublishForm godoc
// @Summary 发布表单版本
// @Description 发布表单版本
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.FormVersion true "Update Form"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/publish [post]
func PublishForm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.FormVersion{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s发布表单版本请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = form.PublishForm(form.PBToFormVersion(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// SwitchFormVersion godoc
// @Summary 切换表单版本
// @Description 切换表单版本
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param versionID query string false "Version ID"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/switch [post]
func SwitchFormVersion(c *gin.Context) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("versionID")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	err := form.SwitchFormVersion(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// GetFormVersionDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetVersionResponse
// @Router /api/form/version/detail [get]
func GetFormVersionDetail(c *gin.Context) {
	resp := &apipb.GetVersionResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := form.GetFormVersionById(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = form.FormVersionToPB(&data)
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterFormRouter(r *gin.Engine) {
	formGroup := r.Group("/api/form")

	formGroup.POST("add", AddForm)
	formGroup.PUT("update", UpdateForm)
	formGroup.GET("query", QueryForm)
	formGroup.DELETE("delete", DeleteForm)
	formGroup.GET("all", GetAllForm)
	formGroup.GET("detail", GetFormDetail)
	formGroup.PUT("schema", UpdateFormSchema)
	formGroup.POST("publish", PublishForm)
	formGroup.POST("switch", SwitchFormVersion)
	formGroup.GET("version/detail", GetFormVersionDetail)
	formGroup.POST("copy", CopyForm)
	formGroup.GET("export", ExportForm)
	formGroup.POST("import", ImportForm)
}

// CopyForm godoc
// @Summary 复制表单
// @Description 复制表单
// @Tags 表单管理
// @Accept  json
// @Produce  json
// @Param account body apipb.DelRequest true "Copy Form"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/copy [post]
func CopyForm(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}

	err = form.CopyForm(req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ExportForm godoc
// @Summary 导出
// @Description 导出
// @Tags 表单管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param formIndex query int false "从1开始"
// @Param formSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param pageName query string false "表单名称"
// @Param group query string false "分组"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/export [get]
func ExportForm(c *gin.Context) {
	req := &apipb.QueryFormRequest{}
	resp := &apipb.QueryFormResponse{
		Code: model.Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	req.PageIndex = 1
	req.PageSize = 1000
	form.QueryForm(req, resp, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename=Form.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

// ImportForm
// @Summary 导入
// @Description 导入
// @Tags 表单管理
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param file formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/form/import [post]
func ImportForm(c *gin.Context) {
	resp := &apipb.QueryFormResponse{
		Code: model.Success,
	}
	//从表单中读取文件
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
	buf, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var forms []*apipb.FormInfo
	err = json.Unmarshal(buf, &forms)
	if err != nil {
		fmt.Println(err)
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	successCount := 0
	failCount := 0
	var names []string
	for _, f := range forms {
		err = form.UpdateFormAll(form.PBToForm(f))
		if err == gorm.ErrRecordNotFound {
			err = form.CreateForm(form.PBToForm(f))
		}
		if err != nil {
			failCount++
			names = append(names, f.PageName+"-"+f.Name)
			fmt.Println(err)
		} else {
			successCount++
		}
	}
	resp.Message = fmt.Sprintf("导入成功数量:%d,导入失败数量:%d,导入失败的表单名称：%s", successCount, failCount, strings.Join(names, ","))
	c.JSON(http.StatusOK, resp)
}
