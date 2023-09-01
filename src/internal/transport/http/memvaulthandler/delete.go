package memvaulthandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/byhowe/memvault/src/internal/kverror"
)

func (h *memVaultHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.JSON(w, http.StatusMethodNotAllowed, e("method "+r.Method+" not allowed"))
		return
	}

	if len(r.URL.Query()) == 0 {
		h.JSON(w, http.StatusNotFound, e("`key` query param required"))
		return
	}

	keys, ok := r.URL.Query()["key"]
	if !ok {
		h.JSON(w, http.StatusNotFound, e("`key` not present"))
		return
	}

	key := keys[0]

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	if err := h.service.Delete(ctx, key); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(w, http.StatusGatewayTimeout, e(err.Error()))
			return
		}

		var kvErr *kverror.Error

		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("memvaulthandler delete", "err", clientMessage)
			}

			if kvErr == kverror.ErrKeyNotFound {
				h.JSON(w, http.StatusNotFound, e(clientMessage))
				return
			}
		}
		h.JSON(w, http.StatusInternalServerError, e(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
