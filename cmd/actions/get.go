package actions

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hikkiyomi/passman/internal/common"
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:    "get",
	Short:  "Fetches data from database.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		result := make([]databases.Record, 0)

		if common.Service != "" {
			foundRecords := common.Database.FindByService(common.Service)
			if foundRecords != nil {
				result = append(result, foundRecords...)
			}
		} else {
			result = append(result, common.Database.FindAll()...)
		}

		toPrint, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(toPrint))
	},
}

func init() {
	GetCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	GetCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")
	GetCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	GetCmd.Flags().StringVar(&common.Service, "service", "", "specifies the service of the saving data.")

	GetCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	GetCmd.MarkFlagRequired("user")

	GetCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	GetCmd.MarkFlagRequired("password")
}
