package handlers

import (
	"net/http"

	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi"
)

// NewRouter - is a function that returns a new router.
func (h *Handler) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(logger.Logger("router", h.Log))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Gateway Microservices Felix"))
	})

	r.Route("/api/orders", func(r chi.Router) {
		r.Get("/", h.GetOrder())
	})

	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", h.GetUser())
	})

	return r
}
