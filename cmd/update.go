package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:    "update",
	Short:  "Updates the data for some service.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt64("id")
		if err != nil {
			log.Fatalf("Couldn't get id flag: %v", err)
		}

		record := database.FindById(id)

		if record != nil {
			record.Data = []byte(data)
			database.Update(*record)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	updateCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	updateCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	updateCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	updateCmd.MarkFlagRequired("user")

	updateCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	updateCmd.MarkFlagRequired("password")

	updateCmd.Flags().Int64("id", 0, "specifies the id of record to be deleted.")
	updateCmd.MarkFlagRequired("id")

	updateCmd.Flags().StringVar(&data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	updateCmd.MarkFlagRequired("data")
}
