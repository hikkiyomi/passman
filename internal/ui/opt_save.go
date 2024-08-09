package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type SaveNode struct {
	BaseNode
}

func NewSaveNode(width, height int, command *cobra.Command) *SaveNode {
	widthForNode := 40

	fields := []Field{
		newText("Choose a form", defaultHeaderTextStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginBottom(1)),
		newChoice(
			"Login & Password",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*SaveNode)
				model.SetNode(NewLoginPasswordForm(currentNode.sizes.width, currentNode.sizes.height, command))
				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
		newChoice(
			"Arbitrary data",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*SaveNode)
				model.SetNode(NewArbitraryDataForm(currentNode.sizes.width, currentNode.sizes.height, command))
				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
	}

	return &SaveNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *SaveNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n SaveNode) View() string {
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
