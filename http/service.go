package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/CloudSilk/curd/gen"
	ctmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/archive"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	ucm "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

func AddDefaultCodeFile(req *apipb.ServiceInfo) {
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "db.go",
		Dir:        "model",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-model-dbinit"},
		Params:     req.Params,
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "main.go",
		Dir:        "",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-go-main"},
		Params:     req.Params,
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       req.Package + "_common.proto",
		Dir:        "proto",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-proto-common"},
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "router.go",
		Dir:        "http",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-http-router-go"},
		Params:     req.Params,
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "Makefile",
		Dir:        "",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-go-makefile"},
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "Dockerfile",
		Dir:        "",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-docker-golang"},
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       ".gitignore",
		Dir:        "",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-go-gitignore"},
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "go.mod",
		Dir:        "",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-go-mod"},
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "gen.sh",
		Dir:        "proto",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-proto-gen"},
	})
	req.CodeFiles = append(req.CodeFiles, &apipb.CodeFileInfo{
		Name:       "docs.go",
		Dir:        "docs",
		Package:    req.Package,
		MetadataID: "4",
		Body:       []string{"default-go-docs"},
	})
}

func AddService(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.ServiceInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Service请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能创建其他租户的服务
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	AddDefaultCodeFile(req)
	err = ctmodel.CreateService(ctmodel.PBToService(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateService(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.ServiceInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Service请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能更新其他租户的服务
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = ctmodel.UpdateService(ctmodel.PBToService(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func QueryService(c *gin.Context) {
	req := &apipb.QueryServiceRequest{}
	resp := &apipb.QueryServiceResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的服务
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	ctmodel.QueryService(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

func DeleteService(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建API请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ctmodel.DeleteService(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func GetServiceDetail(c *gin.Context) {
	resp := &apipb.GetServiceDetailResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := ctmodel.GetServiceByID(idStr, false)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		data.Sort()
		resp.Data = ctmodel.ServiceToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

func CopyService(c *gin.Context) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	err := ctmodel.CopyService(idStr)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func GetAllService(c *gin.Context) {
	resp := &apipb.QueryServiceResponse{
		Code: apipb.Code_Success,
	}
	req := &apipb.QueryServiceRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的服务
	tenantID := ucm.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	apis, err := ctmodel.GetAllService(req)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = ctmodel.ServicesToPB(apis)
	resp.Records = int64(len(apis))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

func GenServiceCode(c *gin.Context) {
	resp := &apipb.GenServiceResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	gen.GenService(idStr, false, resp)
	c.JSON(http.StatusOK, resp)
}

func DownloadServiceCode(c *gin.Context) {
	resp := &apipb.GenServiceResponse{
		Code: apipb.Code_Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	gen.GenService(idStr, true, resp)
	if resp.Code == apipb.Code_Success {
		defer os.RemoveAll(resp.Dir)
		filePath, err := archive.PackageFolder(resp.Dir, "")
		if err != nil {
			resp.Code = apipb.Code_InternalServerError
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer os.Remove(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			resp.Code = apipb.Code_InternalServerError
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		c.Header("Content-Type", "application/octet-stream")

		c.Header("Content-Disposition", "attachment;filename="+resp.Data.Name+"-code.zip")
		c.Header("Content-Transfer-Encoding", "binary")
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.JSON(http.StatusOK, resp)
}

func AddCodeFile(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.QuickAddCodeFileRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Service请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	md, err := gen.GetMetadataById(req.MetadataID)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	if md == nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	var codeFiles []*ctmodel.CodeFile
	codeFiles = append(codeFiles, addModelCodeFile(req, md), addHTTPCodeFile(req, md), addProtoCodeFile(req, md), addConfigAPICodeFile(req, md), addConfigMenuCodeFile(req, md), addConfigPageCodeFile(req, md))
	if req.NeedProvider {
		codeFiles = append(codeFiles, addProviderCodeFile(req, md))
	}
	err = ctmodel.AddCodeFiles(codeFiles)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func addModelCodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	if req.IsTree {
		tpls = []string{"default-model-create-tree", "default-model-update-tree", "default-model-gettree", "default-model-delete3", "default-model-query", "default-model-getbyid", "default-model-struct", "default-model-pbtostruct", "default-model-convert-children"}
	} else {
		switch req.Type {
		case 1:
			tpls = []string{"default-model-create-simple", "default-model-update-simple", "default-model-delete3", "default-model-query", "default-model-getbyid", "default-model-getall", "default-model-struct", "default-model-pbtostruct", "default-model-convert-children"}
		case 2:
			tpls = []string{"default-model-create", "default-model-update-children", "default-model-delete3", "default-model-query", "default-model-getbyid", "default-model-getall", "default-model-copy", "default-model-enable", "default-model-struct", "default-model-pbtostruct", "default-model-convert-children"}
		case 3:
			tpls = []string{"default-model-create", "default-model-update-children", "default-model-delete3", "default-model-query", "default-model-getbyid", "default-model-getall", "default-model-copy", "default-model-enable", "default-model-struct", "default-model-pbtostruct", "default-model-convert-children"}
		}
	}

	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.LowerSnakeCase(md.Name) + ".go",
		Dir:        "model",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "default-model-import",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func addHTTPCodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	if req.IsTree {
		tpls = []string{"default-http-create", "default-http-update", "default-http-delete", "default-http-query", "default-http-getdetail", "default-http-gettree", "default-http-router-tree"}
	} else {
		switch req.Type {
		case 1:
			tpls = []string{"default-http-create", "default-http-update", "default-http-delete", "default-http-query", "default-http-getdetail", "default-http-getall", "default-http-router-simple"}
		case 2:
			tpls = []string{"default-http-create", "default-http-update", "default-http-delete", "default-http-query", "default-http-getdetail", "default-http-getall", "default-http-copy", "default-http-enable", "default-http-router-complex"}
		case 3:
			tpls = []string{"default-http-create", "default-http-update", "default-http-delete", "default-http-query", "default-http-getdetail", "default-http-getall", "default-http-copy", "default-http-enable", "default-http-import-data", "default-http-export", "default-http-router-complex"}
		}
	}

	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.LowerSnakeCase(md.Name) + ".go",
		Dir:        "http",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "default-http-import",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func addProtoCodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	if req.IsTree {
		tpls = append(tpls, "default-proto-tree")
	} else {
		switch req.Type {
		case 1:
			tpls = []string{"default-proto"}
		case 2:
			tpls = []string{"default-proto-complex"}
		case 3:
			tpls = []string{"default-proto-complex"}
		}
	}

	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.LowerSnakeCase(md.Name) + ".proto",
		Dir:        "proto",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func addProviderCodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	if req.IsTree {
		tpls = []string{"default-provider-curd-tree"}
	} else {
		switch req.Type {
		case 1:
			tpls = []string{"default-provider-curd"}
		case 2:
			tpls = []string{"default-provider-curd-complex"}
		case 3:
			tpls = []string{"default-provider-curd-complex"}
		}
	}

	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.LowerSnakeCase(md.Name) + ".go",
		Dir:        "provider",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func addConfigAPICodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	switch req.Type {
	case 1:
		tpls = []string{"default-config-api"}
	case 2:
		tpls = []string{"default-config-api"}
	case 3:
		tpls = []string{"default-config-api"}
	}
	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.CamelName(md.Name) + "API.json",
		Dir:        "deployment",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func addConfigMenuCodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	switch req.Type {
	case 1:
		tpls = []string{"default-config-menu"}
	case 2:
		tpls = []string{"default-config-menu"}
	case 3:
		tpls = []string{"default-config-menu"}
	}
	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.CamelName(md.Name) + "Menu.json",
		Dir:        "deployment",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func addConfigPageCodeFile(req *apipb.QuickAddCodeFileRequest, md *ctmodel.Metadata) *ctmodel.CodeFile {
	var tpls []string
	switch req.Type {
	case 1:
		tpls = []string{"default-config-page"}
	case 2:
		tpls = []string{"default-config-page"}
	case 3:
		tpls = []string{"default-config-page"}
	}
	codeFile := &ctmodel.CodeFile{
		ServiceID:  req.ServiceID,
		Name:       gen.CamelName(md.Name) + "Page.json",
		Dir:        "deployment",
		Package:    req.Package,
		MetadataID: req.MetadataID,
		Start:      "",
		Body:       ctmodel.ObjectToJsonString(tpls),
	}
	return codeFile
}

func RegisterServiceRouter(r *gin.Engine) {
	g := r.Group("/api/curd/service")

	g.POST("add", AddService)
	g.PUT("update", UpdateService)
	g.GET("query", QueryService)
	g.DELETE("delete", DeleteService)
	g.GET("all", GetAllService)
	g.GET("detail", GetServiceDetail)
	g.POST("copy", CopyService)
	g.GET("code", GenServiceCode)
	g.GET("download", DownloadServiceCode)
	// g.POST("codefile/add", AddCodeFile)
}
