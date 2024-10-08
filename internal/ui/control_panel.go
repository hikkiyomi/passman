package ui

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hikkiyomi/passman/cmd/actions"
	"github.com/hikkiyomi/passman/internal/common"
)

type ControlPanelNode struct {
	BaseNode
}

func tryToOpenDatabase(model *model) (bool, tea.Cmd) {
	MapUserContextToDatabaseVariables(model.userContext)

	if _, err := os.Stat(common.Path); err != nil {
		msgCmd := formMessage(
			model,
			"There is no database at given path yet.",
			defaultErrorStyle,
			3*time.Second,
		)

		return false, msgCmd
	}

	actions.GetCmd.PreRun(actions.GetCmd, nil)

	return true, nil
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
		newChoice(
			"Get",
			func(model *model) (bool, tea.Cmd) {
				if success, cmd := tryToOpenDatabase(model); !success {
					return success, cmd
				}

				currentNode := model.node.(*ControlPanelNode)

				model.SetNode(
					NewGetNode(
						currentNode.sizes.width,
						currentNode.sizes.height,
						defaultListStyle,
					),
				)

				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
		newChoice(
			"Update",
			func(model *model) (bool, tea.Cmd) {
				if success, cmd := tryToOpenDatabase(model); !success {
					return success, cmd
				}

				currentNode := model.node.(*ControlPanelNode)

				model.SetNode(
					NewUpdateNode(
						currentNode.sizes.width,
						currentNode.sizes.height,
						defaultListStyle,
						actions.UpdateCmd,
					),
				)

				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
		newChoice(
			"Remove",
			func(model *model) (bool, tea.Cmd) {
				if success, cmd := tryToOpenDatabase(model); !success {
					return success, cmd
				}

				currentNode := model.node.(*ControlPanelNode)

				model.SetNode(
					NewRemoveNode(
						currentNode.sizes.width,
						currentNode.sizes.height,
						defaultListStyle,
						actions.UpdateCmd,
					),
				)

				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
		newChoice(
			"Import",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*ControlPanelNode)

				model.SetNode(
					NewFilePicker(
						currentNode.sizes.width,
						currentNode.sizes.height,
						false,
						newTextInputField(
							"Browser type (empty/chrome/firefox): ",
							textinput.EchoNormal,
							defaultTextInputStyle,
						),
						handlerForImporting,
					),
				)

				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-4).AlignHorizontal(lipgloss.Center),
		),
		newChoice(
			"Export",
			func(model *model) (bool, tea.Cmd) {
				if success, cmd := tryToOpenDatabase(model); !success {
					return success, cmd
				}

				currentNode := model.node.(*ControlPanelNode)

				model.SetNode(
					NewFilePicker(
						currentNode.sizes.width,
						currentNode.sizes.height,
						true,
						newTextInputField(
							"Export into: ",
							textinput.EchoNormal,
							defaultTextInputStyle,
						),
						handlerForExporting,
					),
				)

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

func (n *ControlPanelNode) Clear() {
	*n = *NewControlPanelNode(n.sizes.width, n.sizes.height)
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
