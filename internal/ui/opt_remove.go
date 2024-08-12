package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hikkiyomi/passman/internal/common"
	"github.com/spf13/cobra"
)

type RemoveNode struct {
	list      list.Model
	listStyle lipgloss.Style
	sizes     Sizes
	cmd       *cobra.Command
}

func NewRemoveNode(width, height int, style lipgloss.Style, command *cobra.Command) *RemoveNode {
	items := MapRecordsToItems(common.Database.FindAll())
	listWidth, listHeight := calcSizes(width, height, style)

	return &RemoveNode{
		list: NewRecordList(
			items,
			NewDelegateForRemove("x", "delete item", common.Database),
			listWidth,
			listHeight,
		),
		listStyle: style,
		sizes: Sizes{
			width:  width,
			height: height,
		},
		cmd: command,
	}
}

func (n *RemoveNode) Clear() {
	*n = *NewRemoveNode(n.sizes.width, n.sizes.height, n.listStyle, n.cmd)
}

func (n *RemoveNode) Init() tea.Cmd {
	return nil
}

func (n *RemoveNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		listWidth, listHeight := calcSizes(msg.Width, msg.Height, n.listStyle)
		n.list.SetSize(listWidth, listHeight)
	}

	var cmd tea.Cmd
	n.list, cmd = n.list.Update(msg)

	return n, cmd
}

func (n *RemoveNode) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		"Remove records and submit:",
		defaultListStyle.Render(n.list.View()),
	)
}

func (n *RemoveNode) Handle(m *model) (bool, tea.Cmd) {
	msgCmd := formMessage(
		m,
		"Success",
		defaultMessageStyle,
		3*time.Second,
	)

	cmd := m.rollbackUntil(
		NewControlPanelNode(n.sizes.width, n.sizes.height),
		func(model *model) bool {
			last := model.nodeHistory[len(model.nodeHistory)-1]
			_, ok := last.(*ControlPanelNode)

			return ok
		},
	)

	return true, tea.Batch(msgCmd, cmd)
}
