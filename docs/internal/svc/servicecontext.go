package svc

import (
	"github.com/localrivet/templwind/docs/internal/config"
)

type ServiceContext struct {
	Config *config.Config
	// put other context fields that you need here
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: cfg,
	}
}
