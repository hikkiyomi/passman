package nodes

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Node interface {
	Init() tea.Cmd
	Update(tea.Msg) (Node, tea.Cmd)
	View() string

	GetNext() Node
	GetChoice() string
	HandleChoice() tea.Cmd
}

type Sizes struct {
	width  int
	height int
}

type BaseNode struct {
	cursor  int
	choices []Choice
	next    Node
	sizes   Sizes
}

func newBaseNode(choices []Choice) BaseNode {
	return BaseNode{
		cursor:  0,
		choices: choices,
	}
}

func (n BaseNode) GetNext() Node {
	return n.next
}

func (n BaseNode) GetChoice() string {
	return n.choices[n.cursor].name
}

func (n BaseNode) HandleChoice() tea.Cmd {
	return n.choices[n.cursor].Handle()
}

func (n BaseNode) Init() tea.Cmd {
	return nil
}

func (n *BaseNode) moveCursor(step int) {
	mod := len(n.choices)
	n.cursor = ((n.cursor+step)%mod + mod) % mod
}

func (n BaseNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		n.sizes.width = msg.Width
		n.sizes.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			n.moveCursor(-1)
		case "down":
			n.moveCursor(1)
		}
	}

	return n, nil
}

func (n BaseNode) renderChoices() string {
	result := make([]string, 0)

	for i, choice := range n.choices {
		prefix := "   "

		if i == n.cursor {
			prefix = "-> "
		}

		result = append(result, fmt.Sprintf("%v%v", prefix, choice.name))
	}

	return lipgloss.JoinVertical(lipgloss.Top, result...)
}

func (n BaseNode) View() string {
	return n.renderChoices()
}
