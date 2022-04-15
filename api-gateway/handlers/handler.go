package handlers

import (
	"github.com/api-gateway/pb"
	"github.com/api-gateway/services"
	"google.golang.org/grpc"
)

// Handler - .
type Handler struct {
	Service services.Service
}

// NewHandler - return new Handler.
func NewHandler(orderConn *grpc.ClientConn, userConn *grpc.ClientConn) *Handler {
	return &Handler{
		Service: services.Service{
			OrderSC: pb.NewOrderServiceClient(orderConn),
			UserSC:  pb.NewUserServiceClient(userConn),
		},
	}
}
