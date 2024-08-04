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
			newTextInputField("   Login: ", textinput.EchoNormal, defaultTextInputStyle.MaxWidth(40)),
			newTextInputField("Password: ", textinput.EchoPassword, defaultTextInputStyle.MaxWidth(40)),
		),
		newChoice(
			"ENTER",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*LoginNode)
				model.node = NewSaltNode(currentNode.sizes.width, currentNode.sizes.height)

				values := currentNode.fields[0].Value().([]any)
				model.userContext.login = values[0].(string)
				model.userContext.password = values[1].(string)

				return true, nil
			},
			defaultUnfocusedStyle.Width(30).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(30).AlignHorizontal(lipgloss.Center),
		),
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
