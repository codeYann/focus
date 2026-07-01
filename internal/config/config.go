package config

import (
	"time"

	"github.com/codeYann/focus/internal/domain"
)

// DefaultFocusDuration is the default length of a focus session (25 minutes).
var DefaultFocusDuration = 25 * time.Minute

// DefaultBreakDuration is the default length of a break session (5 minutes).
var DefaultBreakDuration = 5 * time.Minute

// DefaultNotify determines whether desktop notifications are enabled by default.
var DefaultNotify = true

// Config holds the user-configurable settings for the focus timer.
type Config struct {
	FocusDuration time.Duration
	BreakDuration time.Duration
	Notify        bool
}

// Default returns a Config populated with sensible defaults.
func Default() Config {
	return Config{
		FocusDuration: DefaultFocusDuration,
		BreakDuration: DefaultBreakDuration,
		Notify:        DefaultNotify,
	}
}

// Domain converts the application-level config to a domain-level Config.
func (c Config) Domain() domain.Config {
	return domain.Config{
		Focus: c.FocusDuration,
		Break: c.BreakDuration,
	}
}
