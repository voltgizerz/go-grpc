package handlers

import (
	"github.com/api-gateway/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	OrderSC pb.OrderServiceClient
}

func NewServiceClient(conn *grpc.ClientConn) *ServiceClient {
	return &ServiceClient{
		OrderSC: pb.NewOrderServiceClient(conn),
	}
}
