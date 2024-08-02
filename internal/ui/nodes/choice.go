package nodes

import tea "github.com/charmbracelet/bubbletea"

type Choice struct {
	name    string
	handler func() tea.Cmd
}

func (c Choice) Handle() tea.Cmd {
	return c.handler()
}
