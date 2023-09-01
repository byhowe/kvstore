package memvaultservice

import (
	"context"
	"fmt"
)

func (s *memVaultService) Update(ctx context.Context, sr *UpdateRequest) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		value, err := s.storage.Update(sr.Key, sr.Value)
		if err != nil {
			return nil, fmt.Errorf("memvaultservice update error: %w", err)
		}
		return &ItemResponse{
			Key:   sr.Key,
			Value: value,
		}, nil
	}
}
