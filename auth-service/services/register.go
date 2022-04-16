package services

import (
	"context"
	"log"

	"github.com/auth-service/pb"
)


// Register - register user.
func (s *Server) Register(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.LoginResponse

	return &res, nil
}
