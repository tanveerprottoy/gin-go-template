package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/templates-go-gin/internal/app/module/content"
	"github.com/tanveerprottoy/templates-go-gin/internal/pkg/constant"
)

func RegisterContentRoutes(e *gin.Engine, version string, module *content.Module) {
	routes := e.Group(constant.ApiPattern + version + constant.ContentsPattern)
	routes.GET(constant.RootPattern, module.Handler.ReadMany)
	routes.GET(constant.RootPattern+"{id}", module.Handler.ReadOne)
	routes.POST(constant.RootPattern, module.Handler.Create)
	routes.PATCH(constant.RootPattern+"{id}", module.Handler.Update)
	routes.DELETE(constant.RootPattern+"{id}", module.Handler.Delete)
}