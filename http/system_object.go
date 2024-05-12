package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddSystemObject godoc
// @Summary 新增
// @Description 新增
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.SystemObjectInfo true "Add SystemObject"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/systemobject/add [post]
func AddSystemObject(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.SystemObjectInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建系统定义的对象请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := model.CreateSystemObject(model.PBToSystemObject(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateSystemObject godoc
// @Summary 更新
// @Description 更新
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.SystemObjectInfo true "Update SystemObject"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/systemobject/update [put]
func UpdateSystemObject(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.SystemObjectInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新系统定义的对象请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.UpdateSystemObject(model.PBToSystemObject(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QuerySystemObject godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Param language query string false "编程语言"
// @Param type query int false "类型"
// @Param enable query string false "是否启用"
// @Success 200 {object} apipb.QuerySystemObjectResponse
// @Router /api/curd/systemobject/query [get]
func QuerySystemObject(c *gin.Context) {
	req := &apipb.QuerySystemObjectRequest{}
	resp := &apipb.QuerySystemObjectResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	model.QuerySystemObject(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllSystemObject godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllSystemObjectResponse
// @Router /api/curd/systemobject/all [get]
func GetAllSystemObject(c *gin.Context) {
	resp := &apipb.GetAllSystemObjectResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.GetAllSystemObjectRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	list, err := model.GetAllSystemObjects(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.SystemObjectsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetSystemObjectDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetSystemObjectDetailResponse
// @Router /api/curd/systemobject/detail [get]
func GetSystemObjectDetail(c *gin.Context) {
	resp := &apipb.GetSystemObjectDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := model.GetSystemObjectByID(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemObjectToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteSystemObject godoc
// @Summary 删除
// @Description 删除
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete SystemObject"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/systemobject/delete [delete]
func DeleteSystemObject(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除系统定义的对象请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.DeleteSystemObject(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// EnableSystemObject godoc
// @Summary 禁用/启用
// @Description 禁用/启用
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.EnableRequest true "Enable/Disable SystemObject"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/systemobject/enable [post]
func EnableSystemObject(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,启用/禁用系统定义的对象请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.EnableSystemObject(req.Id, req.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// CopySystemObject godoc
// @Summary 复制
// @Description 复制
// @Tags 系统定义的对象管理
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy SystemObject"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/systemobject/copy [post]
func CopySystemObject(c *gin.Context) {
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

	id, err := model.CopySystemObject(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterSystemObjectRouter(r *gin.Engine) {
	g := r.Group("/api/curd/systemobject")

	g.POST("add", AddSystemObject)
	g.PUT("update", UpdateSystemObject)
	g.GET("query", QuerySystemObject)
	g.DELETE("delete", DeleteSystemObject)
	g.GET("all", GetAllSystemObject)
	g.GET("detail", GetSystemObjectDetail)
	g.POST("copy", CopySystemObject)
	g.POST("enable", EnableSystemObject)
}
