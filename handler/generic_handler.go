package handler

import (
	"context"
	"encoding/json"
)

type Generic func(ctx context.Context, req string) (any, error)

func NewGeneric[T any](api func(context.Context, T) (any, error)) Generic {
	return func(ctx context.Context, req string) (any, error) {
		var t T
		if err := json.Unmarshal([]byte(req), &t); err != nil {
			return nil, err
		}
		return api(ctx, t)
	}
}
