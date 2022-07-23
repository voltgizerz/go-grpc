package services

import (
	"context"
	"log"

	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Login - login user.
func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.LoginResponse

	return &res, nil
}
