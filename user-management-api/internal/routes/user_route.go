package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	Handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		Handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", ur.Handler.GetAllUsers)
		users.POST("", ur.Handler.CreateUser)
		users.PUT("/:uuid", ur.Handler.UpdateUser)
		users.DELETE("/:uuid", ur.Handler.DeleteUser)
	}
}
