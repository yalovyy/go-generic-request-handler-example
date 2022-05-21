package storage

import (
	"context"
	"fmt"
)

func New(endpoint string) *Storage {
	return &Storage{endpoint}
}

type Storage struct {
	endpoint string
}

func (s *Storage) Read(_ context.Context, name string) (string, error) {
	return fmt.Sprintf("Reading from '%s' [%s]", name, s.endpoint), nil
}

func (s *Storage) Write(_ context.Context, name string, value string) (string, error) {
	return fmt.Sprintf("Writing '%s' to '%s' from [%s]", value, name, s.endpoint), nil
}
