package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Router *chi.Mux
}

func NewApplication() *Application {
	app := Application{}

	app.AddRouter()

	return &app
}

func (app *Application) Serve() {
	address := fmt.Sprintf("%s:%s", "127.0.0.1", "8080")
	log.Info().Msgf("starting server, listening on %s", address)
	log.Fatal().Err(http.ListenAndServe(address, app.Router)).Msg("router crashed")
}
