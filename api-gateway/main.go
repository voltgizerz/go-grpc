package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/api-gateway/handlers"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var defaultPort = ":1000"

// LoadENV - load env file.
func LoadENV() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	if os.Getenv("PORT") != "" {
		defaultPort = ":" + os.Getenv("PORT")
	}
}

func main() {
	LoadENV()
	log.Println("Dialing gRPC server...")

	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	connOrder, err := grpc.DialContext(ctx, os.Getenv("ORDER_GRPC"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect to order gRPC: %v", err)
	}
	defer connOrder.Close()

	connUser, err := grpc.DialContext(ctx, os.Getenv("USER_GRPC"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect to user gRPC: %v", err)
	}
	defer connUser.Close()

	h := handlers.NewHandler(connOrder, connUser)
	appRouter := h.NewRouter()

	log.Printf("API Gateway listening at http://localhost%s", defaultPort)
	http.ListenAndServe(defaultPort, appRouter)
}
