package nodes

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Node interface {
	Init() tea.Cmd
	Update(tea.Msg) (Node, tea.Cmd)
	View() string

	GetNext() Node
	Handle(*Node) tea.Cmd
}

type Sizes struct {
	width  int
	height int
}

type BaseNode struct {
	cursor int
	fields []Field
	next   Node
	sizes  Sizes
}

func newBaseNode(width, height int, choices ...Field) BaseNode {
	return BaseNode{
		cursor: 0,
		fields: choices,
		sizes: Sizes{
			width:  width,
			height: height,
		},
	}
}

func (n BaseNode) GetNext() Node {
	return n.next
}

func (n *BaseNode) Handle(node *Node) tea.Cmd {
	cmd := n.fields[n.cursor].Handle(node)
	return cmd
}

func (n BaseNode) Init() tea.Cmd {
	return nil
}

func (n *BaseNode) moveCursor(step int) tea.Cmd {
	mod := len(n.fields)

	switch field := n.fields[n.cursor].(type) {
	case *Block:
		ok, cmd := field.moveCursor(step)

		if !ok {
			break
		}

		n.fields[n.cursor] = field

		return cmd
	}

	n.fields[n.cursor].Blur()
	n.cursor = ((n.cursor+step)%mod + mod) % mod

	switch field := n.fields[n.cursor].(type) {
	case *Block:
		var cmd tea.Cmd

		if step > 0 {
			field.cursor = 0
			cmd = field.fields[0].Focus()
		} else {
			field.cursor = len(field.fields) - 1
			cmd = field.fields[len(field.fields)-1].Focus()
		}

		n.fields[n.cursor] = field

		return cmd
	}

	return n.fields[n.cursor].Focus()
}

// An update function for non-base nodes.
// It delegates the update for base node.
func updateNode(baseNode *BaseNode, msg tea.Msg) tea.Cmd {
	node, cmd := baseNode.Update(msg)

	tempNode, ok := node.(*BaseNode)
	if !ok {
		log.Fatal("Couldn't convert tempNode to baseNode while updating non-base node.")
	}

	*baseNode = *tempNode

	return cmd
}

func (n *BaseNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		n.sizes.width = msg.Width
		n.sizes.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			cmd = n.moveCursor(-1)
		case "down", "tab":
			cmd = n.moveCursor(1)
		default:
			n.fields[n.cursor], cmd = n.fields[n.cursor].Update(msg)
		}
	}

	return n, cmd
}

func (n BaseNode) render() string {
	result := make([]string, 0)

	for _, field := range n.fields {
		result = append(result, field.View())
	}

	return lipgloss.JoinVertical(lipgloss.Top, result...)
}

func (n BaseNode) View() string {
	return n.render()
}
