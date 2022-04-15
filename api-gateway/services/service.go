package services

import "github.com/api-gateway/pb"

// Service - .
type Service struct {
	OrderSC pb.OrderServiceClient
	UserSC  pb.UserServiceClient
}
