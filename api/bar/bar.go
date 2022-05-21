package bar

type Bar struct {
	storage Reader
}

func New(storage Reader) *Bar {
	return &Bar{storage}
}
