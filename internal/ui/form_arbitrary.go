package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type ArbitraryDataForm struct {
	BaseNode
	cmd *cobra.Command
}

func NewArbitraryDataForm(width, height int, command *cobra.Command) *ArbitraryDataForm {
	widthForNode := 40

	fields := []Field{
		newBlock(
			defaultBlockStyle.Border(lipgloss.RoundedBorder()).Padding(0, 1),
			newTextInputField(
				"Service: ",
				textinput.EchoNormal,
				defaultTextInputStyle.Width(widthForNode),
			),
			newTextInputField(
				"   Data: ",
				textinput.EchoPassword,
				defaultTextInputStyle.Width(widthForNode),
			),
		),
		newChoice(
			"SAVE",
			func(m *model) (bool, tea.Cmd) {
				currentNode := m.node.(*ArbitraryDataForm)
				values := currentNode.fields[0].Value().([]any)

				service, ok := values[0].(string)
				if service == "" || !ok {
					cmd := formMessage(
						m,
						"Bad service. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				data, ok := values[1].(string)
				if data == "" || !ok {
					cmd := formMessage(
						m,
						"Bad data. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				m.userContext.service = service
				m.userContext.data = data
				MapUserContextToDatabaseVariables(m.userContext)

				command.PreRun(command, nil)
				command.Run(command, nil)

				rollbackCmd := m.rollbackUntil(
					NewControlPanelNode(currentNode.sizes.width, currentNode.sizes.height),
					func(model *model) bool {
						last := model.nodeHistory[len(model.nodeHistory)-1]
						_, ok := last.(*ControlPanelNode)

						return ok
					},
				)

				cmd := formMessage(
					m,
					"Successfully saved",
					defaultMessageStyle,
					3*time.Second,
				)

				return true, tea.Batch(rollbackCmd, cmd)
			},
			defaultUnfocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginTop(1),
			defaultFocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginTop(1),
		),
	}

	return &ArbitraryDataForm{
		BaseNode: newBaseNode(width, height, fields...),
		cmd:      command,
	}
}

func (n *ArbitraryDataForm) Clear() {
	*n = *NewArbitraryDataForm(n.sizes.width, n.sizes.height, n.cmd)
}

func (n *ArbitraryDataForm) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n ArbitraryDataForm) View() string {
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
