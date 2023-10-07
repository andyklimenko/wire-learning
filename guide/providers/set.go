package providers

import "github.com/google/wire"

var Set = wire.NewSet(NewFoo, NewBar, NewBaz)
