package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	saltEnv         string
	path            string
	user            string
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

	rootCmd.PersistentFlags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	rootCmd.PersistentFlags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	rootCmd.PersistentFlags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	rootCmd.PersistentFlags().StringVar(&service, "service", "", "specifies the service of the saving data.")
	rootCmd.PersistentFlags().StringVar(&data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")

	rootCmd.PersistentFlags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	rootCmd.MarkFlagRequired("user")

	rootCmd.PersistentFlags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	rootCmd.MarkFlagRequired("password")
}

func initConfig() {
	viper.AutomaticEnv()
}
