package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/api-gateway/config"
	"github.com/api-gateway/handlers"
	"github.com/api-gateway/services"
)

var defaultPort = ":1000"

func main() {
	config.LoadENV()
	if os.Getenv("PORT") != "" {
		defaultPort = ":" + os.Getenv("PORT")
	}

	log := config.SetupLog()
	log.Info("Dialing gRPC server...")

	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	clientGRPC := config.InitGRPC(ctx)
	service := services.NewService(clientGRPC, log)

	h := handlers.NewHandler(service, log)
	appRouter := h.NewRouter()

	log.Printf("API Gateway listening at http://localhost%s", defaultPort)
	http.ListenAndServe(defaultPort, appRouter)
}
