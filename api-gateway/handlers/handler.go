package handlers

import (
	"github.com/api-gateway/pb"
	"github.com/api-gateway/services"
	"google.golang.org/grpc"
)

type Handler struct {
	Service services.Service
}

func NewHandler(conn *grpc.ClientConn) *Handler {
	return &Handler{
		Service: services.Service{
			OrderSC: pb.NewOrderServiceClient(conn),
		},
	}
}
