package tui

// Command represents an action available to the user, with its keybinding and display label.
type Command struct {
	// Key is the keyboard shortcut that triggers this command.
	Key string
	// Label is the human-readable name shown in the UI.
	Label string
}

// Commands is the fixed set of keybindings available during a session.
var Commands = [...]Command{
	{"s", "Iniciar"},
	{"p", "Pausar"},
	{"r", "Retomar"},
	{"q", "Sair"},
}
