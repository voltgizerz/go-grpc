package services

import (
	oClient "github.com/voltgizerz/public-grpc/order/gen"
	uClient "github.com/voltgizerz/public-grpc/user/gen"
)

// Service - .
type Service struct {
	OrderGRPC oClient.OrderServiceClient
	UserGRPC  uClient.UserServiceClient
}

