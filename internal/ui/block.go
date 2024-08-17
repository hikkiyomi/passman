package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Block struct {
	cursor    int
	fields    []Field
	style     lipgloss.Style
	isFocused bool
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

	_, ok := b.fields[b.cursor].(*Text)
	if ok {
		additionalStep := 1

		if step < 0 {
			additionalStep = -1
		}

		return b.moveCursor(additionalStep)
	}

	return true, b.fields[b.cursor].Focus()
}

func (b *Block) Init() tea.Cmd {
	return b.Focus()
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
		lipgloss.Left,
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

	b.isFocused = false
}

func (b *Block) Focus() tea.Cmd {
	b.isFocused = true
	return b.fields[0].Focus()
}

func (b *Block) Toggle() tea.Cmd {
	var cmd tea.Cmd

	if b.isFocused {
		b.Blur()
	} else {
		cmd = b.Focus()
	}

	return cmd
}

func (b *Block) InheritStyle(style lipgloss.Style) {
	b.style = b.style.Inherit(style)
}

func (b *Block) Value() any {
	result := make([]any, 0, len(b.fields))

	for _, field := range b.fields {
		result = append(result, field.Value())
	}

	return result
}
