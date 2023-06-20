package main

import (
	"fmt"
	"net/http"

	"github.com/akxcix/log-router/pkg/logstore"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Application struct {
	LogStore *logstore.Store
	Handlers *Handler
	Router   *chi.Mux
}

func NewApplication() *Application {
	logStore := logstore.NewStore()
	h := NewHandler(logStore)
	r := CreateRouter(h)

	app := Application{
		LogStore: logStore,
		Handlers: h,
		Router:   r,
	}

	return &app
}

func (app *Application) Serve() {
	address := fmt.Sprintf("%s:%s", "0.0.0.0", "8080")

	app.LogStore.StartProcessing()

	log.Info().Msgf("starting server, listening on %s", address)
	log.Fatal().Err(http.ListenAndServe(address, app.Router)).Msg("router crashed")
}
