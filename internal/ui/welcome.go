package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WelcomeNode struct {
	BaseNode
}

func NewWelcomeNode(width, height int) *WelcomeNode {
	widthForNode := 40

	welcomeText := " _ __   __ _ ___ ___ _ __ ___   __ _ _ __  \n" +
		"| '_ \\ / _` / __/ __| '_ ` _ \\ / _` | '_ \\ \n" +
		"| |_) | (_| \\__ \\__ \\ | | | | | (_| | | | |\n" +
		"| .__/ \\__,_|___/___/_| |_| |_|\\__,_|_| |_|\n" +
		"|_|                                        "

	fields := []Field{
		newText(welcomeText, defaultHeaderTextStyle),
		newChoice(
			"Login",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*WelcomeNode)
				model.SetNode(NewLoginNode(currentNode.sizes.width, currentNode.sizes.height))
				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center),
		),
		newChoice(
			"Exit",
			func(model *model) (bool, tea.Cmd) {
				return false, tea.Quit
			},
			defaultUnfocusedStyle.Width(widthForNode-1).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-1).AlignHorizontal(lipgloss.Center),
		),
	}

	return &WelcomeNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *WelcomeNode) Clear() {
	*n = *NewWelcomeNode(n.sizes.width, n.sizes.height)
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
