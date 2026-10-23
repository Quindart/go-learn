package handler

import (
	"user-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (handler *UserHandler) GetAllUsers(ctx *gin.Context) {
	handler.Service.GetAllUsers()
}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {

}

func (handler *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (handler *UserHandler) DeleteUser(ctx *gin.Context) {

}
