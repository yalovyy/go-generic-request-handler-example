package bar

import (
	"context"
	"fmt"
)

type GetRequest struct {
	Id string
}

func (b *Bar) Get(ctx context.Context, req GetRequest) (any, error) {
	resp, err := b.storage.Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("Bar: %s", resp), nil
}
