package services

import (
	"github.com/sirupsen/logrus"
	oClient "github.com/voltgizerz/public-grpc/order/gen"
	uClient "github.com/voltgizerz/public-grpc/user/gen"
)

// Service - .
type Service struct {
	OrderGRPC oClient.OrderServiceClient
	UserGRPC  uClient.UserServiceClient
	Log 		 *logrus.Logger
}

