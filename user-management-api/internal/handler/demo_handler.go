package handler

import (
	"user-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	service service.DemoService
}

func NewDemoHandler(service service.DemoService) *DemoHandler {
	return &DemoHandler{
		service: service,
	}
}

func (handler *DemoHandler) GetAllDemos(ctx *gin.Context) {
	handler.service.GetAllDemos()
}

func (handler *DemoHandler) CreateDemo(ctx *gin.Context) {

}

func (handler *DemoHandler) UpdateDemo(ctx *gin.Context) {

}
func (handler *DemoHandler) DeleteDemo(ctx *gin.Context) {

}
