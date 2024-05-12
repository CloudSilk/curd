package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	curdmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddPage godoc
// @Summary 新增页面
// @Description 新增页面
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.PageInfo true "Add Page"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/add [post]
func AddPage(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.PageInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建页面配置请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能创建其他租户的页面配置
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = curdmodel.CreatePage(curdmodel.PBToPage(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePage godoc
// @Summary 更新页面配置
// @Description 更新页面配置
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.PageInfo true "Update Page"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/update [put]
func UpdatePage(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.PageInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建页面配置请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能更新其他租户的页面配置
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = curdmodel.UpdatePage(curdmodel.PBToPage(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeletePage godoc
// @Summary 删除页面配置
// @Description 删除页面配置
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete Page"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/delete [delete]
func DeletePage(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建Page请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.DeletePage(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// EnablePage godoc
// @Summary 禁用/启用页面配置
// @Description 禁用/启用页面配置
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.EnableRequest true "Enable/Disable Page"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/enable [post]
func EnablePage(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建页面配置请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.EnablePage(req.Id, req.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryPage godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param enable query int false "是否启用"
// @Param type query int false "类型"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.QueryPageResponse
// @Router /api/curd/page/query [get]
func QueryPage(c *gin.Context) {
	req := &apipb.QueryPageRequest{}
	resp := &apipb.QueryPageResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的页面配置
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	curdmodel.QueryPage(req, resp, false, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllPage godoc
// @Summary 查询所有页面配置
// @Description 查询所有页面配置
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllPageResponse
// @Router /api/curd/page/all [get]
func GetAllPage(c *gin.Context) {
	resp := &apipb.QueryPageResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.QueryPageRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的页面配置
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	data, err := curdmodel.GetAllPage(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = curdmodel.PagesToPB(data)
	resp.Records = int64(len(data))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

// GetPageDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetPageDetailResponse
// @Router /api/curd/page/detail [get]
func GetPageDetail(c *gin.Context) {
	resp := &apipb.GetPageDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	page, err := curdmodel.GetPageByID(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = curdmodel.PageToPB(page)
	}
	c.JSON(http.StatusOK, resp)
}

// GetPageDetailByName godoc
// @Summary 根据名称查询明细
// @Description 根据名称查询明细
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param name query string true "名称"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetPageDetailResponse
// @Router /api/curd/page/detail/name [get]
func GetPageDetailByName(c *gin.Context) {
	resp := &apipb.GetPageDetailResponse{
		Code: apipb.Code_Success,
	}
	name := c.Query("name")
	if name == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error
	page, err := curdmodel.GetPageByName(name)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = curdmodel.PageToPB(page)
	}
	c.JSON(http.StatusOK, resp)
}

// CopyPage godoc
// @Summary 复制页面配置
// @Description 复制页面配置
// @Tags 页面配置
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy Page"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/detail [post]
func CopyPage(c *gin.Context) {
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

	err = curdmodel.CopyPage(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ExportPage godoc
// @Summary 导出
// @Description 导出
// @Tags 页面配置
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param enable query int false "是否启用"
// @Param type query int false "类型"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.QueryPageResponse
// @Router /api/curd/page/export [get]
func ExportPage(c *gin.Context) {
	req := &apipb.QueryPageRequest{}
	resp := &apipb.QueryPageResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的页面配置
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	req.PageIndex = 1
	req.PageSize = 1000
	curdmodel.QueryPage(req, resp, true, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename=Page.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

// ImportPage
// @Summary 导入
// @Description 导入
// @Tags 页面配置
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param files formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/page/import [post]
func ImportPage(c *gin.Context) {
	resp := &apipb.QueryPageResponse{
		Code: apipb.Code_Success,
	}
	//从表单中读取文件
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
	buf, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var pages []*apipb.PageInfo
	err = json.Unmarshal(buf, &pages)
	if err != nil {
		fmt.Println(err)
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	successCount := 0
	failCount := 0
	for _, page := range pages {
		err = curdmodel.UpdatePage(curdmodel.PBToPage(page))
		if err == gorm.ErrRecordNotFound {
			err = curdmodel.CreatePage(curdmodel.PBToPage(page))
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

func RegisterPageRouter(r *gin.Engine) {
	g := r.Group("/api/curd/page")

	g.POST("add", AddPage)
	g.PUT("update", UpdatePage)
	g.GET("query", QueryPage)
	g.DELETE("delete", DeletePage)
	g.GET("all", GetAllPage)
	g.GET("detail", GetPageDetail)
	g.GET("detail/name", GetPageDetailByName)
	g.POST("copy", CopyPage)
	g.POST("enable", EnablePage)
	g.GET("export", ExportPage)
	g.POST("import", ImportPage)
}
