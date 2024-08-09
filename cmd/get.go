package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:    "get",
	Short:  "Fetches data from database.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		result := make([]databases.Record, 0)

		if service != "" {
			foundRecords := database.FindByService(service)
			if foundRecords != nil {
				result = append(result, foundRecords...)
			}
		} else {
			result = append(result, database.FindAll()...)
		}

		toPrint, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(toPrint))
	},
}

func init() {
	rootCmd.AddCommand(GetCmd)

	GetCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	GetCmd.Flags().StringVar(&Path, "path", "./database.db", "specifies the path to database.")
	GetCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	GetCmd.Flags().StringVar(&service, "service", "", "specifies the service of the saving data.")

	GetCmd.Flags().StringVar(&User, "user", "", "specifies the owner of the saving data.")
	GetCmd.MarkFlagRequired("user")

	GetCmd.Flags().StringVar(&MasterPassword, "password", "", "specifies the master password.")
	GetCmd.MarkFlagRequired("password")
}
