package services

import (
	"context"

	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Login - login user.
func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.Log.Printf("Received: %v", in)

	user, err := s.UserRepo.GetUserByUsernameAndPassword(in.Username, in.Password)
	if err != nil {
		s.Log.Errorf("Error: %v", err)
		return &pb.LoginResponse{
			Status: 401,
			Error:  err.Error(),
		}, nil
	}

	if user == nil {
		s.Log.Error("Error: user not found")
		return &pb.LoginResponse{
			Status: 401,
			Error:  "user not found.",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(*user)
	if err != nil {
		s.Log.Errorf("Error: %v", err)
		return &pb.LoginResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: 200,
		Token:  token,
	}, nil
}
