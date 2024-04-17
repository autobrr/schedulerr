package webhook

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/autobrr/schedulerr/internal/scheduler"
)

func WebhookHandler(ws *scheduler.WeeklyScheduler, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var currentSchedule map[string][]scheduler.HourBlock
	var err error

	if r.ContentLength > 0 {
		err = json.NewDecoder(r.Body).Decode(&currentSchedule)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	} else {
		currentSchedule = map[string][]scheduler.HourBlock{
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

	blockEnabled := true // treat blocks as enabled by default

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
