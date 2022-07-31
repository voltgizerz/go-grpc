package handlers

import (
	"github.com/api-gateway/services"
	"github.com/sirupsen/logrus"
)

// Handler - .
type Handler struct {
	Service *services.Service
	Log     *logrus.Logger
}

// NewHandler - init handler.
func NewHandler(service *services.Service, logSetup *logrus.Logger) *Handler {
	return &Handler{
		Service: service,
		Log:     logSetup,
	}
}
