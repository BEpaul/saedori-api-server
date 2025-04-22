package service

import (
	"sync"

	"github.com/bestkkii/saedori-api-server/internal/repository"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository
	Dashboard  *Dashboard
}

func NewService(repository *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: repository,
		}

		serviceInstance.Dashboard = newDashboardService(repository.Dashboard)
	})

	return serviceInstance
}
