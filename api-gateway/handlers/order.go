package handlers

import (
	"net/http"
)

func (h *Handler) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("order_id")
		res, err := h.Service.GetOrder(orderID)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		SuccesResponse(w, http.StatusOK, "Success", res)
		return
	}
}
