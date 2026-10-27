package app

import (
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/service"
)

type Demo2Module struct {
	routes routes.Route
}

func NewDemo2Module() *Demo2Module {
	demo2Repo := repository.NewDemo2Repository()
	demo2Service := service.NewDemo2Service(demo2Repo)

	demo2Handler := handler.NewDemo2Handler(demo2Service)

	demo2Route := routes.NewDemo2Router(demo2Handler)

	return &Demo2Module{
		routes: demo2Route,
	}
}

func (m *Demo2Module) Routes() routes.Route {
	return m.routes
}
