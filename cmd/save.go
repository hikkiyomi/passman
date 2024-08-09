package cmd

import (
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/spf13/cobra"
)

var SaveCmd = &cobra.Command{
	Use:    "save",
	Short:  "Saves the data for some service.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		record := databases.Record{
			Owner:   User,
			Service: service,
			Data:    []byte(data),
		}

		database.Insert(&record)
	},
}

func init() {
	rootCmd.AddCommand(SaveCmd)

	SaveCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	SaveCmd.Flags().StringVar(&Path, "path", "./database.db", "specifies the path to database.")
	SaveCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	SaveCmd.Flags().StringVar(&User, "user", "", "specifies the owner of the saving data.")
	SaveCmd.MarkFlagRequired("user")

	SaveCmd.Flags().StringVar(&MasterPassword, "password", "", "specifies the master password.")
	SaveCmd.MarkFlagRequired("password")

	SaveCmd.Flags().StringVar(&service, "service", "", "specifies the service of the saving data.")
	SaveCmd.MarkFlagRequired("service")

	SaveCmd.Flags().StringVar(&data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	SaveCmd.MarkFlagRequired("data")
}
