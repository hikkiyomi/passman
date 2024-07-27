package cmd

import (
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the data for some service.",
	Run: func(cmd *cobra.Command, args []string) {
		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)

		record := databases.Record{
			Owner:   user,
			Service: service,
			Data:    []byte(data),
		}

		database.Update(record)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.MarkFlagRequired("service")
	rootCmd.MarkFlagRequired("data")
}
