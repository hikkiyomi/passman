package ui

import (
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hikkiyomi/passman/internal/common"
)

type GetNode struct {
	list      list.Model
	listStyle lipgloss.Style
	sizes     Sizes
}

func calcSizes(width, height int, style lipgloss.Style) (int, int) {
	h, v := style.GetFrameSize()
	return width - h, height - v
}

func NewGetNode(width, height int, listStyle lipgloss.Style) *GetNode {
	items := MapRecordsToItems(common.Database.FindAll())
	listWidth, listHeight := calcSizes(width, height, listStyle)

	type delegateKeymap struct {
		copy key.Binding
	}

	delegate := list.NewDefaultDelegate()
	keymap := delegateKeymap{
		copy: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "copy to clipboard"),
		),
	}
	delegateHelp := []key.Binding{keymap.copy}

	delegate.ShortHelpFunc = func() []key.Binding {
		return delegateHelp
	}
	delegate.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{delegateHelp}
	}

	lst := list.New(items, delegate, listWidth, listHeight)

	lst.KeyMap = func() list.KeyMap {
		return list.KeyMap{
			CursorUp: key.NewBinding(
				key.WithKeys("up", "k"),
				key.WithHelp("↑/k", "up"),
			),
			CursorDown: key.NewBinding(
				key.WithKeys("down", "j"),
				key.WithHelp("↓/j", "down"),
			),
			PrevPage: key.NewBinding(
				key.WithKeys("left", "h", "pgup", "b", "u"),
				key.WithHelp("←/h/pgup", "prev page"),
			),
			NextPage: key.NewBinding(
				key.WithKeys("right", "l", "pgdown", "f", "d"),
				key.WithHelp("→/l/pgdn", "next page"),
			),
			GoToStart: key.NewBinding(
				key.WithKeys("home", "g"),
				key.WithHelp("g/home", "go to start"),
			),
			GoToEnd: key.NewBinding(
				key.WithKeys("end", "G"),
				key.WithHelp("G/end", "go to end"),
			),
			Filter: key.NewBinding(
				key.WithKeys("/"),
				key.WithHelp("/", "filter"),
			),
			ClearFilter: key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "clear filter"),
			),
			CancelWhileFiltering: key.NewBinding(
				key.WithKeys("ctrl+z"),
				key.WithHelp("ctrl+z", "cancel"),
			),
			AcceptWhileFiltering: key.NewBinding(
				key.WithKeys("tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
				key.WithHelp("tab", "apply filter"),
			),
			ShowFullHelp: key.NewBinding(
				key.WithKeys("?"),
				key.WithHelp("?", "more"),
			),
			CloseFullHelp: key.NewBinding(
				key.WithKeys("?"),
				key.WithHelp("?", "close help"),
			),
		}
	}()

	return &GetNode{
		list:      lst,
		listStyle: listStyle,
		sizes: Sizes{
			width:  width,
			height: height,
		},
	}
}

func (n *GetNode) Clear() {
	*n = *NewGetNode(n.sizes.width, n.sizes.height, n.listStyle)
}

func (n *GetNode) Init() tea.Cmd {
	return nil
}

func (n *GetNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		listWidth, listHeight := calcSizes(msg.Width, msg.Height, n.listStyle)
		n.list.SetSize(listWidth, listHeight)
	}

	var cmd tea.Cmd
	n.list, cmd = n.list.Update(msg)

	return n, cmd
}

func (n *GetNode) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		defaultListStyle.Render(n.list.View()),
	)
}

func (n *GetNode) Handle(m *model) (bool, tea.Cmd) {
	var desc string

	if item, ok := n.list.SelectedItem().(item); ok {
		desc = item.Description()
	} else {
		return false, nil
	}

	var msgCmd tea.Cmd

	err := clipboard.WriteAll(desc)
	if err != nil {
		msgCmd = formMessage(
			m,
			"Couldn't save data to clipboard.",
			defaultErrorStyle,
			3*time.Second,
		)
	} else {
		msgCmd = formMessage(
			m,
			"Copied to clipboard!",
			defaultMessageStyle,
			3*time.Second,
		)
	}

	currentNode := m.node.(*GetNode)
	cmd := m.rollbackUntil(
		NewControlPanelNode(currentNode.sizes.width, currentNode.sizes.height),
		func(model *model) bool {
			last := model.nodeHistory[len(model.nodeHistory)-1]
			_, ok := last.(*ControlPanelNode)

			return ok
		},
	)

	return true, tea.Batch(cmd, msgCmd)
}
