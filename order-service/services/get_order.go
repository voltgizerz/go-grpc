package services

import (
	"context"
	"log"

	"github.com/bxcodec/faker"
	"github.com/order-service/pb"
)

// GetOrder - .
func (s *Server) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.GetOrderResponse

	err := faker.FakeData(&res)
	if err != nil {
		return nil, err
	}

	res.OrderId = in.OrderId
	return &res, nil
}
