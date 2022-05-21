package foo

import (
	"context"
	"fmt"
)

type PutRequest struct {
	Id    string
	Value string
}

func (b *Foo) Put(ctx context.Context, req PutRequest) (any, error) {
	resp, err := b.storage.Write(ctx, req.Id, req.Value)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("Foo: %s", resp), nil
}
