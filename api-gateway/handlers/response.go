package handlers

import (
	"encoding/json"
	"net/http"
)

// ErrRes - is a struct for error response.
type ErrRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccRes - is a struct for success response.
type SuccRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse - is a function for error response.
func ErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrRes{
		Code:    code,
		Message: message,
	})
}

// SuccesResponse - is a function for success response.
func SuccesResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccRes{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
