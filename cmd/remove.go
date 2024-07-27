package cmd

import (
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes the data for some service.",
	Run: func(cmd *cobra.Command, args []string) {
		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)

		database.DeleteByOwnerAndService(user, service)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	removeCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	removeCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	removeCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	removeCmd.MarkFlagRequired("user")

	removeCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	removeCmd.MarkFlagRequired("password")

	removeCmd.Flags().StringVar(&service, "service", "", "specifies the service of the saving data.")
	removeCmd.MarkFlagRequired("service")
}
