package cmd

import (
	"github.com/bestkkii/saedori-api-server/internal/config"
	"github.com/bestkkii/saedori-api-server/internal/repository"
	"github.com/bestkkii/saedori-api-server/internal/router"
	"github.com/bestkkii/saedori-api-server/internal/service"
)

type Cmd struct {
	config     *config.Config
	repository *repository.Repository
	service    *service.Service
	router     *router.Router
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.router = router.NewRouter(c.service)
	c.router.ServerStart(c.config.Server.Port)

	return c
}
