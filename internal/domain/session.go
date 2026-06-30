package domain

import (
	"errors"
	"time"
)

// Errors returned by Session methods.
var (
	ErrSessionNotIdle    = errors.New("sessão não está no estado inicial")
	ErrSessionNotRunning = errors.New("sessão não está em andamento")
	ErrSessionNotPaused  = errors.New("sessão não está pausada")
	ErrSessionCompleted  = errors.New("sessão já está finalizada")
)

// Session orchestrates the lifecycle of a Pomodoro focus session.
type Session struct {
	ID        string        // Unique identifier for the session.
	Config    Config        // User-configured durations.
	Phase     Phase         // Current session state.
	Paused    bool          // Whether the session is paused.
	Remaining time.Duration // Time left in the current phase.
	Timing    Timing        // Timestamps for state transitions.
}

// New creates a Session with the given ID and configuration.
// The session starts in the Idle phase with no time remaining.
func New(ID string, config Config) *Session {
	return &Session{
		ID:     ID,
		Config: config,
		Phase:  Idle,
	}
}

// IsRunning reports whether the session is in a focus or break period.
func (s Session) IsRunning() bool {
	return (s.Phase == Focus || s.Phase == Break)
}

// IsPaused reports whether the session is currently paused.
func (s Session) IsPaused() bool {
	return s.Paused
}

// IsCompleted reports whether the session has finished.
func (s Session) IsCompleted() bool {
	return s.Phase == Completed
}

func (s *Session) advance(now time.Time) error {
	s.Paused = false

	switch s.Phase {
	case Focus:
		s.Phase = Break
		s.Remaining = s.Config.Break

	case Break:
		if err := s.Timing.Finish(now); err != nil {
			return err
		}

		s.Phase = Completed
		s.Remaining = 0
	}

	return nil
}

// Tick decrements the remaining time by delta and advances to the next phase when time expires.
func (s *Session) Tick(delta time.Duration, now time.Time) error {
	if !s.IsRunning() || s.IsPaused() {
		return nil
	}

	if delta <= 0 {
		return nil
	}

	s.Remaining -= delta

	if s.Remaining <= 0 {
		s.Remaining = 0
		return s.advance(now)
	}

	return nil
}

// Start begins the session, transitioning from Idle to Focus.
func (s *Session) Start(now time.Time) error {
	if s.Phase != Idle {
		return ErrSessionNotIdle
	}

	if err := s.Timing.Start(now); err != nil {
		return err
	}

	s.Phase = Focus
	s.Remaining = s.Config.Focus

	return nil
}

// Pause suspends an active session.
func (s *Session) Pause(now time.Time) error {
	if !s.IsRunning() {
		return ErrSessionNotRunning
	}

	if err := s.Timing.Pause(now); err != nil {
		return err
	}

	s.Paused = true

	return nil
}

// Resume continues a paused session.
func (s *Session) Resume() error {
	if !s.IsPaused() {
		return ErrSessionNotPaused
	}

	if err := s.Timing.Resume(); err != nil {
		return err
	}

	s.Paused = false
	return nil
}

// Finish ends the session immediately, setting it to Completed.
func (s *Session) Finish(now time.Time) error {
	if s.IsCompleted() {
		return ErrSessionCompleted
	}

	if err := s.Timing.Finish(now); err != nil {
		return err
	}

	s.Phase = Completed
	s.Remaining = 0
	return nil
}
