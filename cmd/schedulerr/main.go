package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/autobrr/schedulerr/internal/webhook"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	configPath := flag.String("config", "", "Path to config YAML file")
	flag.Parse()

	scheduler := webhook.NewWeeklyScheduler()

	if *configPath != "" {
		config, err := webhook.LoadConfigFromYAML(*configPath)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load YAML configuration")
		}
		scheduler.AssignSchedule(config)
	}

	http.HandleFunc("/webhook", scheduler.WebhookHandler)

	log.Info().Str("service", "schedulerr").Msg("Service has started on :8585")

	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
