package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	rootCmd.Flags().StringP("salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	rootCmd.MarkFlagRequired("salt")
}

func initConfig() {
	viper.AutomaticEnv()
}
