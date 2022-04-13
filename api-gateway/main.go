package main

import (
	"log"
	"net/http"
	"os"

	"github.com/api-gateway/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var defaultPort = ":1000"

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
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("ORDER_GRPC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to order grpc: %v", err)
	}
	defer conn.Close()

	h := handlers.NewServiceClient(conn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Gateway Microservices Felix"))
	})

	r.Route("/api/orders", func(r chi.Router) {
		r.Get("/", h.GetOrders())
	})

	log.Printf("Api Gateway listening at http://localhost%s", defaultPort)
	http.ListenAndServe(defaultPort, r)

}
