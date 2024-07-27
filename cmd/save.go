package cmd

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getSalt(saltEnv string) string {
	salt, ok := viper.Get(saltEnv).(string)
	if !ok {
		log.Fatal("Couldn't find any salt in provided env variable")
	}

	return salt
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves the data for some service.",
	Run: func(cmd *cobra.Command, args []string) {
		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)

		record := databases.Record{
			Owner:   user,
			Service: service,
			Data:    []byte(data),
		}

		database.Insert(record)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	rootCmd.MarkFlagRequired("service")
	rootCmd.MarkFlagRequired("data")
}
