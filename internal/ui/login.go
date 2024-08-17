package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoginNode struct {
	BaseNode
}

func NewLoginNode(width, height int) *LoginNode {
	widthForNode := 30

	fields := []Field{
		newBlock(
			defaultBlockStyle.Border(lipgloss.RoundedBorder()).Padding(0, 1),
			newTextInputField("   Login: ", textinput.EchoNormal, defaultTextInputStyle.Width(widthForNode)),
			newTextInputField("Password: ", textinput.EchoPassword, defaultTextInputStyle.Width(widthForNode)),
		),
		newChoice(
			"ENTER",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*LoginNode)
				values := currentNode.fields[0].Value().([]any)

				login, ok := values[0].(string)
				if login == "" || !ok {
					cmd := formMessage(
						model,
						"Bad login. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				password, ok := values[1].(string)
				if password == "" || !ok {
					cmd := formMessage(
						model,
						"Bad password. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				model.SetNode(NewSaltNode(currentNode.sizes.width, currentNode.sizes.height))
				model.userContext.login = login
				model.userContext.password = password

				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center),
		),
	}

	return &LoginNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *LoginNode) Clear() {
	*n = *NewLoginNode(n.sizes.width, n.sizes.height)
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
