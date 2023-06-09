package types

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/AdamGriffiths31/Typing/database"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	logger *log.Logger
	db     *database.DB

	Progress *progress.Model
	Percent  float64

	Text          []rune
	CompletedText []rune
	Score         int
	Errors        int

	CompletedTextStyle lipgloss.Style
	NextCharStyle      lipgloss.Style

	Stopwatch stopwatch.Model
	WPM       int
}

func NewModel(text string, logger *log.Logger, db *database.DB) *Model {
	bar := progress.NewModel()
	return &Model{
		logger:             logger,
		db:                 db,
		Text:               []rune(text),
		Progress:           &bar,
		CompletedTextStyle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00")),
		NextCharStyle:      lipgloss.NewStyle().Underline(true),
		Stopwatch:          stopwatch.NewWithInterval(time.Millisecond),
	}
}

func (m Model) Init() tea.Cmd {
	m.logger.Println("Text:", string(m.Text))
	return m.Stopwatch.Init()
}

// Update handles the bubbletea model updating
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
		if msg.Type != tea.KeyRunes && msg.Type != tea.KeySpace {
			//Not an expected input
			return m, nil
		}
		m.validate(msg.Runes[0])
	}

	if m.Percent >= 1 {
		m.db.WriteRecord(int(m.Stopwatch.Elapsed().Milliseconds()))
		return m, tea.Quit
	}
	var cmd tea.Cmd
	m.Stopwatch, cmd = m.Stopwatch.Update(msg)
	return m, cmd
}

// View returns the UI
func (m Model) View() string {
	var sb strings.Builder
	//Render the text already completed
	sb.WriteString(m.CompletedTextStyle.Render(string(m.CompletedText)))
	// Render the active char
	if len(m.Text) != len(m.CompletedText) {
		sb.WriteString(m.NextCharStyle.Render(string(m.Text[m.Score])))
	}
	//Render the text yet to be completed
	if m.Score < len(m.Text)-1 {
		sb.WriteString(string(m.Text[m.Score+1:]))
	}
	return fmt.Sprintf("%s\n%s\n%s\nWPM: %v\tErrors: %d\n\n", m.Progress.ViewAs(m.Percent), sb.String(), m.Stopwatch.View(), m.WPM, m.Errors)
}

// validate checks the input value is the next correct value
func (m *Model) validate(input rune) error {
	expected := m.Text[m.Score]
	if input != expected {
		m.Errors++
		return fmt.Errorf("expected '%c', but got '%c'", expected, input)
	}

	m.updateScore()
	m.updateWPM()
	m.CompletedText = append(m.CompletedText, input)

	return nil
}

// updateScore increments the score and updates the percentage completed
func (m *Model) updateScore() {
	m.Score++
	m.Percent = float64(m.Score) / float64(len(m.Text))
	if m.Score == 1 {
		m.Stopwatch.Toggle()
	}
}

// updateWPM updates the words per minute
func (m *Model) updateWPM() {
	words := strings.Count(string(m.Text), " ") + 1
	if int(m.Stopwatch.Elapsed().Seconds()) > 0 {
		m.WPM = (words / int(m.Stopwatch.Elapsed().Seconds())) * 60
	}
}
