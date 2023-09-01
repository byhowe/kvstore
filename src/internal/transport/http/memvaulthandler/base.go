package memvaulthandler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/byhowe/memvault/src/internal/service/memvaultservice"
	"github.com/byhowe/memvault/src/internal/transport/http/basehttphandler"
)

var _ MemVaultHTTPHandler = (*memVaultHandler)(nil) // compile time proof

// MemVaultHTTPHandler defines /store/ http handler behaviours.
type MemVaultHTTPHandler interface {
	Set(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
}

type memVaultHandler struct {
	basehttphandler.Handler

	service memvaultservice.MemVaultService
}

// StoreHandlerOption represents store handler option type.
type StoreHandlerOption func(*memVaultHandler)

// WithService sets service option.
func WithService(srvc memvaultservice.MemVaultService) StoreHandlerOption {
	return func(s *memVaultHandler) {
		s.service = srvc
	}
}

// WithContextTimeout sets handler context cancel timeout.
func WithContextTimeout(d time.Duration) StoreHandlerOption {
	return func(s *memVaultHandler) {
		s.Handler.CancelTimeout = d
	}
}

// WithServerEnv sets handler server env.
func WithServerEnv(env string) StoreHandlerOption {
	return func(s *memVaultHandler) {
		s.Handler.ServerEnv = env
	}
}

// WithLogger sets handler logger.
func WithLogger(l *slog.Logger) StoreHandlerOption {
	return func(s *memVaultHandler) {
		s.Handler.Logger = l
	}
}

// New instantiates new memVaultHandler instance.
func New(options ...StoreHandlerOption) MemVaultHTTPHandler {
	kvsh := &memVaultHandler{
		Handler: basehttphandler.Handler{},
	}

	for _, o := range options {
		o(kvsh)
	}

	return kvsh
}

func e(msg string) map[string]string {
	return map[string]string{"error": msg}
}
