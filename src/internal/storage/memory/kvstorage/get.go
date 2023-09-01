package kvstorage

import "github.com/byhowe/memvault/src/internal/kverror"

func (ms *memoryStorage) Get(key string) (any, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	value, ok := ms.db[key]
	if !ok {
		return nil, kverror.ErrKeyNotFound.AddData("`" + key + "` does not exist")
	}
	return value, nil
}
