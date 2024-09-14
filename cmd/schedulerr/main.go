package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/autobrr/schedulerr/scheduler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05 MST",
	})

	configPath := flag.String("config", "", "Path to config YAML file")
	flag.Parse()

	sched := scheduler.NewWeeklyScheduler()

	if *configPath != "" {
		if err := sched.LoadConfigFromYAML(*configPath); err != nil {
			log.Fatal().Err(err).Msg("Failed to load YAML configuration")
		}
		log.Info().Str("configPath", *configPath).Msg("Loaded configuration file")
	} else {
		log.Info().Msg("No configuration file loaded")
	}

	http.HandleFunc("/webhook", sched.WebhookHandler)

	log.Info().Str("service", "schedulerr").Msg("Service has started on :8585")

	if err := http.ListenAndServe(":8585", nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
