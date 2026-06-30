package config

import "time"

var (
	DefaultFocusDuration = 25 * time.Minute
	DefaultBreakDuration = 5 * time.Minute
	DefaultNotify        = true
)

type Config struct {
	FocusDuration time.Duration
	BreakDuration time.Duration
	Notify        bool
}

func Default() Config {
	return Config{
		FocusDuration: DefaultFocusDuration,
		BreakDuration: DefaultBreakDuration,
		Notify:        DefaultNotify,
	}
}
