package services

import (
	"github.com/api-gateway/config"
	"github.com/sirupsen/logrus"
	aClient "github.com/voltgizerz/public-grpc/auth/gen"
	oClient "github.com/voltgizerz/public-grpc/order/gen"
	uClient "github.com/voltgizerz/public-grpc/user/gen"
)

// Service - .
type Service struct {
	OrderGRPC oClient.OrderServiceClient
	UserGRPC  uClient.UserServiceClient
	AuthGRPC  aClient.AuthServiceClient
	Log       *logrus.Logger
}

// NewService - init service.
func NewService(service *config.GRPCClient, logSetup *logrus.Logger) *Service {
	return &Service{
		OrderGRPC: service.OrderClient,
		UserGRPC:  service.UserClient,
		AuthGRPC:  service.AuthClient,
		Log:       logSetup,
	}
}
