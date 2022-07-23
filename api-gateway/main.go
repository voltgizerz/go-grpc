package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/api-gateway/handlers"
	"github.com/api-gateway/client"
	"github.com/joho/godotenv"

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

	clientGrpc := client.InitGRPC(ctx)

	h := handlers.NewHandler(clientGrpc)
	appRouter := h.NewRouter()

	log.Printf("API Gateway listening at http://localhost%s", defaultPort)
	http.ListenAndServe(defaultPort, appRouter)
}
