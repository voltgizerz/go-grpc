package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// LoginRequest - is a struct for login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest - is a struct for register request.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register - is a function for register handlers.
func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.Username == "" || req.Password == "" {
			ErrorResponse(w, http.StatusBadRequest, "Username and password are required")
			return
		}

		res, err := h.Service.Register(req.Username, req.Password)
		if err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if res.Error != "" {
			ErrorResponse(w, http.StatusBadRequest, res.Error)
			return
		}

		SuccesResponse(w, http.StatusOK, "User Created", res)
	}
}

// Login - is a function for login handlers.
func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.Username == "" || req.Password == "" {
			ErrorResponse(w, http.StatusBadRequest, "Username and password are required")
			return
		}

		res, err := h.Service.Login(req.Username, req.Password)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if res.Error != "" {
			ErrorResponse(w, http.StatusUnauthorized, res.Error)
			return
		}

		SuccesResponse(w, http.StatusOK, "Login Success", res)
	}
}

// Validate - is a function for validate user handlers.
func (h *Handler) Validate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		res, err := h.Service.Validate(token)
		if err != nil {
			ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		if res.Error != "" {
			ErrorResponse(w, http.StatusBadRequest, res.Error)
			return
		}

		SuccesResponse(w, http.StatusOK, "Token is valid", res)
	}
}
