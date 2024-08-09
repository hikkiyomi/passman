package cmd

import (
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
)

var (
	exportInto   string
	exporterType string
)

var ExportCmd = &cobra.Command{
	Use:    "export",
	Short:  "Exports data into given file.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
		exporter := exporters.GetExporter(exporterType, exportInto, "")
		exportingData := database.FindAll()

		exporter.Export(exportingData)
	},
}

func init() {
	rootCmd.AddCommand(ExportCmd)

	ExportCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	ExportCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	ExportCmd.Flags().StringVar(&exporterType, "export-type", "", "specifies the file type to export into.")
	ExportCmd.Flags().StringVar(&Path, "path", "./database.db", "specifies the path to database.")

	ExportCmd.Flags().StringVar(&User, "user", "", "specifies the owner of the saving data.")
	ExportCmd.MarkFlagRequired("user")

	ExportCmd.Flags().StringVar(&MasterPassword, "password", "", "specifies the master password.")
	ExportCmd.MarkFlagRequired("password")

	ExportCmd.Flags().StringVar(&exportInto, "into", "", "specifies the path to export into.")
	ExportCmd.MarkFlagRequired("into")
}
