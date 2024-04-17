package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/autobrr/schedulerr/internal/config"
	"github.com/autobrr/schedulerr/internal/scheduler"
	"github.com/autobrr/schedulerr/internal/webhook"

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

	scheduler := scheduler.NewWeeklyScheduler()
	loadConfig := config.LoadConfigFromYAML

	if *configPath != "" {
		config, err := loadConfig(*configPath)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load YAML configuration")
		}
		scheduler.AssignSchedule(config)
		log.Info().Str("configPath", *configPath).Msg("Loaded configuration file")
	} else {
		log.Info().Msg("No configuration file loaded")
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		webhook.WebhookHandler(scheduler, w, r)
	})

	log.Info().Str("service", "schedulerr").Msg("Service has started on :8585")

	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
