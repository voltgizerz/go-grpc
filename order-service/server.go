package main

import (
	"net"
	"os"

	"github.com/order-service/config"
	"github.com/order-service/services"
	pb "github.com/voltgizerz/public-grpc/order/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var defaultPort = ":3000"

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

	pb.RegisterOrderServiceServer(s, &services.Server{
		Log: log,
	})
	log.Printf("Server gRPC order listening at http://localhost%s", defaultPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
