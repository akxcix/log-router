package main

import "github.com/go-chi/chi/v5"

func CreateRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", h.handleHealthcheck)
	r.Post("/log", h.handleLog)
	r.NotFound(h.notFound)

	return r
}
