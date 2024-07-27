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

		toPrint, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(toPrint))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
