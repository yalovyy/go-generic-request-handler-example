package router

import (
	"context"
	"fmt"
	"ghandler/api/bar"
	"ghandler/api/foo"
	"ghandler/environment"
	"ghandler/handler"
)

type Router struct {
	routes map[string]handler.Generic
}

func New(env environment.Environment) *Router {
	storage := env.GetStorage()

	r := new(Router)
	r.routes = map[string]handler.Generic{
		"GET /bar": handler.NewGeneric(bar.New(storage).Get),
		"PUT /foo": handler.NewGeneric(foo.New(storage).Put),
	}
	return r
}

func (r *Router) Route(ctx context.Context, path string, req string) (any, error) {
	if h, ok := r.routes[path]; ok {
		return h(ctx, req)
	}
	return nil, fmt.Errorf("no handler found for '%s' path", path)
}
