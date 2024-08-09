package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hikkiyomi/passman/cmd/actions"
)

type ControlPanelNode struct {
	BaseNode
}

func NewControlPanelNode(width, height int) *ControlPanelNode {
	widthForNode := 40

	fields := []Field{
		newText("Choose command", lipgloss.NewStyle().MarginBottom(1).Width(widthForNode).AlignHorizontal(lipgloss.Center)),
		newChoice(
			"Save",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*ControlPanelNode)
				model.SetNode(NewSaveNode(currentNode.sizes.width, currentNode.sizes.height, actions.SaveCmd))
				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
	}

	return &ControlPanelNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *ControlPanelNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n *ControlPanelNode) View() string {
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
