package kvstorage

import "github.com/byhowe/memvault/src/internal/kverror"

func (ms *memoryStorage) Delete(key string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if !ms.keyExists(key) {
		return kverror.ErrKeyNotFound.AddData("`" + key + "` does not exist")
	}

	delete(ms.db, key)

	return nil
}
