package services

import (
	"context"
	"log"

	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Register - register user.
func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.RegisterResponse

	return &res, nil
}
