package memvaultservice

import (
	"context"
	"fmt"
)

func (s *memVaultService) Get(ctx context.Context, key string) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		value, err := s.storage.Get(key)
		if err != nil {
			return nil, fmt.Errorf("memvaultservice get error: %w", err)
		}
		return &ItemResponse{
			Key:   key,
			Value: value,
		}, nil
	}
}
