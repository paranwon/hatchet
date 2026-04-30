package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
)

// main is the entry point for the Hatchet TUI application.
// It initializes the root model and starts the Bubble Tea program.
func main() {
	p := tea.NewProgram(
		NewRootModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running hatchet TUI: %v\n", err)
		os.Exit(1)
	}
}
