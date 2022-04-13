package services

import (
	"context"
	"time"

	"github.com/api-gateway/pb"
)

func (s *Service) GetOrders() (*pb.GetOrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := s.OrderSC.GetOrder(ctx, &pb.GetOrderRequest{OrderId: 1767858})
	if err != nil {
		return nil, err
	}
	return res, nil
}
