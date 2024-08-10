package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type delegateKeymap struct {
	copy key.Binding
}

func NewGetDelegate() list.DefaultDelegate {
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

	return delegate
}

func NewRecordList(
	items []list.Item,
	delegate list.DefaultDelegate,
	width int,
	height int,
) list.Model {
	result := list.New(items, delegate, width, height)

	result.KeyMap = func() list.KeyMap {
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

	return result
}
