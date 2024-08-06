package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modelKeymap struct {
	Back   key.Binding
	Submit key.Binding
	Quit   key.Binding
}

var defaultModelKeymap = modelKeymap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "exit the program"),
	),
}

func (k modelKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Submit, k.Quit}
}

func (k modelKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type Sizes struct {
	width  int
	height int
}

type model struct {
	node        Node
	userContext UserContext
	nodeHistory []Node
	help        help.Model
	keymap      modelKeymap
	sizes       Sizes
}

func NewModel() model {
	return model{
		node:   NewWelcomeNode(0, 0),
		help:   help.New(),
		keymap: defaultModelKeymap,
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
	case tea.WindowSizeMsg:
		m.sizes.width = msg.Width
		m.sizes.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Submit):
			var msgCmd tea.Cmd

			hasChanged, msgCmd := m.node.Handle(&m)
			cmds = append(cmds, msgCmd)

			if hasChanged {
				// About tea.ClearScreen and why is it necessary: read lower.
				cmds = append(cmds, m.node.Init(), tea.ClearScreen)
			}
		case key.Matches(msg, m.keymap.Back):
			length := len(m.nodeHistory)

			if length == 0 {
				break
			}

			m.node = m.nodeHistory[length-1]
			m.nodeHistory = m.nodeHistory[:length-1]

			// Perhaps it is really unnecessary.
			// Apparently without clearing the screen
			// jumping between nodes trashes your terminal output.
			// I do not really know the proper fix for this.
			// So just clear the screen entirely and re-render it without any trash.
			//
			// I also noticed that if I don't clear the screen before filepicker node,
			// I get a stutter for like 5 seconds. My terminal just lags and won't do anything.
			// With clearing the screen that does not happen. So that's a reason to do clearing.
			cmds = append(cmds, tea.ClearScreen)
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

	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.node.View(),
		lipgloss.NewStyle().
			Width(m.sizes.width).
			AlignVertical(lipgloss.Bottom).
			AlignHorizontal(lipgloss.Center).
			Render(m.help.View(m.keymap)),
	)
}
