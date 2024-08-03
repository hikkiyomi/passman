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
		newChoice("Login", func(node *Node) tea.Cmd {
			tempNode := *node
			currentNode := tempNode.(*WelcomeNode)
			*node = NewLoginNode(currentNode.sizes.width, currentNode.sizes.height)
			return nil
		}),
		newChoice("Exit", func(node *Node) tea.Cmd {
			return tea.Quit
		}),
	}

	return &WelcomeNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n WelcomeNode) Init() tea.Cmd {
	return n.fields[0].Focus()
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
			"WELCOME TO PASSMAN",
			"",
			n.BaseNode.View(),
		),
	)
}
