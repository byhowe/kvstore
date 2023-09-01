package kvstorage

import "github.com/byhowe/memvault/src/internal/kverror"

func (ms *memoryStorage) Update(key string, value any) (any, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if !ms.keyExists(key) {
		return nil, kverror.ErrKeyNotFound.AddData("`" + key + "` does not exist")
	}

	ms.db[key] = value
	return value, nil
}
