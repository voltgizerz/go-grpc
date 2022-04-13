package handlers

import (
	"net/http"

	"github.com/api-gateway/services"
)

func (h *ServiceClient) GetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := services.Service{OrderSC: h.OrderSC}
		res, err := service.GetOrders()
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		SuccesResponse(w, http.StatusOK, "Success", res)
		return
	}
}
