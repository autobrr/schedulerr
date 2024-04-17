package scheduler

import (
	"testing"
)

func TestAssignSchedule(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[string][]HourBlock
		expected *WeeklyScheduler
	}{
		{
			name: "Assign full schedule",
			input: map[string][]HourBlock{
				"Sunday":    {{Hour: 10, Enabled: true}},
				"Monday":    {{Hour: 11, Enabled: false}},
				"Tuesday":   {{Hour: 12, Enabled: true}},
				"Wednesday": {{Hour: 13, Enabled: false}},
				"Thursday":  {{Hour: 14, Enabled: true}},
				"Friday":    {{Hour: 15, Enabled: false}},
				"Saturday":  {{Hour: 16, Enabled: true}},
			},
			expected: &WeeklyScheduler{
				Sunday:    []HourBlock{{Hour: 10, Enabled: true}},
				Monday:    []HourBlock{{Hour: 11, Enabled: false}},
				Tuesday:   []HourBlock{{Hour: 12, Enabled: true}},
				Wednesday: []HourBlock{{Hour: 13, Enabled: false}},
				Thursday:  []HourBlock{{Hour: 14, Enabled: true}},
				Friday:    []HourBlock{{Hour: 15, Enabled: false}},
				Saturday:  []HourBlock{{Hour: 16, Enabled: true}},
			},
		},
		{
			name: "Case insensitivity",
			input: map[string][]HourBlock{
				"sUnDaY": {{Hour: 10, Enabled: true}},
			},
			expected: &WeeklyScheduler{
				Sunday: []HourBlock{{Hour: 10, Enabled: true}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			scheduler := NewWeeklyScheduler()
			scheduler.AssignSchedule(tt.input)

			// Compare each day to ensure the schedule is assigned correctly
			if diff := compareHourBlocks(scheduler.Sunday, tt.expected.Sunday); diff != "" {
				t.Errorf("Mismatch on Sunday: %s", diff)
			}
			if diff := compareHourBlocks(scheduler.Monday, tt.expected.Monday); diff != "" {
				t.Errorf("Mismatch on Monday: %s", diff)
			}
			if diff := compareHourBlocks(scheduler.Tuesday, tt.expected.Tuesday); diff != "" {
				t.Errorf("Mismatch on Tuesday: %s", diff)
			}
			if diff := compareHourBlocks(scheduler.Wednesday, tt.expected.Wednesday); diff != "" {
				t.Errorf("Mismatch on Wednesday: %s", diff)
			}
			if diff := compareHourBlocks(scheduler.Thursday, tt.expected.Thursday); diff != "" {
				t.Errorf("Mismatch on Thursday: %s", diff)
			}
			if diff := compareHourBlocks(scheduler.Friday, tt.expected.Friday); diff != "" {
				t.Errorf("Mismatch on Friday: %s", diff)
			}
			if diff := compareHourBlocks(scheduler.Saturday, tt.expected.Saturday); diff != "" {
				t.Errorf("Mismatch on Saturday: %s", diff)
			}
		})
	}
}

func TestNewWeeklyScheduler(t *testing.T) {
	t.Parallel()

	scheduler := NewWeeklyScheduler()
	if scheduler == nil {
		t.Error("NewWeeklyScheduler() returned nil, expected non-nil *WeeklyScheduler")
		return
	}

	// Check if all days are initialized and empty
	if len(scheduler.Sunday) != 0 || len(scheduler.Monday) != 0 || len(scheduler.Tuesday) != 0 ||
		len(scheduler.Wednesday) != 0 || len(scheduler.Thursday) != 0 || len(scheduler.Friday) != 0 ||
		len(scheduler.Saturday) != 0 {
		t.Error("NewWeeklyScheduler() should initialize all days with empty hour blocks")
	}
}

func compareHourBlocks(a, b []HourBlock) string {
	if len(a) != len(b) {
		return "different lengths"
	}
	for i := range a {
		if a[i] != b[i] {
			return "hour block mismatch"
		}
	}
	return ""
}
