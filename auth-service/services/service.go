package services

import (
	"github.com/auth-service/repository"
	"github.com/sirupsen/logrus"
	pb "github.com/voltgizerz/public-grpc/auth/gen"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedAuthServiceServer
	Jwt      JwtWrapper
	UserRepo repository.UserRepository
	Log      *logrus.Logger
}
