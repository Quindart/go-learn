package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type DemoRoutes struct {
	handler *handler.DemoHandler
}

func NewDemoRoutes(handler *handler.DemoHandler) *DemoRoutes {
	return &DemoRoutes{
		handler: handler,
	}
}

func (demoRoute *DemoRoutes) Register(r *gin.RouterGroup) {
	demo := r.Group("/demos")
	{
		demo.GET("", demoRoute.handler.GetAllDemos)
		demo.POST("", demoRoute.handler.CreateDemo)
		demo.PUT("/:uuid", demoRoute.handler.UpdateDemo)
		demo.DELETE("/:uuid", demoRoute.handler.DeleteDemo)
	}
}
