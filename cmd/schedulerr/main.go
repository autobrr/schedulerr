package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/autobrr/schedulerr/scheduler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var configPath string

func findConfigFile() (string, error) {
	if configPath != "" {
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}
		return "", fmt.Errorf("specified config file not found: %s", configPath)
	}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		configDir := filepath.Join(homeDir, ".config", "schedulerr", "config.yaml")
		if _, err := os.Stat(configDir); err == nil {
			return configDir, nil
		}
	}

	if _, err := os.Stat("config.yaml"); err == nil {
		return "config.yaml", nil
	}

	return "", fmt.Errorf("no config file found in standard locations")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the schedulerr HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		sched := scheduler.NewWeeklyScheduler()

		configFile, err := findConfigFile()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to find configuration file")
		}

		if err := sched.LoadConfigFromYAML(configFile); err != nil {
			log.Fatal().Err(err).Msg("Failed to load YAML configuration")
		}
		log.Info().Str("configPath", configFile).Msg("Loaded configuration file")

		http.HandleFunc("/webhook", sched.WebhookHandler)

		log.Info().Str("service", "schedulerr").Msg("Service has started on :8585")

		if err := http.ListenAndServe(":8585", nil); err != nil {
			log.Fatal().Err(err).Msg("Failed to start the server")
		}
	},
}

var rootCmd = &cobra.Command{
	Use:               "schedulerr",
	Short:             "schedulerr - A scheduling service",
	DisableAutoGenTag: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to config YAML file")

	rootCmd.AddGroup(&cobra.Group{
		ID:    "main",
		Title: "Main Commands:",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "other",
		Title: "Other Commands:",
	})

	rootCmd.SetHelpCommandGroupID("other")
	serveCmd.GroupID = "main"
	rootCmd.AddCommand(serveCmd)
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05 MST",
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}
