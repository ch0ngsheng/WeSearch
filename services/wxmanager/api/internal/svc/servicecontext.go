package svc

import (
	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
