package nodes

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WelcomeNode struct {
	BaseNode
}

func NewWelcomeNode() WelcomeNode {
	choices := []Choice{
		{
			name: "Login",
			handler: func() tea.Cmd {
				return nil
			},
		},
		{
			name: "Exit",
			handler: func() tea.Cmd {
				return tea.Quit
			},
		},
	}

	return WelcomeNode{
		BaseNode: newBaseNode(choices),
	}
}

func (n WelcomeNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	var cmd tea.Cmd
	node, cmd := n.BaseNode.Update(msg)

	tempNode, ok := node.(BaseNode)
	if !ok {
		log.Fatalf("Couldn't convert node to BaseNode")
	}

	n.BaseNode = tempNode

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
