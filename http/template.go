package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/curd/gen"
	ctmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

func AddTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.TemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建模板请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
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
	err = ctmodel.CreateTemplate(ctmodel.PBToTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		gen.AddTemplate(ctmodel.PBToTemplate(req))
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.TemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建模板请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ctmodel.UpdateTemplate(ctmodel.PBToTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		gen.AddTemplate(ctmodel.PBToTemplate(req))
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteTemplate(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建模板请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ctmodel.DeleteTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		gen.DeleteTemplate(req.Id)
	}
	c.JSON(http.StatusOK, resp)
}

func QueryTemplate(c *gin.Context) {
	req := &apipb.QueryTemplateRequest{}
	resp := &apipb.QueryTemplateResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ctmodel.QueryTemplate(req, resp)

	c.JSON(http.StatusOK, resp)
}

func GetAllTemplate(c *gin.Context) {
	resp := &apipb.QueryTemplateResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.QueryTemplateRequest{}
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
	tpls, err := ctmodel.GetAllTemplates(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = ctmodel.TemplatesToPB(tpls)
	resp.Records = int64(len(tpls))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

func GetTemplateDetail(c *gin.Context) {
	resp := &apipb.GetTemplateDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := ctmodel.GetTemplateById(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = ctmodel.TemplateToPB(&data)
	}
	c.JSON(http.StatusOK, resp)
}

func GenTemplateCode(c *gin.Context) {
	resp := &apipb.GenTemplateCodeResponse{
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

	str, err := gen.GenCode(tplIDStr, mdIDStr)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = str
	c.JSON(http.StatusOK, resp)
}

// CopyTemplate godoc
// @Summary 复制
// @Description 复制
// @Tags 代码模板管理
// @Accept  json
// @Produce  json
// @Param data body apipb.DelRequest true "Copy Template"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/template/copy [post]
func CopyTemplate(c *gin.Context) {
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

	err = ctmodel.CopyTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterTemplateRoute(r *gin.Engine) {
	tplGroup := r.Group("/api/curd/template")

	tplGroup.POST("add", AddTemplate)
	tplGroup.PUT("update", UpdateTemplate)
	tplGroup.GET("query", QueryTemplate)
	tplGroup.DELETE("delete", DeleteTemplate)
	tplGroup.GET("all", GetAllTemplate)
	tplGroup.GET("detail", GetTemplateDetail)
	tplGroup.GET("code", GenTemplateCode)
	tplGroup.POST("copy", CopyTemplate)
}
