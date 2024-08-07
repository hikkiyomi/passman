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

func NewFilePicker(width, height int, createNew bool) *FilePicker {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()

	// If createNew is true, then we are creating new database, hence selecting directories.
	// In these directories we will create a new database file.
	// Otherwise, we should select files which represent the database itself.
	fp.DirAllowed = createNew
	fp.FileAllowed = !createNew

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

	f.filepicker, cmd = f.filepicker.Update(msg)

	if didSelect, path := f.filepicker.DidSelectFile(msg); didSelect {
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
		s.WriteString("Pick your database file / directory, where you want to place it:")
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
	node, _ := f.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))

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
	}

	return true, tea.Batch(msgCmd)
}
