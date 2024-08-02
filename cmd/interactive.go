package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hikkiyomi/passman/internal/ui"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Runs user-friendly interface.",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := tea.LogToFile("debug.log", "DEBUG")
		if err != nil {
			log.Fatalf("Couldn't handle logging to file: %v", err)
		}
		defer f.Close()

		p := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			log.Fatalf("Something went wrong while running tea.Program: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
