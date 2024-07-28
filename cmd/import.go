package cmd

import (
	"fmt"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	browser      string
	importFrom   string
	importerType string
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports data from given file.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("user", user)

		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)
		importer := exporters.GetExporter(importerType, importFrom, browser)
		importingData := importer.Import()

		countService := map[string]int{}

		// In case if importing data contains more than 1 record for one service.
		renameService := func(service string) string {
			count := countService[service]

			if count >= 2 {
				service = fmt.Sprintf("%s (%d)", service, count-1)
			}

			return service
		}

		for _, record := range importingData {
			countService[record.Service]++
			record.Service = renameService(record.Service)
			database.Insert(record)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	importCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	importCmd.Flags().StringVar(&importerType, "import-type", "csv", "specifies the file type to import from.")
	importCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	importCmd.Flags().StringVar(&browser, "browser", "", "specifies the browser from which the data has come originally. Currently only Chrome and Firefox are supported.")

	importCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	importCmd.MarkFlagRequired("user")

	importCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	importCmd.MarkFlagRequired("password")

	importCmd.Flags().StringVar(&importFrom, "from", "", "specifies the path to import from.")
	importCmd.MarkFlagRequired("from")
}
