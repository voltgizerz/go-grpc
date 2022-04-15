package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter - is a function that returns a new router.
func (h *Handler) NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
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
