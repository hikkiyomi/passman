package actions

import (
	"github.com/hikkiyomi/passman/internal/common"
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/spf13/cobra"
)

var SaveCmd = &cobra.Command{
	Use:    "save",
	Short:  "Saves the data for some service.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		record := databases.Record{
			Owner:   common.User,
			Service: common.Service,
			Data:    []byte(common.Data),
		}

		common.Database.Insert(&record)
	},
}

func init() {
	SaveCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	SaveCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")
	SaveCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	SaveCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	SaveCmd.MarkFlagRequired("user")

	SaveCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	SaveCmd.MarkFlagRequired("password")

	SaveCmd.Flags().StringVar(&common.Service, "service", "", "specifies the service of the saving data.")
	SaveCmd.MarkFlagRequired("service")

	SaveCmd.Flags().StringVar(&common.Data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	SaveCmd.MarkFlagRequired("data")
}
