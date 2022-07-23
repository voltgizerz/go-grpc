package main

import (
	"log"
	"net"
	"os"

	"github.com/auth-service/services"
	"github.com/auth-service/repository"
	"github.com/joho/godotenv"
	pb "github.com/voltgizerz/public-grpc/auth/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(s)
	jwt := services.NewJwtWrapper()
	userRepo := repository.NewUserRepository()
	
	pb.RegisterAuthServiceServer(s, &services.Server{
		Jwt: jwt,
		UserRepo: userRepo,
	})

	log.Printf("Server gRPC auth listening at http://localhost%s", defaultPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
