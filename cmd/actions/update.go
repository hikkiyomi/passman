package actions

import (
	"github.com/hikkiyomi/passman/internal/common"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:    "update",
	Short:  "Updates the data for some service.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		record := common.Database.FindById(common.UpdateId)

		if record != nil {
			record.Data = []byte(common.Data)
			common.Database.Update(*record)
		}
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	UpdateCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")
	UpdateCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	UpdateCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	UpdateCmd.MarkFlagRequired("user")

	UpdateCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	UpdateCmd.MarkFlagRequired("password")

	UpdateCmd.Flags().Int64Var(&common.UpdateId, "id", 0, "specifies the id of record to be deleted.")
	UpdateCmd.MarkFlagRequired("id")

	UpdateCmd.Flags().StringVar(&common.Data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	UpdateCmd.MarkFlagRequired("data")
}
