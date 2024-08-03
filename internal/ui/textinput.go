package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextInput struct {
	textinput textinput.Model
	style     lipgloss.Style
}

func newTextInputField(style lipgloss.Style, prompt string, echoMode textinput.EchoMode) *TextInput {
	ti := textinput.New()

	ti.Prompt = prompt
	ti.EchoMode = echoMode

	return &TextInput{
		textinput: ti,
		style:     style,
	}
}

func (t TextInput) Init() tea.Cmd {
	return nil
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
	t.textinput.Blur()
}

func (t *TextInput) Focus() tea.Cmd {
	return t.textinput.Focus()
}
