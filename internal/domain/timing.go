package domain

import (
	"errors"
	"time"
)

// Errors returned by Timing methods.
var (
	ErrTimingAlreadyStarted  = errors.New("Erro ao iniciar o timing. Não é possível iniciar o timing duas vezes")
	ErrTimingAlreadyPaused   = errors.New("Erro ao pausar o timing. Timing já encontra-se pausado.")
	ErrTimingNotPaused       = errors.New("Erro ao retomar o timing! Timing não está pausado.")
	ErrTimingAlreadyFinished = errors.New("Erro ao finalizar o timing. Timing já encontra-se finalizado.")
)

// Timing tracks timestamps for session state transitions.
type Timing struct {
	startedAt  *time.Time
	finishedAt *time.Time
	pausedAt   *time.Time
}

func (t Timing) hasStarted() bool {
	return t.startedAt != nil
}

func (t Timing) hasFinished() bool {
	return t.finishedAt != nil
}

func (t Timing) hasPaused() bool {
	return t.pausedAt != nil
}

// Start records the start time. Returns ErrTimingAlreadyStarted if already started.
func (t *Timing) Start(now time.Time) error {
	if t.hasStarted() {
		return ErrTimingAlreadyStarted
	}

	t.startedAt = &now

	return nil
}

// Pause records the pause time. Returns ErrTimingAlreadyPaused if already paused.
func (t *Timing) Pause(now time.Time) error {
	if t.hasPaused() {
		return ErrTimingAlreadyPaused
	}

	t.pausedAt = &now

	return nil
}

// Resume clears the pause time. Returns ErrTimingNotPaused if not paused.
func (t *Timing) Resume() error {
	if !t.hasPaused() {
		return ErrTimingNotPaused
	}

	t.pausedAt = nil
	return nil
}

// Finish records the finish time. Returns ErrTimingAlreadyFinished if already finished.
func (t *Timing) Finish(now time.Time) error {
	if t.hasFinished() {
		return ErrTimingAlreadyFinished
	}

	t.finishedAt = &now

	return nil
}
