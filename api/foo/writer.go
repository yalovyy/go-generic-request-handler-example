package foo

import "context"

type Writer interface {
	Write(_ context.Context, name string, value string) (string, error)
}
