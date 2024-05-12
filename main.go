package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	curdconfig "github.com/CloudSilk/curd/config"
	"github.com/CloudSilk/curd/docs"
	"github.com/CloudSilk/curd/gen"
	"github.com/CloudSilk/curd/http"
	"github.com/CloudSilk/curd/model"
	"github.com/CloudSilk/curd/provider"
	"github.com/CloudSilk/curd/service"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/utils"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	middleware.InitIdentity()
	service.Init()
	config.SetProviderService(&provider.PageProvider{})
	config.SetProviderService(&provider.MetadataProvider{})
	config.SetProviderService(&provider.CellProvider{})
	// config.SetProviderService(&provider.FileTemplateProvider{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	configCenter := config.GetRootConfig().ConfigCenter
	nacosAddr := configCenter.Address
	list := strings.Split(nacosAddr, ":")
	port, err := strconv.ParseUint(list[1], 10, 64)
	if err != nil {
		panic(err)
	}
	curdconfig.Init(configCenter.Namespace, list[0], port, configCenter.Username, configCenter.Password)
	constants.SetPlatformTenantID(curdconfig.DefaultConfig.PlatformTenantID)

	model.Init(curdconfig.DefaultConfig.Mysql, curdconfig.DefaultConfig.Debug)
	fmt.Println("started server")
	gen.LoadCache()
	Start(48081)
}

func Start(port int) {
	docs.SwaggerInfo.Title = "CURD API"
	docs.SwaggerInfo.Description = "This is a CURD server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	if os.Getenv("CURD_DISABLE_AUTH") != "true" {
		r.Use(middleware.AuthRequiredWithRPC)
	}
	r.Use(utils.Cors())
	http.RegisterRouter(r)
	r.GET("/swagger/curd/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%d", port))
}
