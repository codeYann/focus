// Package domain defines the core domain model for a Pomodoro-style focus session.
//
// It provides the foundational types and business logic for managing focus sessions,
// including session state transitions, timing tracking, and configuration.
//
// The main types are:
//   - Config: holds user-configurable durations for focus and break periods.
//   - Phase: represents the possible states of a session (Idle, Focus, Break, Completed).
//   - Session: the central entity that orchestrates a session lifecycle
//     (Start, Pause, Resume, Finish, Tick) with state validation.
//   - Timing: tracks timestamps for each phase transition and enforces
//     correct sequencing (start →  pause/resume →  finish).
//
// Each type exposes sentinel errors (prefixed with Err) to allow callers to
// handle specific error conditions using errors.Is.
package domain
