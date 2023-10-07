//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/andyklimenko/wire-learning/guide/providers"
	"github.com/google/wire"
	"time"
)

// we have to pass ctx here because it's required by NewBaz provider
func initializeBaz(ctx context.Context) (providers.Baz, error) {
	wire.Build(providers.Set)
	return providers.Baz{}, nil
}

func main() {
	initializeBaz(context.Background(), time.Now())
}
