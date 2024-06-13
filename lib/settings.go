package lib

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sallescosta/gPush/helpers"
)

type errMsg error

type SettingsModel struct {
	choices       []string
	cursor        int
	selected      map[int]struct{}
	originalRepo  string
	secondaryRepo string
	quitting      bool
	err           error
	showPush      bool
}

func Settings(originalRepo, secondaryRepo string, showPush bool) SettingsModel {
	choices := []string{
		"[i] change origin repo",
		"[s] change secondary repo",
		"[q] quit",
	}

	if showPush {
		choices = append([]string{"[p] git push"}, choices...)
	}

	return SettingsModel{
		choices:       choices,
		selected:      make(map[int]struct{}),
		originalRepo:  originalRepo,
		secondaryRepo: secondaryRepo,
		showPush:      showPush,
	}
}

type Msg int

const (
	up Msg = iota
	down
	enter
)

func (m SettingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	choice := m.choices[m.cursor]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.quitting = true
			return m, tea.Quit
		case tea.KeyUp.String(), tea.KeyShiftTab.String():
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown.String(), tea.KeyTab.String():
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "i":
			helpers.SetRepos("origin")
			return m, tea.Quit
		case "s":
			helpers.SetRepos("secondary")
			return m, tea.Quit

		case "p":
			if !m.showPush {
				return m, nil
			}
			helpers.PushOption()
			return m, tea.Quit

		case tea.KeyEnter.String(), tea.KeySpace.String():
			if m.showPush {
				if !m.showPush {
					return m, nil
				} else {
					if choice == "[p] git push" {
						helpers.PushOption()
						return m, tea.Quit
					}
				}
			}
			if choice == "[i] change origin repo" {
				helpers.SetRepos("origin")
				return m, tea.Quit
			}
			if choice == "[s] change secondary repo" {
				helpers.SetRepos("secondary")
				return m, tea.Quit
			}
			if choice == "[q] quit" {
				return m, tea.Quit
			}
		}
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		return m, cmd
	}
	return m, nil
}

func (m SettingsModel) Init() tea.Cmd {
	return nil
}

func (m SettingsModel) View() string {
	var s string

	s += "Settings menu:\n\n"
	for i, choice := range m.choices {
		cursor := "[  ]"
		if m.cursor == i {
			cursor = "[🔥]"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	if m.quitting {
		s += "\n"
	}

	if m.err != nil {
		s += fmt.Sprintf("\n\nError: %v", m.err)
	}

	return s
}
