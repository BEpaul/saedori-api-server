package repository

import "sync"

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	Dashboard *DashboardRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			Dashboard: newDashboardRepository(),
		}
	})

	return repositoryInstance
}
