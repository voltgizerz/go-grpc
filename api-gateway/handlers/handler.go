package handlers

import (
	"github.com/api-gateway/client"
	"github.com/api-gateway/services"
)

// Handler - .
type Handler struct {
	Service services.Service
}

// NewHandler - init handler.
func NewHandler(service *client.GrpcClient) *Handler {
	return &Handler{
		Service: services.Service{
			OrderGRPC: service.OrderClient,
			UserGRPC:  service.UserClient,
		},
	}
}