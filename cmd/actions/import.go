package actions

import (
	"log"

	"github.com/hikkiyomi/passman/internal/common"
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
)

var ImportCmd = &cobra.Command{
	Use:    "import",
	Short:  "Imports data from given file.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		importer, err := exporters.GetExporter(common.ImporterType, common.ImportFrom, common.Browser)
		if err != nil {
			log.Fatal(err)
		}

		importingData := importer.Import()

		for _, record := range importingData {
			common.Database.Insert(&record)
		}
	},
}

func init() {
	ImportCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	ImportCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	ImportCmd.Flags().StringVar(&common.ImporterType, "import-type", "", "specifies the file type to import from.")
	ImportCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")
	ImportCmd.Flags().StringVar(&common.Browser, "browser", "", "specifies the browser from which the data has come originally. Currently only Chrome and Firefox are supported.")

	ImportCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	ImportCmd.MarkFlagRequired("user")

	ImportCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	ImportCmd.MarkFlagRequired("password")

	ImportCmd.Flags().StringVar(&common.ImportFrom, "from", "", "specifies the path to import from.")
	ImportCmd.MarkFlagRequired("from")
}
