package services

import (
	"context"
	"math/rand"
	"time"
	pb "github.com/voltgizerz/public-grpc/user/gen"
)

// GetUser - returns a user.
func (s *Service) GetUser(uID string) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userID := int64(rand.Intn(1000))
	res, err := s.UserGRPC.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return res, nil
}
