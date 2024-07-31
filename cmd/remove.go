package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:    "remove",
	Short:  "Removes the data for some service.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt64("id")
		if err != nil {
			log.Fatalf("Couldn't get id flag: %v", err)
		}

		record := database.FindById(id)

		if record != nil {
			database.Delete(*record)
		}
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

	removeCmd.Flags().Int64("id", 0, "specifies the id of record to be deleted.")
	removeCmd.MarkFlagRequired("id")
}
