package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/api-gateway/config"
	"github.com/api-gateway/handlers"
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

	clientGrpc := config.InitGRPC(ctx)

	h := handlers.NewHandler(clientGrpc, log)
	appRouter := h.NewRouter()

	log.Printf("API Gateway listening at http://localhost%s", defaultPort)
	http.ListenAndServe(defaultPort, appRouter)
}
