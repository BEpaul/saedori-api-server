package cmd

import (
	"github.com/bestkkii/saedori-api-server/internal/config"
	"github.com/bestkkii/saedori-api-server/internal/repository"
	"github.com/bestkkii/saedori-api-server/internal/router"

	"github.com/bestkkii/saedori-api-server/internal/scheduler"
	"github.com/bestkkii/saedori-api-server/internal/service"
)

type Cmd struct {
	config     *config.Config
	repository *repository.Repository
	service    *service.Service
	router     *router.Router
}

func NewCmd() *Cmd {
	c := &Cmd{
		config: config.NewConfig(),
	}

	dashboardRepo := repository.NewRepository().Dashboard
	dashboard := &scheduler.Dashboard{DashboardRepository: dashboardRepo}
	dashboard.StartScheduler()

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.router = router.NewRouter(c.service)
	c.router.ServerStart(c.config.Server.Port)

	return c
}
