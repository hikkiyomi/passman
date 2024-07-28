package cmd

import (
	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/hikkiyomi/passman/internal/exporters"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	exportInto   string
	exporterType string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports data into given file.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("user", user)

		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)
		exporter := exporters.GetExporter(exporterType, exportInto, "")
		exportingData := database.FindByOwner(user)

		exporter.Export(exportingData)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	exportCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")
	exportCmd.Flags().StringVar(&exporterType, "export-type", "", "specifies the file type to export into.")
	exportCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")

	exportCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	exportCmd.MarkFlagRequired("user")

	exportCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	exportCmd.MarkFlagRequired("password")

	exportCmd.Flags().StringVar(&exportInto, "into", "", "specifies the path to export into.")
	exportCmd.MarkFlagRequired("into")
}
