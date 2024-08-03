package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoginNode struct {
	BaseNode
}

func NewLoginNode(width, height int) *LoginNode {
	fields := []Field{
		newBlock(
			defaultBlockStyle.Border(lipgloss.RoundedBorder()).PaddingLeft(1).PaddingRight(1),
			newTextInputField(defaultTextInputStyle.MaxWidth(40), "   Login: ", textinput.EchoNormal),
			newTextInputField(defaultTextInputStyle.MaxWidth(40), "Password: ", textinput.EchoPassword),
		),
		newChoice(defaultUnfocusedStyle, defaultFocusedStyle, "ENTER", func(model *model) (bool, tea.Cmd) {
			// TODO: GO TO SALT NODE
			return false, nil
		}),
	}

	return &LoginNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n LoginNode) Init() tea.Cmd {
	return n.fields[0].Focus()
}

func (n *LoginNode) Update(msg tea.Msg) (Node, tea.Cmd) {
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
			n.BaseNode.View(),
		),
	)
}
