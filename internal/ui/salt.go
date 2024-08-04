package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type SaltNode struct {
	BaseNode
}

func NewSaltNode(width, height int) *SaltNode {
	fields := []Field{
		newBlock(
			defaultBlockStyle.Border(lipgloss.RoundedBorder()).Padding(0, 1, 0),
			newTextInputField(defaultTextInputStyle, "SALT ENV: ", textinput.EchoPassword),
			// TODO: MAKE FIELD FOR STATIC TEXT
			newTextInputField(defaultTextInputStyle, "    SALT: ", textinput.EchoPassword),
		),
		newChoice(
			defaultUnfocusedStyle,
			defaultFocusedStyle,
			"ENTER",
			func(model *model) (bool, tea.Cmd) {
				currentNode := model.node.(*SaltNode)
				model.node = nil

				values := currentNode.fields[0].Value().([]any)
				saltEnv := values[0].(string)
				salt := values[1].(string)

				if saltEnv == "" && salt == "" {
					log.Fatal("Either salt env or salt should be passed.")
				} else if saltEnv != "" {
					var ok bool
					salt, ok = viper.Get(saltEnv).(string)
					if !ok {
						log.Fatal("Couldn't cast salt from saltEnv to string.")
					}
				}

				model.userContext.salt = salt

				return true, nil
			},
		),
	}

	return &SaltNode{
		BaseNode: newBaseNode(width, height, fields...),
	}
}
