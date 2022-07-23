package services

import (
	"context"

	"github.com/auth-service/models"
	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Register - register user.
func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.Log.Printf("Received: %v", in)

	user := models.User{
		Username: in.Username,
		Password: in.Password,
	}

	err := s.UserRepo.CreateUser(user)
	if err != nil {
		s.Log.Errorf("Error: %v", err)
		return &pb.RegisterResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: 201,
	}, nil
}
