package services

import 	pb "github.com/voltgizerz/public-grpc/auth/gen"

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedAuthServiceServer
}
