package actions

import (
	"log"

	"github.com/hikkiyomi/passman/internal/common"
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

		record := common.Database.FindById(id)

		if record != nil {
			common.Database.Delete(*record)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	RemoveCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")
	RemoveCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	RemoveCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	RemoveCmd.MarkFlagRequired("user")

	RemoveCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	RemoveCmd.MarkFlagRequired("password")

	RemoveCmd.Flags().Int64("id", 0, "specifies the id of record to be deleted.")
	RemoveCmd.MarkFlagRequired("id")
}
