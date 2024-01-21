package webhook

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DaySchedule [24]HourBlock

type Config struct {
	Sunday    DaySchedule `yaml:"sunday"`
	Monday    DaySchedule `yaml:"monday"`
	Tuesday   DaySchedule `yaml:"tuesday"`
	Wednesday DaySchedule `yaml:"wednesday"`
	Thursday  DaySchedule `yaml:"thursday"`
	Friday    DaySchedule `yaml:"friday"`
	Saturday  DaySchedule `yaml:"saturday"`
}

func LoadConfig(filename string) (*WeeklyScheduler, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	ws := NewWeeklyScheduler()
	ws.Sunday = config.Sunday
	ws.Monday = config.Monday
	ws.Tuesday = config.Tuesday
	ws.Wednesday = config.Wednesday
	ws.Thursday = config.Thursday
	ws.Friday = config.Friday
	ws.Saturday = config.Saturday

	return ws, nil
}
