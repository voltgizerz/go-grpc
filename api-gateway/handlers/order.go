package handlers

import (
	"net/http"
)

// GetOrder - is a function handler for GET /api/orders/{order_id}.
func (h *Handler) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("order_id")
		res, err := h.Service.GetOrder(orderID)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if res == nil {
			ErrorResponse(w, http.StatusNotFound, "Order not found")
			return
		}

		SuccesResponse(w, http.StatusOK, "Success", res)
	}
}
