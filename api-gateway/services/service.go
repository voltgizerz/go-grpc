package services

import (
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
