package services

import (
	"github.com/sirupsen/logrus"
	pb "github.com/voltgizerz/public-grpc/user/gen"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedUserServiceServer
	Log *logrus.Logger
}
