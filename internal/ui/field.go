package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Field interface {
	Init() tea.Cmd
	Update(tea.Msg) (Field, tea.Cmd)
	View() string

	Blur()
	Focus() tea.Cmd
	Toggle() tea.Cmd

	Handle(*model) (bool, tea.Cmd)
	InheritStyle(lipgloss.Style)
	Value() any
}
