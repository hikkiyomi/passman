package ui

import "github.com/charmbracelet/lipgloss"

var (
	itemStyle         = lipgloss.NewStyle()
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
	textInputStyle    = lipgloss.NewStyle().Width(30)
)
