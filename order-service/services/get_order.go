package services

import (
	"context"
	"log"
	"math/rand"

	"github.com/bxcodec/faker"
	pb "github.com/voltgizerz/public-grpc/order/gen"
)

var status = []string{"Pending", "Processing", "Shipped", "Delivered"}

// GetOrder - get single fake data order.
func (s *Server) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	log.Printf("Received: %v", in)
	var res pb.GetOrderResponse
	err := faker.FakeData(&res)
	if err != nil {
		return nil, err
	}

	res.FirstName = faker.FirstName
	res.LastName = faker.LastName
	res.OrderStatus = status[rand.Intn(len(status))]
	res.OrderId = in.OrderId
	return &res, nil
}
