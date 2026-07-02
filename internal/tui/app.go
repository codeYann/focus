package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/google/uuid"

	"github.com/codeYann/focus/internal/config"
	"github.com/codeYann/focus/internal/domain"
)

// Run creates a new focus session and starts the Bubble Tea program.
// It returns an error if the program fails to run.
func Run(cfg config.Config) error {
	ID := uuid.New()
	session := domain.New(ID.String(), cfg.Domain())

	model := NewModel(session)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		return err
	}

	return nil
}
