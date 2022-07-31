package services

import (
	"context"
	"time"

	"github.com/api-gateway/utils"
	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Register - registers a user.
func (s *Service) Register(username, password string) (*pb.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer utils.RecoverPanic()

	res, err := s.AuthGRPC.Register(ctx, &pb.RegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Login - login a user.
func (s *Service) Login(username, password string) (*pb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer utils.RecoverPanic()

	res, err := s.AuthGRPC.Login(ctx, &pb.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Validate - validates a user.
func (s *Service) Validate(token string) (*pb.ValidateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer utils.RecoverPanic()

	res, err := s.AuthGRPC.Validate(ctx, &pb.ValidateRequest{
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
