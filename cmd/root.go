package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	user            string
	saltEnv         string
	path            string
	masterPassword  string
	service         string
	data            string
	chosenEncryptor string
)

var rootCmd = &cobra.Command{
	Use:   "passman",
	Short: "passman is a CLI application for managing your passwords. Use `passman --help` for more information.",
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
