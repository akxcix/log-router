package main

import "github.com/go-chi/chi/v5"

func (a *Application) AddRouter() {
	r := chi.NewRouter()
	h := NewHandler()

	r.Get("/healthcheck", h.handleHealthcheck)

	a.Router = r
}
