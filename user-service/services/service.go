package services

import "github.com/user-service/pb"

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedUserServiceServer
}
