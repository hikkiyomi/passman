package actions

import (
	"github.com/hikkiyomi/passman/internal/common"
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
)

var ExportCmd = &cobra.Command{
	Use:    "export",
	Short:  "Exports data into given file.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		exporter := exporters.GetExporter(common.ExporterType, common.ExportInto, "")
		exportingData := common.Database.FindAll()

		exporter.Export(exportingData)
	},
}

func init() {
	ExportCmd.Flags().StringVarP(&common.SaltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	ExportCmd.Flags().StringVar(&common.ChosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	ExportCmd.Flags().StringVar(&common.ExporterType, "export-type", "", "specifies the file type to export into.")
	ExportCmd.Flags().StringVar(&common.Path, "path", "./database.db", "specifies the path to database.")

	ExportCmd.Flags().StringVar(&common.User, "user", "", "specifies the owner of the saving data.")
	ExportCmd.MarkFlagRequired("user")

	ExportCmd.Flags().StringVar(&common.MasterPassword, "password", "", "specifies the master password.")
	ExportCmd.MarkFlagRequired("password")

	ExportCmd.Flags().StringVar(&common.ExportInto, "into", "", "specifies the path to export into.")
	ExportCmd.MarkFlagRequired("into")
}
