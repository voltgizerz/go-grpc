package services

import (
	"github.com/sirupsen/logrus"
	pb "github.com/voltgizerz/public-grpc/order/gen"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedOrderServiceServer
	Log *logrus.Logger
}
