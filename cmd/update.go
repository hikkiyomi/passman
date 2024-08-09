package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
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
	rootCmd.AddCommand(UpdateCmd)

	UpdateCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	UpdateCmd.Flags().StringVar(&Path, "path", "./database.db", "specifies the path to database.")
	UpdateCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	UpdateCmd.Flags().StringVar(&User, "user", "", "specifies the owner of the saving data.")
	UpdateCmd.MarkFlagRequired("user")

	UpdateCmd.Flags().StringVar(&MasterPassword, "password", "", "specifies the master password.")
	UpdateCmd.MarkFlagRequired("password")

	UpdateCmd.Flags().Int64("id", 0, "specifies the id of record to be deleted.")
	UpdateCmd.MarkFlagRequired("id")

	UpdateCmd.Flags().StringVar(&data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	UpdateCmd.MarkFlagRequired("data")
}
