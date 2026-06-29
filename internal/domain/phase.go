package domain

// Phase represents the current state of a session.
type Phase uint8

const (
	Idle      Phase = iota // Initial state before a session starts.
	Focus                  // Session is in a focus period.
	Break                  // Session is in a break period.
	Completed              // Session has finished.
)
