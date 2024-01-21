package main

import (
	"net/http"
	"os"

	"github.com/autobrr/schedulerr/internal/webhook"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	scheduler := webhook.NewWeeklyScheduler()

	http.HandleFunc("/webhook", scheduler.WebhookHandler)

	log.Info().Str("service", "schedulerr").Msg("Service has started on :8585")

	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
