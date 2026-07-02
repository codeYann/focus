package tui

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/codeYann/focus/internal/domain"
)

// Model is the Bubble Tea model for the focus timer UI.
// It holds the session state, available commands, and any errors
// that occur during user interaction.
type Model struct {
	session  domain.Session
	commands []Command
	error    error
}

// NewModel creates a new Model with the given session and default commands.
func NewModel(session domain.Session) Model {
	return Model{
		session:  session,
		commands: Commands[:],
	}
}

// Init initializes the model. It returns nil as no startup command is needed.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model state accordingly.
// Supported key presses:
//   - s: start the session
//   - p: pause the session
//   - r: resume the session
//   - q, esc, ctrl+c: quit the program
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "s":
			if err := m.session.Start(time.Now()); err != nil {
				m.error = err
			}
			return m, nil

		case "p":
			if err := m.session.Pause(time.Now()); err != nil {
				m.error = err
			}
			return m, nil

		case "r":
			if err := m.session.Resume(); err != nil {
				m.error = err
			}
			return m, nil
		}
	}
	return m, nil
}

// View renders the current session state and keybindings to the terminal.
func (m Model) View() tea.View {
	if m.error != nil {
		return tea.NewView(fmt.Sprintf("\n Erro ao apresentar uma nova sessão: %v\n\n", m.error))
	}

	var s strings.Builder
	s.WriteString("⏳ Focus \n\n")

	remaining := m.session.Remaining

	if remaining != 0 {
		s.WriteString(remaining.String())
		s.WriteString("\n\n")
	}

	for _, command := range m.commands {
		s.WriteString("[")
		s.WriteString(command.Key)
		s.WriteString("]")
		s.WriteString(" ")
		s.WriteString(command.Label)
		s.WriteString("  ")
	}

	view := tea.NewView(s.String())

	view.AltScreen = true

	return view
}
