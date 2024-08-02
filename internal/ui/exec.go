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
		node: nodes.NewWelcomeNode(0, 0),
	}
}

func (m model) Init() tea.Cmd {
	return m.node.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			var msgCmd tea.Cmd

			m.node, msgCmd = m.node.Handle()
			m.node = m.node.GetNext()
			cmds = append(cmds, msgCmd)

			if m.node != nil {
				cmds = append(cmds, m.node.Init())
			}
		}
	}

	var cmd tea.Cmd

	if m.node != nil {
		m.node, cmd = m.node.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.node == nil {
		return ""
	}

	return m.node.View()
}