package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Text struct {
	value string
	style lipgloss.Style
}

func newText(value string, style lipgloss.Style) *Text {
	return &Text{
		value: value,
		style: style,
	}
}

func (t Text) Init() tea.Cmd {
	return nil
}

func (t *Text) Update(msg tea.Msg) (Field, tea.Cmd) {
	return t, nil
}

func (t Text) View() string {
	return t.style.Render(t.value)
}

func (t *Text) Blur() {
}

func (t *Text) Focus() tea.Cmd {
	return nil
}

func (t *Text) Handle(model *model) (bool, tea.Cmd) {
	return false, nil
}

func (t *Text) InheritStyle(style lipgloss.Style) {
	t.style = t.style.Inherit(style)
}

func (t *Text) Value() any {
	return t.value
}
