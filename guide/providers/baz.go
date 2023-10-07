package providers

import (
	"context"
	"errors"
)

type Baz struct {
	X int
}

// NewBaz requires context, depends on Bar and can return an error
func NewBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("no zero values allowed")
	}

	return Baz{X: bar.X}, nil
}
