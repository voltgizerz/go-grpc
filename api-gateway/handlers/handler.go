package handlers

import (
	"github.com/api-gateway/config"
	"github.com/api-gateway/services"
	"github.com/sirupsen/logrus"
)

// Handler - .
type Handler struct {
	Service services.Service
	Log     *logrus.Logger
}

// NewHandler - init handler.
func NewHandler(service *config.GrpcClient, logSetup *logrus.Logger) *Handler {
	return &Handler{
		Service: services.Service{
			OrderGRPC: service.OrderClient,
			UserGRPC:  service.UserClient,
			Log:       logSetup,
		},
		Log: logSetup,
	}
}
