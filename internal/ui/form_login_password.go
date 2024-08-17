package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	validator "github.com/wagslane/go-password-validator"
)

type LoginPasswordForm struct {
	BaseNode
	progress progress.Model
	cmd      *cobra.Command
}

func ParseValue[T comparable](v *T, placeholder any, model *model, errorMsg string) tea.Cmd {
	var defaultValue T

	if tmp, ok := placeholder.(T); ok && tmp != defaultValue {
		*v = tmp
	} else {
		return formMessage(
			model,
			errorMsg,
			defaultErrorStyle,
			3*time.Second,
		)
	}

	return nil
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
			func(m *model) (bool, tea.Cmd) {
				currentNode := m.node.(*LoginPasswordForm)
				values := currentNode.fields[0].Value().([]any)

				var service string
				var login string
				var password string

				if serviceMsgCmd := ParseValue(
					&service,
					values[0],
					m,
					"Bad service. It should be non-empty and string.",
				); serviceMsgCmd != nil {
					return false, serviceMsgCmd
				}

				if loginMsgCmd := ParseValue(
					&login,
					values[1],
					m,
					"Bad login. It should be non-empty and string.",
				); loginMsgCmd != nil {
					return false, loginMsgCmd
				}

				if passwordMsgCmd := ParseValue(
					&password,
					values[2],
					m,
					"Bad password. It should be non-empty and string.",
				); passwordMsgCmd != nil {
					return false, passwordMsgCmd
				}

				m.userContext.service = service
				m.userContext.data = fmt.Sprintf("login: %s, password: %s", login, password)
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
					"Success",
					defaultMessageStyle,
					3*time.Second,
				)

				return true, tea.Batch(rollbackCmd, cmd)
			},
			defaultUnfocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginTop(1),
			defaultFocusedStyle.Width(widthForNode).AlignHorizontal(lipgloss.Center).MarginTop(1),
		),
	}

	return &LoginPasswordForm{
		BaseNode: newBaseNode(width, height, fields...),
		progress: progress.New(progress.WithDefaultGradient()),
		cmd:      command,
	}
}

func (n *LoginPasswordForm) Clear() {
	*n = *NewLoginPasswordForm(n.sizes.width, n.sizes.height, n.cmd)
}

func (n *LoginPasswordForm) Update(msg tea.Msg) (Node, tea.Cmd) {
	var cmds []tea.Cmd

	cmds = append(cmds, updateNode(&n.BaseNode, msg))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if block, isBlock := n.fields[n.cursor].(*Block); isBlock {
			if tif, isTextInputField := block.fields[block.cursor].(*TextInput); isTextInputField {
				if strings.HasPrefix(tif.textinput.Prompt, "Password") {
					entropy := validator.GetEntropy(tif.textinput.Value())
					cmd := n.progress.SetPercent(min(1, entropy/120))
					cmds = append(cmds, cmd)
				}
			}
		}
	case progress.FrameMsg:
		progressModel, cmd := n.progress.Update(msg)
		n.progress = progressModel.(progress.Model)
		cmds = append(cmds, cmd)
	}

	return n, tea.Batch(cmds...)
}

func (n LoginPasswordForm) View() string {
	return lipgloss.Place(
		n.sizes.width,
		n.sizes.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			n.BaseNode.View(),
			"\n",
			"Password Strength:\n",
			n.progress.View(),
		),
	)
}
