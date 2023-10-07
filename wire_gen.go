// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/andyklimenko/wire-learning/guide/providers"
	"time"
)

// Injectors from main.go:

func initializeBaz(ctx context.Context, t time.Time) (providers.Baz, error) {
	foo := providers.NewFoo()
	bar := providers.NewBar(foo)
	baz, err := providers.NewBaz(ctx, t, bar)
	if err != nil {
		return providers.Baz{}, err
	}
	return baz, nil
}

// main.go:

func main() {
	initializeBaz(context.Background(), time.Now())
}
