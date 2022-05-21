package bar

import "context"

type Reader interface {
	Read(_ context.Context, name string) (string, error)
}
