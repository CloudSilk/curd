package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CloudSilk/curd/gen"
	curdmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddMetadata(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.MetadataInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建元数据请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 只有平台租户才能为其他租户创建元数据
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = curdmodel.CreateMetadata(curdmodel.PBToMetadata(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		gen.AddMetadata(curdmodel.PBToMetadata(req))
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateMetadata(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.MetadataInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建元数据请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能更新其他租户元数据
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = curdmodel.UpdateMetadata(curdmodel.PBToMetadata(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		gen.AddMetadata(curdmodel.PBToMetadata(req))
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteMetadata(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建元数据请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = curdmodel.DeleteMetadata(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		gen.DeleteMetadata(req.Id)
	}
	c.JSON(http.StatusOK, resp)
}

func QueryMetadata(c *gin.Context) {
	req := &apipb.QueryMetadataRequest{}
	resp := &apipb.QueryMetadataResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的元数据
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	curdmodel.QueryMetadata(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

func GetAllMetadata(c *gin.Context) {
	resp := &apipb.QueryMetadataResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.QueryMetadataRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的元数据
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	metadatas, err := curdmodel.GetAllMetadatas(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = curdmodel.MetadatasToPB(metadatas)
	resp.Records = int64(len(metadatas))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

func GetMetadataDetail(c *gin.Context) {
	resp := &apipb.GetMetadataDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := curdmodel.GetMetadataById(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = curdmodel.MetadataToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

func GetMetadataFields(c *gin.Context) {
	resp := &apipb.GetMetadataFieldsResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := curdmodel.GetMetadataFieldByMDId(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = curdmodel.MetadataFieldsToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

func GetMetadataTree(c *gin.Context) {
	resp := &apipb.QueryMetadataResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.QueryMetadataRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的元数据
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	data, records, err := curdmodel.GetMetadataTree(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Records = records
		resp.Data = curdmodel.MetadatasToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

func CopyMetadata(c *gin.Context) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}

	req := &apipb.DelRequest{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	err = curdmodel.CopyMetadata(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ImportMetadata
// @Summary 导入
// @Description 导入
// @Tags Metadata管理
// @Accept  mpfd
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param files formData file true "要上传的文件"
// @Success 200 {object} apipb.CommonResponse
// @Router /api//metadata/import [post]
func ImportMetadata(c *gin.Context) {
	resp := &apipb.QueryMetadataResponse{
		Code: apipb.Code_Success,
	}
	//从Metadata中读取文件
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

	var list []*apipb.MetadataInfo
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
		md := curdmodel.PBToMetadata(f)
		err = curdmodel.UpdateMetadata(md)
		if err == gorm.ErrRecordNotFound {
			err = curdmodel.CreateMetadata(md)
		}
		if err != nil {
			failCount++
			fmt.Println(err)
		} else {
			successCount++
			gen.AddMetadata(md)
		}
	}
	resp.Message = fmt.Sprintf("导入成功数量:%d,导入失败数量:%d", successCount, failCount)
	c.JSON(http.StatusOK, resp)
}

// ExportMetadata godoc
// @Summary 导出
// @Description 导出
// @Tags Metadata管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param ids query []string false "IDs"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/curd/metadata/export [get]
func ExportMetadata(c *gin.Context) {
	req := &apipb.QueryMetadataRequest{}
	resp := &apipb.QueryMetadataResponse{
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
	curdmodel.QueryMetadata(req, resp, true)
	if resp.Code != apipb.Code_Success {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment;filename=Metadata.json")
	c.Header("Content-Transfer-Encoding", "binary")
	buf, _ := json.Marshal(resp.Data)
	c.Writer.Write(buf)
}

func RegisterMetadataRouter(r *gin.Engine) {
	mdGroup := r.Group("/api/curd/metadata")
	mdGroup.POST("add", AddMetadata)
	mdGroup.PUT("update", UpdateMetadata)
	mdGroup.GET("query", QueryMetadata)
	mdGroup.DELETE("delete", DeleteMetadata)
	mdGroup.GET("all", GetAllMetadata)
	mdGroup.GET("detail", GetMetadataDetail)
	mdGroup.GET("tree", GetMetadataTree)
	mdGroup.GET("fields", GetMetadataFields)
	mdGroup.POST("import", ImportMetadata)
	mdGroup.POST("copy", CopyMetadata)
	mdGroup.GET("export", ExportMetadata)
}
