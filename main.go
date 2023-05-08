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
	mode := flag.String("mode", "", "Sets the mode of the program. Options include: random")
	wordCount := flag.Int("words", 10, "Sets the word count for random mode")
	flag.Parse()
	flag.PrintDefaults()
	var text string
	switch *mode {
	case "random":
		text = randomtext.GenerateText(*wordCount)
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
