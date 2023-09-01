package memvaultservice

import (
	"context"

	"github.com/byhowe/memvault/src/internal/storage/memory/kvstorage"
)

var _ MemVaultService = (*memVaultService)(nil)

// MemVaultService
type MemVaultService interface {
	Set(context.Context, *SetRequest) (*ItemResponse, error)
	Get(context.Context, string) (*ItemResponse, error)
	Update(context.Context, *UpdateRequest) (*ItemResponse, error)
	Delete(context.Context, string) error
	List(context.Context) (*ListResponse, error)
}

type memVaultService struct {
	storage kvstorage.Storer
}

// ServiceOption represents service option type.
type ServiceOption func(*memVaultService)

// WithStorage sets storage option.
func WithStorage(strg kvstorage.Storer) ServiceOption {
	return func(s *memVaultService) {
		s.storage = strg
	}
}

// New instantiates new service instance.
func New(options ...ServiceOption) MemVaultService {
	kvs := &memVaultService{}

	for _, o := range options {
		o(kvs)
	}

	return kvs
}
