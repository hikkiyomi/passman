package cmd

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
	"github.com/hikkiyomi/passman/internal/encryption"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getSalt(saltEnv string) string {
	salt, ok := viper.Get(saltEnv).(string)
	if !ok {
		log.Fatal("Couldn't find any salt in provided env variable")
	}

	return salt
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves the data for some service.",
	Run: func(cmd *cobra.Command, args []string) {
		salt := getSalt(saltEnv)
		encryptor := encryption.GetEncryptor(chosenEncryptor, masterPassword, salt)
		database := databases.Open(path, encryptor)

		record := databases.Record{
			Owner:   user,
			Service: service,
			Data:    []byte(data),
		}

		database.Insert(record)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	rootCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	rootCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	rootCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	rootCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	rootCmd.MarkFlagRequired("user")

	rootCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	rootCmd.MarkFlagRequired("password")

	rootCmd.Flags().StringVar(&service, "service", "", "specifies the service of the saving data.")
	rootCmd.MarkFlagRequired("service")

	rootCmd.Flags().StringVar(&data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	rootCmd.MarkFlagRequired("data")
}
