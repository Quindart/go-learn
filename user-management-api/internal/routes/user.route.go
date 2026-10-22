package routes

import "user-management-api/internal/handler"

type UserRoutes struct {
	Handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		Handler: handler,
	}
}
