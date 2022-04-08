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

func (m MyObserver) Start(_ context.Context) {
	fmt.Printf("calling %s\n", m.name)
}

func (m MyObserver) Success(_ context.Context, response any) {
	fmt.Printf("called %s succesfully and got %v\n", m.name, response)
}

func (m MyObserver) Error(_ context.Context, err error) {
	fmt.Printf("some error happened in %s: %s\n", m.name, err)
}

func (m MyObserver) Finish(_ context.Context, duration time.Duration) {
	fmt.Printf("call %s took %dms", m.name, duration.Milliseconds())
}
