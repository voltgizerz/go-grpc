package handlers

import (
	"github.com/api-gateway/pb"
	"github.com/api-gateway/services"
	"google.golang.org/grpc"
)

type Handler struct {
	Service services.Service
}

func NewHandler(orderConn *grpc.ClientConn, userConn *grpc.ClientConn) *Handler {
	return &Handler{
		Service: services.Service{
			OrderSC: pb.NewOrderServiceClient(orderConn),
			UserSC:  pb.NewUserServiceClient(userConn),
		},
	}
}
