package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddCell godoc
// @Summary 新增
// @Description 新增
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.CellInfo true "Add Cell"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/add [post]
func AddCell(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CellInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Cell请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 只有平台租户才能为其他租户创建Cell
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = model.CreateCell(model.PBToCell(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCell godoc
// @Summary 更新
// @Description 更新
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.CellInfo true "Update Cell"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/update [put]
func UpdateCell(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CellInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新Cell请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 只有平台租户才能为其他租户更新Cell
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = model.UpdateCell(model.PBToCell(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCell godoc
// @Summary 删除
// @Description 删除
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete Cell"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/delete [delete]
func DeleteCell(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除Cell请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.DeleteCell(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryCell godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags Cell管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param mustSource query string false "Must Source"
// @Param name query string false "名称"
// @Param system query string false "系统"
// @Param view query string false "视图"
// @Param shape query string false "形状"
// @Param isEdge query string false "边"
// @Param common query string false "是否常用"
// @Param resizing query string false "调整大小"
// @Param group query string false "分组"
// @Param mustTarget query string false "Must Target"
// @Success 200 {object} apipb.QueryCellResponse
// @Router /api/curd/cell/query [get]
func QueryCell(c *gin.Context) {
	req := &apipb.QueryCellRequest{}
	resp := &apipb.QueryCellResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//只有平台租户才能查询其他租户的Cell
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	model.QueryCell(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetCellDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetCellDetailResponse
// @Router /api/curd/cell/detail [get]
func GetCellDetail(c *gin.Context) {
	resp := &apipb.GetCellDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := model.GetCellByID(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CellToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllCell godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllCellResponse
// @Router /api/curd/cell/all [get]
func GetAllCell(c *gin.Context) {
	req := &apipb.GetAllCellRequest{}
	resp := &apipb.GetAllCellResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的项目
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}

	list, err := model.GetAllCells(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.CellsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// CopyCell godoc
// @Summary 复制
// @Description 复制
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy Cell"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/copy [post]
func CopyCell(c *gin.Context) {
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

	err = model.CopyCell(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// EnableCell godoc
// @Summary 禁用/启用
// @Description 禁用/启用
// @Tags Cell管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.EnableRequest true "Enable/Disable Cell"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/enable [post]
func EnableCell(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,启用/禁用Cell请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.EnableCell(req.Id, req.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ImportCell
// @Summary 导入
// @Description 导入
// @Tags Cell管理
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param files formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/import [post]
func ImportCell(c *gin.Context) {
	resp := &apipb.QueryCellResponse{
		Code: apipb.Code_Success,
	}
	//从Cell中读取文件
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

	var list []*apipb.CellInfo
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
		err = model.UpdateCell(model.PBToCell(f))
		if err == gorm.ErrRecordNotFound {
			err = model.CreateCell(model.PBToCell(f))
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

// ExportCell godoc
// @Summary 导出
// @Description 导出
// @Tags Cell管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param mustSource query string false "Must Source"
// @Param name query string false "名称"
// @Param system query string false "系统"
// @Param view query string false "视图"
// @Param shape query string false "形状"
// @Param isEdge query string false "边"
// @Param common query string false "是否常用"
// @Param resizing query string false "调整大小"
// @Param group query string false "分组"
// @Param mustTarget query string false "Must Target"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/cell/export [get]
func ExportCell(c *gin.Context) {
	req := &apipb.QueryCellRequest{}
	resp := &apipb.QueryCellResponse{
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
	model.QueryCell(req, resp, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename=Cell.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

func RegisterCellRouter(r *gin.Engine) {
	g := r.Group("/api/curd/cell")

	g.POST("add", AddCell)
	g.PUT("update", UpdateCell)
	g.GET("query", QueryCell)
	g.DELETE("delete", DeleteCell)
	g.GET("detail", GetCellDetail)
	g.POST("copy", CopyCell)
	g.POST("enable", EnableCell)
	g.GET("all", GetAllCell)
	g.GET("export", ExportCell)
	g.POST("import", ImportCell)
}
