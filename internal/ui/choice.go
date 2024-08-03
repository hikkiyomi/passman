package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Choice struct {
	name           string
	handler        func(*model) (bool, tea.Cmd)
	isFocused      bool
	unfocusedStyle lipgloss.Style
	focusedStyle   lipgloss.Style
}

func newChoice(
	unfocusedStyle lipgloss.Style,
	focusedStyle lipgloss.Style,
	name string,
	handler func(*model) (bool, tea.Cmd),
) *Choice {
	return &Choice{
		name:           name,
		handler:        handler,
		unfocusedStyle: unfocusedStyle,
		focusedStyle:   focusedStyle,
	}
}

func (c Choice) Init() tea.Cmd {
	return nil
}

func (c *Choice) Update(msg tea.Msg) (Field, tea.Cmd) {
	return c, nil
}

func (c Choice) View() string {
	if c.isFocused {
		return c.focusedStyle.Render("> " + c.name)
	}

	return c.unfocusedStyle.Render(c.name)
}

func (c Choice) Handle(model *model) (bool, tea.Cmd) {
	return c.handler(model)
}

func (c *Choice) Blur() {
	c.isFocused = false
}

func (c *Choice) Focus() tea.Cmd {
	c.isFocused = true
	return nil
}
