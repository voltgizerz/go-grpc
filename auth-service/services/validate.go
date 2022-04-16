package services

import (
	"context"
	"log"

	"github.com/auth-service/pb"
)

// Validate - validate user.
func (s *Server) Validate(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.LoginResponse

	return &res, nil
}
