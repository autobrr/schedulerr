package config

import (
	"os"

	"github.com/autobrr/schedulerr/internal/scheduler"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromYAML(filePath string) (map[string][]scheduler.HourBlock, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config map[string][]scheduler.HourBlock
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return config, nil
}
