package ui

import "github.com/charmbracelet/lipgloss"

var (
	defaultUnfocusedStyle = lipgloss.NewStyle().PaddingLeft(4)
	defaultFocusedStyle   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#FFFF00"))
	defaultTextInputStyle = lipgloss.NewStyle().Width(30)
	defaultBlockStyle     = lipgloss.NewStyle().MarginTop(1).MarginBottom(1)
)
