package http

import (
	"context"
	"net/http"

	curdmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	Data map[string]interface{} `json:"data"`
}

// Add godoc
// @Summary 新增
// @Description 新增
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Param data body AddRequest true "Add Object"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/common/{pageName}/add [post]
func Add(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &AddRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:PageName为空", transID)
		return
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.Create(pageName, req.Data)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Update godoc
// @Summary 更新
// @Description 更新
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Param data body AddRequest true "Update Object"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/common/{pageName}/update [put]
func Update(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &AddRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:PageName为空", transID)
		return
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.Update(pageName, req.Data)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Delete godoc
// @Summary 删除
// @Description 删除
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete Page"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/common/{pageName}/delete [delete]
func Delete(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:PageName为空", transID)
		return
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.Delete(pageName, req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Enable godoc
// @Summary 禁用/启用
// @Description 禁用/启用
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Param data body apipb.EnableRequest true "Enable/Disable Page"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/common/{pageName}/enable [post]
func Enable(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.EnableRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:PageName为空", transID)
		return
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.Enable(pageName, req.Id, req.Enable)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Query godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} curdmodel.QueryResponse
// @Router /api/curd/common/{pageName}/query [get]
func Query(c *gin.Context) {
	req := &curdmodel.QueryRequest{}
	resp := &curdmodel.QueryResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "请求参数无效:PageName为空")
		return
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	req.PageName = pageName
	curdmodel.Query(req, resp)

	c.JSON(http.StatusOK, resp)
}

// GetAll godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Success 200 {object} curdmodel.QueryResponse
// @Router /api/curd/common/{pageName}/all [get]
func GetAll(c *gin.Context) {
	resp := &curdmodel.QueryResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := curdmodel.GetAll(pageName)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = data
	resp.Records = int64(len(data))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

// GetDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} model.CommonDetailResponse
// @Router /api/curd/common/{pageName}/detail [get]
func GetDetail(c *gin.Context) {
	resp := model.CommonDetailResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	resp.Data, err = curdmodel.GetDetailById(pageName, idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// GetDetailByName godoc
// @Summary 根据名称查询明细
// @Description 根据名称查询明细
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param name query string true "名称"
// @Param authorization header string true "jwt token"
// @Success 200 {object} model.CommonDetailResponse
// @Router /api/curd/common/{pageName}/detail/name [get]
func GetDetailByName(c *gin.Context) {
	resp := model.CommonDetailResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	name := c.Query("name")
	if name == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error
	resp.Data, err = curdmodel.GetDetailByName(pageName, name)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Copy godoc
// @Summary 复制
// @Description 复制
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param data body apipb.DelRequest true "Copy Page"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/common/{pageName}/detail [post]
func Copy(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:PageName为空", transID)
		return
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}

	err = curdmodel.Copy(pageName, req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// GetTree godoc
// @Summary 树形数据
// @Description 树形数据
// @Tags 通用增删改查接口
// @Accept  json
// @Produce  json
// @Param pageName path string true "页面配置名称"
// @Param authorization header string true "jwt token"
// @Success 200 {object} curdmodel.QueryResponse
// @Router /api/curd/common/{pageName}/tree [get]
func GetTree(c *gin.Context) {
	resp := &curdmodel.QueryResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	pageName := c.Param("pageName")
	if pageName == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error
	resp.Data, resp.Records, err = curdmodel.GetTree(pageName)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// :pageName是否为了控制接口的权限
func RegisterCurdRouter(r *gin.Engine) {
	g := r.Group("/api/curd/common")

	g.POST("/:pageName/add", Add)
	g.PUT("/:pageName/update", Update)
	g.GET("/:pageName/query", Query)
	g.GET("/:pageName/tree", GetTree)
	g.DELETE("/:pageName/delete", Delete)
	g.GET("/:pageName/all", GetAll)
	g.GET("/:pageName/detail", GetDetail)
	g.GET("/:pageName/detail/name", GetDetailByName)
	g.POST("/:pageName/copy", Copy)
	g.POST("/:pageName/enable", Enable)
}
