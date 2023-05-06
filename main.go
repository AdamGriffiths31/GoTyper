package main

import (
	"flag"
	"time"

	randomtext "github.com/AdamGriffiths31/Typing/randomText"
	"github.com/AdamGriffiths31/Typing/types"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	mode := flag.String("mode", "", "Sets the mode of the program")
	flag.Parse()

	var text string

	switch *mode {
	case "random":
		text = randomtext.GenerateText()
	default:
		text = "the quick brown fox jumps over the lazy dog"
	}

	execute(text)
}

func execute(text string) {
	bar := progress.NewModel()
	program := tea.NewProgram(types.Model{
		Text:               []rune(text),
		Progress:           &bar,
		CompletedTextStyle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00")),
		NextCharStyle:      lipgloss.NewStyle().Underline(true),
		Stopwatch:          stopwatch.NewWithInterval(time.Millisecond),
	})
	program.Start()
}
