package usecaseprovider

import (
	repoprovider "github.com/pandhu/hehemock/app/providers/repository"
	healthzusecase "github.com/pandhu/hehemock/app/usecases/healthz"
)

// Usecase Usecase provider struct
type Usecase struct {
	HealthzUsecase *healthzusecase.HealthzUsecase
}

// InitUsecase initialize all service provider
func InitUsecase(repo *repoprovider.Repo) *Usecase {

	return &Usecase{
		HealthzUsecase: healthzusecase.NewHealthzUsecase(repo.Healthz),
	}
}
