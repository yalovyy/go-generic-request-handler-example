package foo

type Foo struct {
	storage Writer
}

func New(storage Writer) *Foo {
	return &Foo{storage}
}
