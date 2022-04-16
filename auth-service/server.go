package main

import (
	"log"
	"net"
	"os"

	"github.com/auth-service/pb"
	"github.com/auth-service/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var defaultPort = ":3000"

// Init - .
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if os.Getenv("PORT") != "" {
		defaultPort = ":" + os.Getenv("PORT")
	}
}

func main() {
	Init()
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &services.Server{})
	log.Printf("Server gRPC auth listening at http://localhost%s", defaultPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
