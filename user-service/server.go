package main

import (
	"net"
	"os"

	"github.com/user-service/config"
	"github.com/user-service/services"
	pb "github.com/voltgizerz/public-grpc/user/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var defaultPort = ":3001"

func main() {
	config.LoadENV()
	if os.Getenv("PORT") != "" {
		defaultPort = ":" + os.Getenv("PORT")
	}

	log := config.SetupLog()
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUserServiceServer(s, &services.Server{
		Log: log,
	})
	log.Printf("Server gRPC user listening at http://localhost%s", defaultPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
