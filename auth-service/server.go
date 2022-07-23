package main

import (
	"net"
	"os"

	"github.com/auth-service/config"
	"github.com/auth-service/repository"
	"github.com/auth-service/services"
	pb "github.com/voltgizerz/public-grpc/auth/gen"
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
	jwt := services.NewJwtWrapper()
	userRepo := repository.NewUserRepository()

	pb.RegisterAuthServiceServer(s, &services.Server{
		Jwt:      jwt,
		UserRepo: userRepo,
		Log:      log,
	})

	log.Printf("Server gRPC auth listening at http://localhost%s", defaultPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
