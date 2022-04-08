package service

import (
	"context"
	"go-web-toolkit/application/observability"
	"time"
)

type MyProxyService struct {
	service *MyService
	proxy   observability.Proxy[time.Time]
}

func (m MyProxyService) ToTime(ctx context.Context, value string) (time.Time, error) {
	return m.proxy.Run(ctx, func() (time.Time, error) {
		return m.service.ToTime(ctx, value)
	})
}

func NewMyProxyService(service *MyService, proxy observability.Proxy[time.Time]) *MyProxyService {
	return &MyProxyService{
		service: service,
		proxy:   proxy,
	}
}
