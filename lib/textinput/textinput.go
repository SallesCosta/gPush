package lib

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"regexp"
)

type (
	errMsg error
)

var InputValue string

type SecondaryRepoModel struct {
	model     tea.Model
	TextInput textinput.Model
	err       error
	repo      string
}

func sanitizeInput(input string) error {
	matched, err := regexp.Match(`^[a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=%]+$`, []byte(input))
	if err != nil {
		return fmt.Errorf("error while matching regex: %v", err)
	}
	if !matched {
		return fmt.Errorf("input does not match the expected pattern")
	}
	return nil
}

func InputModel(repo string) SecondaryRepoModel {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 500
	ti.Width = 200
	ti.Validate = sanitizeInput

	return SecondaryRepoModel{
		TextInput: ti,
		err:       nil,
		repo:      repo,
	}
}

func (m SecondaryRepoModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SecondaryRepoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			InputValue = m.TextInput.Value()
			return m, tea.Quit
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m SecondaryRepoModel) View() string {
	return fmt.Sprintf(
		"Repsitory url: \n\n%s\n\n%s",
		m.TextInput.View(),
		"(esc to quit)",
	) + "\n"
}
