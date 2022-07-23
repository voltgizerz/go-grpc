package handlers

import (
	"net/http"
)

// GetUser - is a function handler for GET /api/users/{user_id}.
func (h *Handler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		res, err := h.Service.GetUser(userID)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if res == nil {
			ErrorResponse(w, http.StatusNotFound, "User not found")
			return
		}

		SuccesResponse(w, http.StatusOK, "Success", res)
	}
}
