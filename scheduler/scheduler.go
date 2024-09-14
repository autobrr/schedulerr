package scheduler

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type HourBlock struct {
	Hour    int  `json:"hour" yaml:"hour"`
	Enabled bool `json:"enabled" yaml:"enabled"`
}

type ScheduleData map[string][]HourBlock

type WeeklyScheduler struct {
	Schedule ScheduleData
}

func NewWeeklyScheduler() *WeeklyScheduler {
	return &WeeklyScheduler{
		Schedule: make(ScheduleData),
	}
}

func (ws *WeeklyScheduler) AssignSchedule(schedule ScheduleData) {
	for day, blocks := range schedule {
		ws.Schedule[strings.ToLower(day)] = blocks
	}
}

func (ws *WeeklyScheduler) LoadConfigFromYAML(filePath string) error {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var config ScheduleData
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return err
	}

	ws.AssignSchedule(config)
	return nil
}

func (ws *WeeklyScheduler) IsHourEnabled(scheduleData ScheduleData, t time.Time) bool {
	currentHour := t.Hour()
	dayOfWeek := strings.ToLower(t.Weekday().String())

	if blocks, found := scheduleData[dayOfWeek]; found {
		for _, block := range blocks {
			if block.Hour == currentHour {
				return block.Enabled
			}
		}
	}
	return true // default to enabled if no specific block is found
}

func (ws *WeeklyScheduler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var scheduleData ScheduleData

	if r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(&scheduleData); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	} else {
		scheduleData = ws.Schedule
	}

	if ws.IsHourEnabled(scheduleData, time.Now()) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Hour block not enabled", http.StatusForbidden)
	}
}
