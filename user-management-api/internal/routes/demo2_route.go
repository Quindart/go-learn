package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type Demo2Routes struct {
	handler *handler.Demo2Handler
}

func NewDemo2Router(handler *handler.Demo2Handler) *Demo2Routes {
	return &Demo2Routes{
		handler: handler,
	}
}
func (dr *Demo2Routes) Register(r *gin.RouterGroup) {
	demo := r.Group("/demo2s")
	{
		demo.GET("", dr.handler.GetAllDemos)
		demo.POST("", dr.handler.CreateDemo)
		demo.PUT("/:uuid", dr.handler.UpdateDemo)
		demo.DELETE("/:uuid", dr.handler.DeleteDemo)
	}
}
