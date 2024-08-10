package ui

import (
	"time"

	"github.com/atotto/clipboard"
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

	return &GetNode{
		list:      NewRecordList(items, NewGetDelegate(), listWidth, listHeight),
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
