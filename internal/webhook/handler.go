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

func (ws *WeeklyScheduler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a map of string to slices of HourBlocks.
	var schedule map[string][]HourBlock
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	currentHour := currentTime.Hour()
	dayOfWeek := currentTime.Weekday().String()

	dayOfWeek = strings.ToLower(dayOfWeek)

	// Assume the block is enabled by default.
	blockEnabled := true

	if blocks, found := schedule[dayOfWeek]; found {
		blockEnabled = false // Set to false initially, we'll enable it if we find a matching block.

		// Iterate over the provided blocks to find a match for the current hour.
		for _, block := range blocks {
			if block.Hour == currentHour {
				blockEnabled = block.Enabled
				break
			}
		}
	}

	// If the block is enabled or not provided, return 200.
	if blockEnabled {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Hour block not enabled", http.StatusForbidden)
	}
}
