package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type SaltNode struct {
	BaseNode
}

func NewSaltNode(width, height int) *SaltNode {
	widthForNode := 40

	fields := []Field{
		newText(
			"Provide an environment variable with your salt or enter the salt itself. Environment variable's a higher priority.",
			lipgloss.NewStyle().Width(widthForNode).AlignHorizontal(lipgloss.Center),
		),
		newBlock(
			defaultBlockStyle.Border(lipgloss.RoundedBorder()).Padding(0, 1).Width(widthForNode),
			newTextInputField(" ENV: ", textinput.EchoPassword, defaultTextInputStyle.Width(widthForNode)),
			newText("OR", lipgloss.NewStyle().Width(widthForNode-3).AlignHorizontal(lipgloss.Center)),
			newTextInputField("SALT: ", textinput.EchoPassword, defaultTextInputStyle.Width(widthForNode)),
		),
		newChoice(
			"ENTER",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*SaltNode)
				values := currentNode.fields[1].Value().([]any)
				saltEnv := values[0].(string)
				salt := values[2].(string)

				if saltEnv == "" && salt == "" {
					cmd := formMessage(
						model,
						"Either env or salt should be non-empty.",
						defaultErrorStyle,
						3*time.Second,
					)

					return false, cmd
				} else if saltEnv != "" {
					var ok bool
					salt, ok = viper.Get(saltEnv).(string)
					if !ok {
						cmd := formMessage(
							model,
							"Couldn't get salt from env variable.",
							defaultErrorStyle,
							3*time.Second,
						)

						return false, cmd
					}
				}

				model.SetNode(NewDatabaseSelectionNode(currentNode.sizes.width, currentNode.sizes.height))
				model.userContext.salt = salt

				return true, nil
			},
			defaultUnfocusedStyle.Width(widthForNode-3).AlignHorizontal(lipgloss.Center),
			defaultFocusedStyle.Width(widthForNode-3).AlignHorizontal(lipgloss.Center),
		),
	}

	return &SaltNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}

func (n *SaltNode) Clear() {
	*n = *NewSaltNode(n.sizes.width, n.sizes.height)
}

func (n *SaltNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	cmd := updateNode(&n.BaseNode, msg)
	return n, cmd
}

func (n SaltNode) View() string {
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
