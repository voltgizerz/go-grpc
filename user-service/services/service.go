package services

import pb "github.com/voltgizerz/public-grpc/user/gen"

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedUserServiceServer
}
