package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetches data from database.",
	Run: func(cmd *cobra.Command, args []string) {
		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)
		result := make([]databases.Record, 0)

		if service != "" {
			foundRecord := database.FindByOwnerAndService(user, service)
			if foundRecord != nil {
				result = append(result, *foundRecord)
			}
		} else {
			result = append(result, database.FindByOwner(user)...)
		}

		toPrint, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(toPrint))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	getCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	getCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	getCmd.Flags().StringVar(&service, "service", "", "specifies the service of the saving data.")

	getCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	getCmd.MarkFlagRequired("user")

	getCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	getCmd.MarkFlagRequired("password")
}
