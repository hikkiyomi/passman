package nodes

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoginNode struct {
	BaseNode
}

func NewLoginNode(width, height int) LoginNode {
	fields := []Field{
		newBlock(
			lipgloss.NewStyle().Margin(1),
			newTextInputField("   Login: ", textinput.EchoNormal),
			newTextInputField("Password: ", textinput.EchoPassword),
		),
		newChoice("ENTER", func(node *BaseNode) tea.Cmd {
			return nil
		}),
	}

	return LoginNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n LoginNode) Init() tea.Cmd {
	return n.fields[0].Focus()
}

func (n LoginNode) Handle() (Node, tea.Cmd) {
	cmd := handleNode(&n.BaseNode)
	return n, cmd
}

func (n LoginNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n LoginNode) View() string {
	return lipgloss.Place(
		n.sizes.width,
		n.sizes.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Top,
			"LOGIN PAGE",
			"",
			n.BaseNode.View(),
		),
	)
}
