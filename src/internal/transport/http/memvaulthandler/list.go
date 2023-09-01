package memvaulthandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/byhowe/memvault/src/internal/kverror"
)

func (h *memVaultHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.JSON(w, http.StatusMethodNotAllowed, e("method "+r.Method+" not allowed"))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	serviceResponse, err := h.service.List(ctx)
	if err != nil {
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
				h.Logger.Error("memvaulthandler list", "err", clientMessage)
			}
		}

		h.JSON(
			w,
			http.StatusInternalServerError,
			e(err.Error()),
		)
		return
	}

	var handlerResponse ListResponse
	for _, item := range *serviceResponse {
		handlerResponse = append(handlerResponse, ItemResponse{
			Key:   item.Key,
			Value: item.Value,
		})
	}

	if len(handlerResponse) == 0 {
		h.JSON(
			w,
			http.StatusNotFound,
			e("nothing found"),
		)
		return
	}

	h.JSON(
		w,
		http.StatusOK,
		handlerResponse,
	)
}
