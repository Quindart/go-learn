package app

import (
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/service"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRouter := routes.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRouter,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}
