package healthz

import (
	"context"
)

type pinger interface {
	DB(ctx context.Context) error
}

type HealthzUsecase struct {
	Pinger pinger
}

func NewHealthzUsecase(repo pinger) *HealthzUsecase {
	return &HealthzUsecase{repo}
}

func (hu HealthzUsecase) Ping() (ok string) {
	return "OK"
}
