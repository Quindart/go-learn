package app

import (
	"user-management-api/internal/config"
	"user-management-api/internal/routes"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.Config
	Router *gin.Engine
}

type Module interface {
	Routes() routes.Route
}

func NewApplication(cfg *config.Config) *Application {

	r := gin.Default()

	modules := []Module{
		NewUserModule(),
		NewDemoModule(),
		NewDemo2Module(),
	}
	routes.RegisterRouter(r, getModuleRoutes(modules)...)

	return &Application{
		Config: cfg,
		Router: r,
	}
}

func (app *Application) Run() error {

	return app.Router.Run(app.Config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {

	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}
	return routeList
}
