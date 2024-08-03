package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Field interface {
	Init() tea.Cmd
	Update(tea.Msg) (Field, tea.Cmd)
	View() string

	Blur()
	Focus() tea.Cmd
	Handle(*model) tea.Cmd
}
