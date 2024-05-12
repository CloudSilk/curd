package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterMetadataRouter(r)
	RegisterPageRouter(r)
	RegisterCurdRouter(r)
	RegisterServiceRouter(r)
	RegisterTemplateRoute(r)
	RegisterCellRouter(r)
	RegisterFormRouter(r)
	RegisterFileTemplateRouter(r)
	RegisterFunctionalTemplateRouter(r)
	RegisterSystemObjectRouter(r)
}
