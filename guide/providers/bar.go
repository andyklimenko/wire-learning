package providers

type Bar struct {
	X int
}

// NewBar depends on Foo
func NewBar(f Foo) Bar {
	return Bar{X: -f.X}
}
