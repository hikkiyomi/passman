package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
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
	rootCmd.AddCommand(RemoveCmd)

	RemoveCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	RemoveCmd.Flags().StringVar(&Path, "path", "./database.db", "specifies the path to database.")
	RemoveCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	RemoveCmd.Flags().StringVar(&User, "user", "", "specifies the owner of the saving data.")
	RemoveCmd.MarkFlagRequired("user")

	RemoveCmd.Flags().StringVar(&MasterPassword, "password", "", "specifies the master password.")
	RemoveCmd.MarkFlagRequired("password")

	RemoveCmd.Flags().Int64("id", 0, "specifies the id of record to be deleted.")
	RemoveCmd.MarkFlagRequired("id")
}
