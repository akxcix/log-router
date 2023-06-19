package main

import "github.com/rs/zerolog/log"

func main() {
	log.Logger.Info().Msg("starting server")
	app := NewApplication()
	app.Serve()
}
