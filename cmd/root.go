package cmd

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hikkiyomi/passman/cmd/actions"
	"github.com/hikkiyomi/passman/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "passman",
	Short: "passman is a CLI application for managing your passwords. Use `passman --help` for more information.",
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(actions.ExportCmd)
	rootCmd.AddCommand(actions.GetCmd)
	rootCmd.AddCommand(actions.ImportCmd)
	rootCmd.AddCommand(actions.RemoveCmd)
	rootCmd.AddCommand(actions.SaveCmd)
	rootCmd.AddCommand(actions.UpdateCmd)
}

func initConfig() {
	viper.AutomaticEnv()
}
