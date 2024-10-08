package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Node interface {
	Init() tea.Cmd
	Update(tea.Msg) (Node, tea.Cmd)
	View() string

	Handle(*model) (bool, tea.Cmd)
	Clear()
}

type baseNodeKeymap struct {
	Up   key.Binding
	Down key.Binding
}

var defaultBaseNodeKeymap = baseNodeKeymap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "tab"),
		key.WithHelp("↓/tab", "move down"),
	),
}

func (k baseNodeKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down}
}

func (k baseNodeKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type BaseNode struct {
	cursor int
	fields []Field
	sizes  Sizes
	help   help.Model
	keymap baseNodeKeymap
}

func newBaseNode(width, height int, choices ...Field) BaseNode {
	return BaseNode{
		cursor: 0,
		fields: choices,
		sizes: Sizes{
			width:  width,
			height: height,
		},
		help:   help.New(),
		keymap: defaultBaseNodeKeymap,
	}
}

func (n *BaseNode) Clear() {
	*n = newBaseNode(n.sizes.width, n.sizes.height, n.fields...)
}

func (n *BaseNode) Handle(model *model) (bool, tea.Cmd) {
	return n.fields[n.cursor].Handle(model)
}

func (n *BaseNode) Init() tea.Cmd {
	firstNonTextField := func() int {
		result := -1

		for i, field := range n.fields {
			_, ok := field.(*Text)
			if !ok {
				result = i
				break
			}
		}

		return result
	}()

	if firstNonTextField != -1 {
		n.cursor = firstNonTextField
		return n.fields[firstNonTextField].Init()
	}

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
	case *Text:
		additionalStep := 1

		if step < 0 {
			additionalStep = -1
		}

		return n.moveCursor(additionalStep)
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
	return n.render() + "\n\n" + n.help.View(n.keymap)
}
