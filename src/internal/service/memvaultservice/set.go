package memvaultservice

import (
	"context"
	"fmt"
)

func (s *memVaultService) Set(ctx context.Context, sr *SetRequest) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		val, err := s.storage.Set(sr.Key, sr.Value)
		if err != nil {
			return nil, fmt.Errorf("memvaultservice set error: %w", err)
		}

		return &ItemResponse{
			Key:   sr.Key,
			Value: val,
		}, nil
	}
}
