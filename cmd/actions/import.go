package actions

import (
	"github.com/hikkiyomi/passman/internal/common"
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
)

var (
	browser      string
	importFrom   string
	importerType string
)

var ImportCmd = &cobra.Command{
	Use:    "import",
	Short:  "Imports data from given file.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		importer := exporters.GetExporter(importerType, importFrom, browser)
		importingData := importer.Import()

		for _, record := range importingData {
			common.Database.Insert(&record)
		}
	},
}

func init() {
	ImportCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	ImportCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	ImportCmd.Flags().StringVar(&importerType, "import-type", "", "specifies the file type to import from.")
	ImportCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")
	ImportCmd.Flags().StringVar(&browser, "browser", "", "specifies the browser from which the data has come originally. Currently only Chrome and Firefox are supported.")

	ImportCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	ImportCmd.MarkFlagRequired("user")

	ImportCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	ImportCmd.MarkFlagRequired("password")

	ImportCmd.Flags().StringVar(&importFrom, "from", "", "specifies the path to import from.")
	ImportCmd.MarkFlagRequired("from")
}
