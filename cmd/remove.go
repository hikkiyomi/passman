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
	rootCmd.MarkFlagRequired("service")
}
