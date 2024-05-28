package healthz

import (
	"context"

	"gorm.io/gorm"
)

type HealthzRepository struct {
	Conn *gorm.DB
}

func NewHealthzRepository(conn *gorm.DB) HealthzRepository {
	return HealthzRepository{conn}
}

func (hr HealthzRepository) DB(ctx context.Context) error {
	db, err := hr.Conn.DB()
	if err != nil {
		return err
	}
	return db.PingContext(ctx)
}
