package cmd

import (
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
)

var (
	importFrom   string
	importerType string
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports data from given file.",
	Run: func(cmd *cobra.Command, args []string) {
		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)
		importer := exporters.GetExporterByType(importerType, importFrom)
		importingData := importer.Import()

		for _, record := range importingData {
			database.Insert(record)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	importCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	importCmd.Flags().StringVar(&importerType, "import-type", "csv", "specifies the file type to import from.")

	importCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	importCmd.MarkFlagRequired("user")

	importCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	importCmd.MarkFlagRequired("password")

	importCmd.Flags().StringVar(&importFrom, "from", "", "specifies the path to import from.")
	importCmd.MarkFlagRequired("from")
}
