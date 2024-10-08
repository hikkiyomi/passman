package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextInput struct {
	textinput textinput.Model
	style     lipgloss.Style
	isFocused bool
}

func newTextInputField(
	prompt string,
	echoMode textinput.EchoMode,
	style lipgloss.Style,
) *TextInput {
	ti := textinput.New()

	ti.Prompt = prompt
	ti.EchoMode = echoMode

	return &TextInput{
		textinput: ti,
		style:     style.Inline(true),
	}
}

func (t *TextInput) Init() tea.Cmd {
	return t.Focus()
}

func (t *TextInput) Update(msg tea.Msg) (Field, tea.Cmd) {
	var cmd tea.Cmd
	t.textinput, cmd = t.textinput.Update(msg)
	return t, cmd
}

func (t TextInput) View() string {
	return t.style.Render(t.textinput.View())
}

func (t *TextInput) Handle(model *model) (bool, tea.Cmd) {
	return false, nil
}

func (t *TextInput) Blur() {
	t.isFocused = false
	t.textinput.Blur()
}

func (t *TextInput) Focus() tea.Cmd {
	t.isFocused = true
	return t.textinput.Focus()
}

func (t *TextInput) Toggle() tea.Cmd {
	var cmd tea.Cmd

	if t.isFocused {
		t.Blur()
	} else {
		cmd = t.Focus()
	}

	return cmd
}

func (t *TextInput) InheritStyle(style lipgloss.Style) {
	t.style = t.style.Inherit(style)
}

func (t *TextInput) Value() any {
	return t.textinput.Value()
}
