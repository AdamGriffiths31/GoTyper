package main

import (
	"time"

	"github.com/AdamGriffiths31/Typing/types"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	bar := progress.NewModel()
	program := tea.NewProgram(types.Model{
		Text:               []rune("the quick brown fox jumps over the lazy dog"),
		Progress:           &bar,
		CompletedTextStyle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00")),
		NextCharStyle:      lipgloss.NewStyle().Underline(true),
		Stopwatch:          stopwatch.NewWithInterval(time.Millisecond),
	})
	program.Start()
}
