package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrRes{
		Code:    code,
		Message: message,
	})
}

func SuccesResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccRes{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
