package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hikkiyomi/passman/internal/ui/nodes"
)

type model struct {
	node nodes.Node
}

func NewModel() model {
	return model{
		node: nodes.NewWelcomeNode(),
	}
}

func (m model) Init() tea.Cmd {
	return m.node.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var msgCmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			msgCmd = m.node.HandleChoice()
			m.node = m.node.GetNext()
		}
	}

	var cmd tea.Cmd
	m.node, cmd = m.node.Update(msg)

	return m, tea.Batch(msgCmd, cmd)
}

func (m model) View() string {
	return m.node.View()
}
