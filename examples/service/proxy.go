package service

import (
	"context"
	"go-web-toolkit/application/observability"
	"time"
)

type MyProxyService struct {
	service *MyService
	proxy   observability.Proxy[string, time.Time]
}

func (m MyProxyService) ToTime(ctx context.Context, s string) (time.Time, error) {
	return m.proxy.Run(ctx, s, m.service.ToTime)
}

func NewMyProxyService(service *MyService, proxy observability.Proxy[string, time.Time]) *MyProxyService {
	return &MyProxyService{
		service: service,
		proxy:   proxy,
	}
}
