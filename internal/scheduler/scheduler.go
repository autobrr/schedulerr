package scheduler

import (
	"strings"
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
