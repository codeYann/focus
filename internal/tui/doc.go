// Package tui provides a terminal user interface for focus sessions using
// the Bubble Tea framework.
//
// It wires together the domain model with a keyboard-driven terminal UI,
// handling user input and rendering the session state in real time.
//
// The main types are:
//   - Model: the Bubble Tea model implementing Init, Update, and View.
//   - Command: maps a key press to a human-readable action label.
//   - Run: the top-level entry point that creates a Model and starts
//     the Bubble Tea event loop.
package tui
