package domain

import "time"

// Config holds the durations for a Pomodoro session.
type Config struct {
	Focus time.Duration // Duration of the focus period.
	Break time.Duration // Duration of the break period.
}
