package services

import "github.com/auth-service/pb"

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedAuthServiceServer
}
