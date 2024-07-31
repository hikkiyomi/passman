package cmd

import (
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
)

var (
	browser      string
	importFrom   string
	importerType string
)

var importCmd = &cobra.Command{
	Use:    "import",
	Short:  "Imports data from given file.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		importer := exporters.GetExporter(importerType, importFrom, browser)
		importingData := importer.Import()

		for _, record := range importingData {
			database.Insert(&record)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	importCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	importCmd.Flags().StringVar(&importerType, "import-type", "", "specifies the file type to import from.")
	importCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	importCmd.Flags().StringVar(&browser, "browser", "", "specifies the browser from which the data has come originally. Currently only Chrome and Firefox are supported.")

	importCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	importCmd.MarkFlagRequired("user")

	importCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	importCmd.MarkFlagRequired("password")

	importCmd.Flags().StringVar(&importFrom, "from", "", "specifies the path to import from.")
	importCmd.MarkFlagRequired("from")
}
