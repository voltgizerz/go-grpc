package services

import (
	"context"
	"log"

	"github.com/auth-service/pb"
)

var status = []string{"Pending", "Processing", "Shipped", "Delivered"}

// GetOrder - get single fake data order.
func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.LoginResponse

	return &res, nil
}
