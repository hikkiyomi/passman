package ui

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FilePicker struct {
	filepicker filepicker.Model
	keymap     fpKeymapAdapter
	help       help.Model
	selected   string
	sizes      Sizes
	createNew  bool
}

type fpKeymapAdapter struct {
	keymap filepicker.KeyMap
}

func newFpKeymapAdapter(keymap filepicker.KeyMap) fpKeymapAdapter {
	return fpKeymapAdapter{keymap}
}

func rebindKeymap() filepicker.KeyMap {
	return filepicker.KeyMap{
		GoToTop:  key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "first")),
		GoToLast: key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "last")),
		Down:     key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "move down")),
		Up:       key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "move up")),
		PageUp:   key.NewBinding(key.WithKeys("pgup"), key.WithHelp("pgup", "page up")),
		PageDown: key.NewBinding(key.WithKeys("pgdown"), key.WithHelp("pgdown", "page down")),
		Back:     key.NewBinding(key.WithKeys("left"), key.WithHelp("←", "back")),
		Open:     key.NewBinding(key.WithKeys("right", "enter"), key.WithHelp("→/enter", "open")),
		Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	}
}

func NewFilePicker(width, height int, createNew bool) *FilePicker {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()

	// If createNew is true, then we are creating new database, hence selecting directories.
	// In these directories we will create a new database file.
	// Otherwise, we should select files which represent the database itself.
	fp.DirAllowed = createNew
	fp.FileAllowed = !createNew

	fp.KeyMap = rebindKeymap()

	fp, _ = fp.Update(tea.WindowSizeMsg{
		Width:  width,
		Height: int(lipgloss.Position(height) * lipgloss.Center),
	})

	return &FilePicker{
		filepicker: fp,
		keymap:     newFpKeymapAdapter(fp.KeyMap),
		help:       help.New(),
		sizes: Sizes{
			width:  width,
			height: height,
		},
		createNew: createNew,
	}
}

func (f *FilePicker) Init() tea.Cmd {
	return f.filepicker.Init()
}

func (f *FilePicker) Update(msg tea.Msg) (Node, tea.Cmd) {
	var cmd tea.Cmd
	var modifiedMsg tea.Msg

	if tmpMsg, ok := msg.(tea.WindowSizeMsg); ok {
		f.sizes.height = tmpMsg.Height
		f.sizes.width = tmpMsg.Width

		tmpMsg.Height = int(lipgloss.Position(f.sizes.height) * lipgloss.Center)
		modifiedMsg = tmpMsg
	} else {
		modifiedMsg = msg
	}

	f.filepicker, cmd = f.filepicker.Update(modifiedMsg)

	if didSelect, path := f.filepicker.DidSelectFile(modifiedMsg); didSelect {
		f.selected = path
	}

	return f, cmd
}

func (fk fpKeymapAdapter) ShortHelp() []key.Binding {
	return []key.Binding{
		fk.keymap.Back,
		fk.keymap.Down,
		fk.keymap.Up,
		fk.keymap.Open,
		fk.keymap.Select,
	}
}

func (fk fpKeymapAdapter) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		fk.ShortHelp(),
		{fk.keymap.GoToLast, fk.keymap.GoToTop},
		{fk.keymap.PageDown, fk.keymap.PageUp},
	}
}

func (f FilePicker) renderHelp() string {
	return f.help.View(f.keymap)
}

func (f FilePicker) View() string {
	var s strings.Builder

	if f.selected == "" {
		if f.createNew {
			s.WriteString("Pick directory:")
		} else {
			s.WriteString("Pick file:")
		}
	}

	s.WriteString("\n\n" + f.filepicker.View() + "\n\n\n" + f.renderHelp())

	return lipgloss.Place(
		f.sizes.width,
		f.sizes.height,
		lipgloss.Center,
		lipgloss.Center,
		s.String(),
	)
}

func (f FilePicker) chooseErrorMessage() string {
	if f.createNew {
		return "You can only select directories while creating new database."
	}

	return "You can only select files while opening existing database."
}

func (f *FilePicker) Handle(model *model) (bool, tea.Cmd) {
	var msgCmd tea.Cmd = nil
	node, cmd := f.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))

	tempNode, ok := node.(*FilePicker)
	if !ok {
		log.Fatal("Couldn't convert node into FilePicker")
	}

	f = tempNode

	if f.selected == "" {
		msgCmd = formMessage(
			model,
			f.chooseErrorMessage(),
			defaultErrorStyle,
			3*time.Second,
		)
	} else {
		model.userContext.path = f.selected
		msgCmd = formMessage(
			model,
			"You picked: "+f.selected,
			defaultMessageStyle,
			3*time.Second,
		)
	}

	return true, tea.Batch(cmd, msgCmd)
}
