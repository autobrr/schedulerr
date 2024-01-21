package webhook

import (
	"encoding/json"
	"net/http"
	"time"
)

type HourBlock struct {
	Hour    int
	Enabled bool
}

// WeeklyScheduler represents the schedule for a whole week, with 1-hour blocks for each day.
type WeeklyScheduler struct {
	Sunday    [24]HourBlock
	Monday    [24]HourBlock
	Tuesday   [24]HourBlock
	Wednesday [24]HourBlock
	Thursday  [24]HourBlock
	Friday    [24]HourBlock
	Saturday  [24]HourBlock
}

func NewWeeklyScheduler() *WeeklyScheduler {
	ws := &WeeklyScheduler{}
	for i := 0; i < 24; i++ {
		ws.Sunday[i].Hour = i
		ws.Monday[i].Hour = i
		ws.Tuesday[i].Hour = i
		ws.Wednesday[i].Hour = i
		ws.Thursday[i].Hour = i
		ws.Friday[i].Hour = i
		ws.Saturday[i].Hour = i
	}
	return ws
}

type Payload struct {
	Time string `json:"time"` // RFC3339
}

func (ws *WeeklyScheduler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	pt, err := time.Parse(time.RFC3339, payload.Time)
	if err != nil {
		http.Error(w, "Invalid time format", http.StatusBadRequest)
		return
	}

	currentHour := pt.Hour()
	dayOfWeek := pt.Weekday()

	// Default to enabled unless disabled
	enabled := true

	// Check if the current hour is explicitly enabled or disabled
	switch dayOfWeek {
	case time.Sunday:
		enabled = ws.Sunday[currentHour].Enabled
	case time.Monday:
		enabled = ws.Monday[currentHour].Enabled
	case time.Tuesday:
		enabled = ws.Tuesday[currentHour].Enabled
	case time.Wednesday:
		enabled = ws.Wednesday[currentHour].Enabled
	case time.Thursday:
		enabled = ws.Thursday[currentHour].Enabled
	case time.Friday:
		enabled = ws.Friday[currentHour].Enabled
	case time.Saturday:
		enabled = ws.Saturday[currentHour].Enabled
	}

	if enabled {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Hour block not enabled", http.StatusForbidden)
	}
}
