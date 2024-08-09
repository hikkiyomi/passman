package cmd

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/hikkiyomi/passman/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	User            string
	saltEnv         string
	Salt            string
	Path            string
	MasterPassword  string
	service         string
	data            string
	chosenEncryptor string = "aes"
	database        *databases.Database
)

func getSalt(saltEnv string) string {
	salt, ok := viper.Get(saltEnv).(string)
	if !ok {
		log.Fatal("Couldn't find any salt in provided env variable")
	}

	return salt
}

func initDatabase(cmd *cobra.Command, args []string) {
	viper.Set("user", User)

	Salt = getSalt(saltEnv)
	encryptor := encryption.GetEncryptor(chosenEncryptor, MasterPassword, Salt)
	database = databases.Open(User, Path, encryptor)
}

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
}

func initConfig() {
	viper.AutomaticEnv()
}
