package services

import "github.com/order-service/pb"

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedOrderServiceServer
}
