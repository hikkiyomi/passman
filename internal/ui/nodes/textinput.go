package nodes

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInput struct {
	textinput textinput.Model
}

func newTextInputField(prompt string, echoMode textinput.EchoMode) *TextInput {
	ti := textinput.New()

	ti.Prompt = prompt
	ti.EchoMode = echoMode

	return &TextInput{
		textinput: ti,
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
	return textInputStyle.Render(t.textinput.View())
}

func (t *TextInput) Handle(node *BaseNode) tea.Cmd {
	return nil
}

func (t *TextInput) Blur() {
	t.textinput.Blur()
}

func (t *TextInput) Focus() tea.Cmd {
	return t.textinput.Focus()
}