package handler

import (
	"user-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Demo2Handler struct {
	service service.Demo2Service
}

func NewDemo2Handler(sv service.Demo2Service) *Demo2Handler {
	return &Demo2Handler{
		service: sv,
	}
}

func (handler *Demo2Handler) GetAllDemos(ctx *gin.Context) {
	handler.service.GetAllDemos()
}

func (handler *Demo2Handler) CreateDemo(ctx *gin.Context) {

}

func (handler *Demo2Handler) UpdateDemo(ctx *gin.Context) {

}
func (handler *Demo2Handler) DeleteDemo(ctx *gin.Context) {

}
