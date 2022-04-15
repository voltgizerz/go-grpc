package services

import (
	"context"
	"log"

	"github.com/bxcodec/faker"
	"github.com/user-service/pb"
)

var status = []string{"Pending", "Processing", "Shipped", "Delivered"}

// GetOrder - get single fake data order.
func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.GetUserResponse

	err := faker.FakeData(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
