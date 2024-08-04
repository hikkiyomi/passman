package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WelcomeNode struct {
	BaseNode
}

func NewWelcomeNode(width, height int) *WelcomeNode {
	fields := []Field{
		newText("WELCOME TO PASSMAN", lipgloss.NewStyle().Margin(0, 1)),
		newChoice(
			"Login",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*WelcomeNode)
				model.node = NewLoginNode(currentNode.sizes.width, currentNode.sizes.height)
				return true, nil
			},
			defaultUnfocusedStyle,
			defaultFocusedStyle,
		),
		newChoice(
			"Exit",
			func(model *model) (bool, tea.Cmd) {
				return false, tea.Quit
			},
			defaultUnfocusedStyle,
			defaultFocusedStyle,
		),
	}

	return &WelcomeNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *WelcomeNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n WelcomeNode) View() string {
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
