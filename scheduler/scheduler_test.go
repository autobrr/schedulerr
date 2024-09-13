package scheduler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLoadConfigFromYAML(t *testing.T) {
	ws := NewWeeklyScheduler()

	err := ws.LoadConfigFromYAML("testdata/schedule.yaml")
	if err != nil {
		t.Fatalf("Failed to load YAML configuration: %v", err)
	}

	if len(ws.Schedule) == 0 {
		t.Fatal("Expected non-empty schedule after loading configuration")
	}

	mondayBlocks, ok := ws.Schedule["monday"]
	if !ok {
		t.Fatal("Expected schedule to contain 'monday'")
	}

	found := false
	for _, block := range mondayBlocks {
		if block.Hour == 9 && block.Enabled {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected to find an enabled hour block at 9 AM on Monday")
	}
}

func TestIsHourEnabled(t *testing.T) {
	ws := NewWeeklyScheduler()

	// set up a schedule with a disabled block at the current hour
	currentTime := time.Now()
	dayOfWeek := strings.ToLower(currentTime.Weekday().String())
	currentHour := currentTime.Hour()

	ws.Schedule = ScheduleData{
		dayOfWeek: []HourBlock{
			{Hour: currentHour, Enabled: false},
		},
	}

	enabled := ws.IsHourEnabled(ws.Schedule, currentTime)
	if enabled {
		t.Error("Expected current hour to be disabled")
	}

	// test with an enabled block
	ws.Schedule[dayOfWeek][0].Enabled = true
	enabled = ws.IsHourEnabled(ws.Schedule, currentTime)
	if !enabled {
		t.Error("Expected current hour to be enabled")
	}

	// test default behavior when no block is found
	ws.Schedule = ScheduleData{}
	enabled = ws.IsHourEnabled(ws.Schedule, currentTime)
	if !enabled {
		t.Error("Expected default behavior to be enabled when no block is found")
	}
}

func TestWebhookHandler(t *testing.T) {
	ws := NewWeeklyScheduler()

	// set up a schedule with a disabled block at the current hour
	currentTime := time.Now()
	dayOfWeek := strings.ToLower(currentTime.Weekday().String())
	currentHour := currentTime.Hour()

	ws.Schedule = ScheduleData{
		dayOfWeek: []HourBlock{
			{Hour: currentHour, Enabled: false},
		},
	}

	req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
	w := httptest.NewRecorder()
	ws.WebhookHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected status %d, got %d", http.StatusForbidden, resp.StatusCode)
	}

	ws.Schedule[dayOfWeek][0].Enabled = true
	w = httptest.NewRecorder()
	ws.WebhookHandler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// rest with custom schedule in request body
	customSchedule := ScheduleData{
		dayOfWeek: []HourBlock{
			{Hour: currentHour, Enabled: false},
		},
	}
	body, err := json.Marshal(customSchedule)
	if err != nil {
		t.Fatalf("Failed to marshal custom schedule: %v", err)
	}

	req = httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	ws.WebhookHandler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected status %d with custom schedule, got %d", http.StatusForbidden, resp.StatusCode)
	}
}
