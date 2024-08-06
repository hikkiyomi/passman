package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	node        Node
	userContext UserContext
	nodeHistory []Node
}

func NewModel() model {
	return model{
		node: NewWelcomeNode(0, 0),
	}
}

func (m *model) SetNode(node Node) {
	m.nodeHistory = append(m.nodeHistory, m.node)
	m.node = node
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

			hasChanged, msgCmd := m.node.Handle(&m)
			cmds = append(cmds, msgCmd)

			if hasChanged {
				cmds = append(cmds, m.node.Init())
			}
		case "esc":
			length := len(m.nodeHistory)

			if length == 0 {
				break
			}

			m.node = m.nodeHistory[length-1]
			m.nodeHistory = m.nodeHistory[:length-1]
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
