package types

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Progress *progress.Model
	Percent  float64

	Text  []rune
	Score int
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles the bubbletea model updating
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.validate(msg.Runes[0])
	}
	return m, nil
}

// View returns the UI
func (m Model) View() string {
	return fmt.Sprintf("%s\n%s\nScore: %d", m.Progress.ViewAs(m.Percent), string(m.Text), m.Score)
}

// validate checks the input value is the next correct value
func (m *Model) validate(input rune) error {
	if m.Text[m.Score] == input {
		m.updateScore()
	}
	return nil
}

func (m *Model) updateScore() {
	m.Score++
	m.Percent = float64(m.Score / len(m.Text))
}
