package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DatabaseSelectionNode struct {
	BaseNode
}

func NewDatabaseSelectionNode(width, height int) *DatabaseSelectionNode {
	fields := []Field{
		newChoice(
			"Create new database",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*DatabaseSelectionNode)
				model.SetNode(NewFilePicker(currentNode.sizes.width, currentNode.sizes.height, true))
				return true, nil
			},
			defaultUnfocusedStyle,
			defaultFocusedStyle,
		),
		newChoice(
			"Open existing",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*DatabaseSelectionNode)
				model.SetNode(NewFilePicker(currentNode.sizes.width, currentNode.sizes.height, false))
				return true, nil
			},
			defaultUnfocusedStyle,
			defaultFocusedStyle,
		),
	}

	return &DatabaseSelectionNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *DatabaseSelectionNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n DatabaseSelectionNode) View() string {
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
