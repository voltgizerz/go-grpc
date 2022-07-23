package services

import (
	"context"
	"log"

	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Validate - validate user.
func (s *Server) Validate(ctx context.Context, in *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.ValidateResponse

	return &res, nil
}
