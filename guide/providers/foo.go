package providers

type Foo struct {
	X int
}

// NewFoo has no external dependencies
func NewFoo() Foo {
	return Foo{X: 42}
}
