package service

import (
	"context"
	"time"
)

type MyService struct{}

func NewMyService() *MyService {
	return &MyService{}
}

func (s MyService) ToTime(_ context.Context, value string) (time.Time, error) {
	time.Sleep(500 * time.Millisecond)
	return time.Parse("2006-01-02", value)
}
