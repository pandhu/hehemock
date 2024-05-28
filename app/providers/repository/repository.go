package repoprovider

import (
	healthzrp "github.com/pandhu/hehemock/app/repositories/healthz"
	"gorm.io/gorm"
)

// Repo Repository struct
type Repo struct {
	Healthz healthzrp.HealthzRepository
}

// InitRepo initialize all repositories
func InitRepo(db *gorm.DB) *Repo {
	// conf := config.All()

	return &Repo{
		Healthz: healthzrp.NewHealthzRepository(db),
	}
}
