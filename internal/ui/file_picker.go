package ui

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type filePickerAdapter struct {
	filepicker filepicker.Model
	isFocused  bool
}

func newFilePickerAdapter(filepicker filepicker.Model) *filePickerAdapter {
	return &filePickerAdapter{
		filepicker: filepicker,
		isFocused:  false,
	}
}

func (f *filePickerAdapter) Init() tea.Cmd {
	cmd := f.Focus()
	return tea.Batch(f.filepicker.Init(), cmd)
}

func (f *filePickerAdapter) Update(msg tea.Msg) (*filePickerAdapter, tea.Cmd) {
	_, ok := msg.(tea.WindowSizeMsg)

	if !f.isFocused && !ok {
		return f, nil
	}

	var cmd tea.Cmd
	f.filepicker, cmd = f.filepicker.Update(msg)

	return f, cmd
}

func (f *filePickerAdapter) View() string {
	return f.filepicker.View()
}

func (f *filePickerAdapter) Blur() {
	f.isFocused = false
}

func (f *filePickerAdapter) Focus() tea.Cmd {
	f.isFocused = true
	return nil
}

func (f *filePickerAdapter) Toggle() tea.Cmd {
	var cmd tea.Cmd

	if f.isFocused {
		f.Blur()
	} else {
		cmd = f.Focus()
	}

	return cmd
}

func (f *filePickerAdapter) DidSelectFile(msg tea.Msg) (bool, string) {
	if !f.isFocused {
		return false, ""
	}

	return f.filepicker.DidSelectFile(msg)
}

type FilePicker struct {
	filePickerAdapter *filePickerAdapter
	textinput         *TextInput

	keymap    fpKeymapAdapter
	help      help.Model
	selected  string
	sizes     Sizes
	createNew bool
}

type fpKeymapAdapter struct {
	fpKeymap  filepicker.KeyMap
	switchKey key.Binding
}

func newFpKeymapAdapter(keymap filepicker.KeyMap) fpKeymapAdapter {
	return fpKeymapAdapter{
		fpKeymap:  keymap,
		switchKey: key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "switch focus")),
	}
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

func NewFilePicker(
	width int,
	height int,
	createNew bool,
	textInputField *TextInput,
) *FilePicker {
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

	adapter := newFilePickerAdapter(fp)

	return &FilePicker{
		filePickerAdapter: adapter,
		textinput:         textInputField,
		keymap:            newFpKeymapAdapter(fp.KeyMap),
		help:              help.New(),
		sizes: Sizes{
			width:  width,
			height: height,
		},
		createNew: createNew,
	}
}

func (f *FilePicker) Init() tea.Cmd {
	return f.filePickerAdapter.Init()
}

func (f *FilePicker) Update(msg tea.Msg) (Node, tea.Cmd) {
	var cmds []tea.Cmd
	var modifiedMsg tea.Msg

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		f.sizes.height = msg.Height
		f.sizes.width = msg.Width
		msg.Height = int(lipgloss.Position(f.sizes.height) * lipgloss.Center)
		modifiedMsg = msg
	case tea.KeyMsg:
		if key.Matches(msg, f.keymap.switchKey) && f.textinput != nil {
			f.filePickerAdapter.Toggle()
			f.textinput.Toggle()
		}

		modifiedMsg = msg
	default:
		modifiedMsg = msg
	}

	var cmd tea.Cmd
	f.filePickerAdapter, cmd = f.filePickerAdapter.Update(modifiedMsg)
	cmds = append(cmds, cmd)

	if f.textinput != nil {
		tempField, cmd := f.textinput.Update(modifiedMsg)
		f.textinput = tempField.(*TextInput)
		cmds = append(cmds, cmd)
	}

	if didSelect, path := f.filePickerAdapter.DidSelectFile(modifiedMsg); didSelect {
		f.selected = path
	}

	return f, tea.Batch(cmds...)
}

func (fk fpKeymapAdapter) ShortHelp() []key.Binding {
	return []key.Binding{
		fk.fpKeymap.Back,
		fk.fpKeymap.Down,
		fk.fpKeymap.Up,
		fk.fpKeymap.Open,
		fk.fpKeymap.Select,
		fk.switchKey,
	}
}

func (fk fpKeymapAdapter) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		fk.ShortHelp(),
		{fk.fpKeymap.GoToLast, fk.fpKeymap.GoToTop},
		{fk.fpKeymap.PageDown, fk.fpKeymap.PageUp},
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

	s.WriteString("\n\n")
	s.WriteString(f.filePickerAdapter.View())
	s.WriteString("\n\n\n")

	if f.textinput != nil {
		s.WriteString(f.textinput.View())
		s.WriteString("\n\n\n")
	}

	s.WriteString(f.renderHelp())

	return lipgloss.Place(
		f.sizes.width,
		f.sizes.height,
		lipgloss.Center,
		lipgloss.Center,
		s.String(),
	)
}

func (f *FilePicker) checkIfNameIsEmpty() bool {
	if f.textinput == nil {
		return false
	}

	filename := f.textinput.Value().(string)

	return filename == ""
}

func (f FilePicker) chooseErrorMessage() string {
	if f.checkIfNameIsEmpty() {
		return "Please provide a name for your database."
	}

	if f.createNew {
		return "You can only select directories while creating new database."
	}

	return "You can only select files while opening existing database."
}

func (f *FilePicker) formPath() string {
	if f.createNew {
		filename := f.textinput.Value().(string)
		return filepath.Join(f.selected, filename)
	}

	return f.selected
}

func (f *FilePicker) Handle(model *model) (bool, tea.Cmd) {
	if !f.filePickerAdapter.isFocused {
		return false, nil
	}

	var msgCmd tea.Cmd = nil
	node, cmd := f.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))

	tempNode, ok := node.(*FilePicker)
	if !ok {
		log.Fatal("Couldn't convert node into FilePicker")
	}

	f = tempNode

	if f.selected == "" || f.checkIfNameIsEmpty() {
		msgCmd = formMessage(
			model,
			f.chooseErrorMessage(),
			defaultErrorStyle,
			3*time.Second,
		)
	} else {
		model.userContext.path = f.formPath()

		msgCmd = formMessage(
			model,
			"Opened database at "+model.userContext.path,
			defaultMessageStyle,
			3*time.Second,
		)

		currentNode := model.node.(*FilePicker)
		model.SetNode(NewControlPanelNode(currentNode.sizes.width, currentNode.sizes.height))
	}

	return true, tea.Batch(cmd, msgCmd)
}
