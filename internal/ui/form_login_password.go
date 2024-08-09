package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type LoginPasswordForm struct {
	BaseNode
}

func NewLoginPasswordForm(width, height int, command *cobra.Command) *LoginPasswordForm {
	widthForNode := 40

	fields := []Field{
		newBlock(
			defaultBlockStyle.Border(lipgloss.RoundedBorder()).Padding(0, 1),
			newTextInputField(
				" Service: ",
				textinput.EchoNormal,
				defaultTextInputStyle.Width(widthForNode),
			),
			newTextInputField(
				"   Login: ",
				textinput.EchoNormal,
				defaultTextInputStyle.Width(widthForNode),
			),
			newTextInputField(
				"Password: ",
				textinput.EchoPassword,
				defaultTextInputStyle.Width(widthForNode),
			),
		),
		newChoice(
			"SAVE",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*LoginPasswordForm)
				values := currentNode.fields[0].Value().([]any)

				service, ok := values[0].(string)
				if service == "" || !ok {
					cmd := formMessage(
						model,
						"Bad service. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				login, ok := values[1].(string)
				if login == "" || !ok {
					cmd := formMessage(
						model,
						"Bad login. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				password, ok := values[2].(string)
				if password == "" || !ok {
					cmd := formMessage(
						model,
						"Bad password. It should be non-empty and string.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				}

				model.userContext.service = service
				model.userContext.data = fmt.Sprintf("login: %s, password: %s", login, password)
				MapUserContextToDatabaseVariables(model.userContext)

				command.PreRun(command, nil)
				command.Run(command, nil)
				model.SetNode(NewControlPanelNode(currentNode.sizes.width, currentNode.sizes.height))

				cmd := formMessage(
					model,
					"Successfully saved",
					defaultMessageStyle,
					3*time.Second,
				)

				return true, cmd
			},
			defaultUnfocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginTop(1),
			defaultFocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginTop(1),
		),
	}

	return &LoginPasswordForm{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *LoginPasswordForm) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n LoginPasswordForm) View() string {
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
