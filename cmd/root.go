package cmd

import (
	"os"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
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
	database        *databases.Database
)

func initDatabase(cmd *cobra.Command, args []string) {
	viper.Set("user", user)

	salt := getSalt(saltEnv)
	encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
	database = databases.Open(path, encryptor)
}

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
