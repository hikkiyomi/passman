package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hikkiyomi/passman/internal/common"
	"github.com/spf13/cobra"
)

type UpdateNode struct {
	list      list.Model
	listStyle lipgloss.Style
	sizes     Sizes
	cmd       *cobra.Command
}

func NewUpdateNode(width, height int, style lipgloss.Style, command *cobra.Command) *UpdateNode {
	items := MapRecordsToItems(common.Database.FindAll())
	listWidth, listHeight := calcSizes(width, height, style)

	return &UpdateNode{
		list:      NewRecordList(items, NewDelegateWithChangedBind("enter", "update"), listWidth, listHeight),
		listStyle: style,
		sizes: Sizes{
			width:  width,
			height: height,
		},
		cmd: command,
	}
}

func (n *UpdateNode) Clear() {
	*n = *NewUpdateNode(n.sizes.width, n.sizes.height, n.listStyle, n.cmd)
}

func (n *UpdateNode) Init() tea.Cmd {
	return nil
}

func (n *UpdateNode) Update(msg tea.Msg) (Node, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		listWidth, listHeight := calcSizes(msg.Width, msg.Height, n.listStyle)
		n.list.SetSize(listWidth, listHeight)
	}

	var cmd tea.Cmd
	n.list, cmd = n.list.Update(msg)

	return n, cmd
}

func (n *UpdateNode) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		defaultListStyle.Render(n.list.View()),
	)
}

func (n *UpdateNode) Handle(m *model) (bool, tea.Cmd) {
	var it item

	if tmpItem, ok := n.list.SelectedItem().(item); ok {
		it = tmpItem
	} else {
		return false, nil
	}

	originalRun := n.cmd.Run
	n.cmd.Run = func(cmd *cobra.Command, args []string) {
		common.UpdateId = it.rawContent.Id
		defer func() {
			common.UpdateId = -1
		}()

		originalRun(cmd, args)
	}

	m.SetNode(
		NewSaveNode(
			n.sizes.width,
			n.sizes.height,
			n.cmd,
		),
	)

	return true, nil
}
