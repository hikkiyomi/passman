package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Block struct {
	cursor int
	fields []Field
	style  lipgloss.Style
}

func newBlock(style lipgloss.Style, fields ...Field) *Block {
	return &Block{
		fields: fields,
		style:  style,
	}
}

// Moves cursor.
// Returns false if there is nowhere to move, otherwise returns true.
func (b *Block) moveCursor(step int) (bool, tea.Cmd) {
	b.fields[b.cursor].Blur()

	if b.cursor+step < 0 || b.cursor+step >= len(b.fields) {
		return false, nil
	}

	b.cursor += step

	return true, b.fields[b.cursor].Focus()
}

func (b Block) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(b.fields))

	for _, field := range b.fields {
		cmds = append(cmds, field.Init())
	}

	return tea.Batch(cmds...)
}

func (b *Block) Update(msg tea.Msg) (Field, tea.Cmd) {
	cmds := make([]tea.Cmd, 0, len(b.fields))

	for i := 0; i < len(b.fields); i++ {
		var cmd tea.Cmd
		b.fields[i], cmd = b.fields[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return b, tea.Batch(cmds...)
}

func (b Block) View() string {
	views := make([]string, 0, len(b.fields))

	for _, field := range b.fields {
		views = append(views, field.View())
	}

	toRender := lipgloss.JoinVertical(
		lipgloss.Top,
		views...,
	)

	return b.style.Render(toRender)
}

func (b *Block) Handle(model *model) (bool, tea.Cmd) {
	return b.fields[b.cursor].Handle(model)
}

func (b *Block) Blur() {
	for _, field := range b.fields {
		field.Blur()
	}
}

func (b *Block) Focus() tea.Cmd {
	return b.fields[0].Focus()
}
