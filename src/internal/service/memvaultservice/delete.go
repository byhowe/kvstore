package memvaultservice

import (
	"context"
	"fmt"
)

func (s *memVaultService) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := s.storage.Delete(key); err != nil {
			return fmt.Errorf("memvaultservice delete error: %w", err)
		}
		return nil
	}
}
