package main

import (
	"github.com/AdamGriffiths31/Typing/types"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bar := progress.NewModel()
	program := tea.NewProgram(types.Model{
		Text:     []rune("type this text"),
		Progress: &bar,
	})
	program.Start()
}
