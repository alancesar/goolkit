package main

import (
	"context"
	"fmt"
	"time"
)

type MyObserver struct {
	name string
}

func NewMyObserver(name string) *MyObserver {
	return &MyObserver{
		name: name,
	}
}

func (m MyObserver) Start(_ context.Context, req string) {
	fmt.Printf("calling %s with %v\n", m.name, req)
}

func (m MyObserver) Success(_ context.Context, res time.Time) {
	fmt.Printf("called %s succesfully and got %v\n", m.name, res)
}

func (m MyObserver) Error(_ context.Context, err error) {
	fmt.Printf("some error happened in %s: %s\n", m.name, err)
}

func (m MyObserver) Finish(_ context.Context, duration time.Duration) {
	fmt.Printf("%s call took %dms\n", m.name, duration.Milliseconds())
}
