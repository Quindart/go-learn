package app

import (
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/service"
)

type DemoModule struct {
	routes routes.Route
}

func NewDemoModule() *DemoModule {
	demoRepo := repository.NewDemoRepository()
	demoService := service.NewDemoService(demoRepo)
	demoHandler := handler.NewDemoHandler(demoService)

	demoRouter := routes.NewDemoRoutes(demoHandler)

	return &DemoModule{
		routes: demoRouter,
	}
}

func (m *DemoModule) Routes() routes.Route {
	return m.routes
}
