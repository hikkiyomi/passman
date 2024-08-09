package ui

import "github.com/charmbracelet/lipgloss"

var (
	defaultHeaderTextStyle = lipgloss.NewStyle().MarginBottom(1)
	defaultUnfocusedStyle  = lipgloss.NewStyle().PaddingLeft(4)
	defaultFocusedStyle    = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#FFFF00"))
	defaultTextInputStyle  = lipgloss.NewStyle().Width(30)
	defaultBlockStyle      = lipgloss.NewStyle().Margin(1, 0, 1)
	defaultMessageStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#007F16"))
	defaultErrorStyle      = lipgloss.NewStyle().Background(lipgloss.Color("#9D0019"))
)
