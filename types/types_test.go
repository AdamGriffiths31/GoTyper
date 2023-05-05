package types

import "testing"

func TestValidateCorrectChar(t *testing.T) {
	m := Model{
		Text:          []rune("hello"),
		CompletedText: []rune("he"),
		Score:         2,
	}

	err := m.validate('l')
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if string(m.CompletedText) != "hel" {
		t.Errorf("Expected hel but got %s", string(m.CompletedText))
	}

	if m.Score != 3 {
		t.Errorf("Expected 3 but got %d", m.Score)
	}

	if m.Percent != 0.6 {
		t.Errorf("Expected 0.6 but got %v", m.Percent)
	}

}

func TestValidateIncorrectChar(t *testing.T) {
	m := Model{
		Text:          []rune("hello"),
		CompletedText: []rune("he"),
		Score:         2,
	}

	err := m.validate('t')
	if err == nil {
		t.Error("Expected error")
	}

	if string(m.CompletedText) != "he" {
		t.Errorf("Expected he but got %s", string(m.CompletedText))
	}

	if m.Score != 2 {
		t.Errorf("Expected 2 but got %d", m.Score)
	}

	if m.Percent != 0 {
		t.Errorf("Expected 0 but got %v", m.Percent)
	}

}
