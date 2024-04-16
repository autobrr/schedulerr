package webhook

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// HourBlock represents a block of one hour with enabled status.
type HourBlock struct {
	Hour    int  `json:"hour"`
	Enabled bool `json:"enabled"`
}

// WeeklyScheduler represents the schedule for a whole week, with 1-hour blocks for each day.
type WeeklyScheduler struct {
	Sunday    []HourBlock
	Monday    []HourBlock
	Tuesday   []HourBlock
	Wednesday []HourBlock
	Thursday  []HourBlock
	Friday    []HourBlock
	Saturday  []HourBlock
}

func NewWeeklyScheduler() *WeeklyScheduler {
	return &WeeklyScheduler{}
}

func (ws *WeeklyScheduler) AssignSchedule(schedule map[string][]HourBlock) {
	for day, blocks := range schedule {
		day = strings.ToLower(day)
		switch day {
		case "sunday":
			ws.Sunday = blocks
		case "monday":
			ws.Monday = blocks
		case "tuesday":
			ws.Tuesday = blocks
		case "wednesday":
			ws.Wednesday = blocks
		case "thursday":
			ws.Thursday = blocks
		case "friday":
			ws.Friday = blocks
		case "saturday":
			ws.Saturday = blocks
		}
	}
}

func (ws *WeeklyScheduler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var currentSchedule map[string][]HourBlock
	var err error

	if r.ContentLength > 0 {
		err = json.NewDecoder(r.Body).Decode(&currentSchedule)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	} else {
		// if no body is present, use config.yaml
		currentSchedule = map[string][]HourBlock{
			"sunday":    ws.Sunday,
			"monday":    ws.Monday,
			"tuesday":   ws.Tuesday,
			"wednesday": ws.Wednesday,
			"thursday":  ws.Thursday,
			"friday":    ws.Friday,
			"saturday":  ws.Saturday,
		}
	}

	currentTime := time.Now()
	currentHour := currentTime.Hour()
	dayOfWeek := strings.ToLower(currentTime.Weekday().String())

	// treat blocks as enabled by default
	blockEnabled := true

	if blocks, found := currentSchedule[dayOfWeek]; found {
		for _, block := range blocks {
			if block.Hour == currentHour {
				blockEnabled = block.Enabled
				break
			}
		}
	}

	if blockEnabled {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Hour block not enabled", http.StatusForbidden)
	}
}
