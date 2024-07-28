package cmd

import (
	"log"

	"github.com/hikkiyomi/passman/internal/databases"
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
	Use:    "save",
	Short:  "Saves the data for some service.",
	PreRun: initDatabase,
	Run: func(cmd *cobra.Command, args []string) {
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

	saveCmd.Flags().StringVarP(&saltEnv, "salt", "s", "PASSMAN_SALT", "specifies the environment variable where the salt resides.")
	saveCmd.Flags().StringVar(&path, "path", "./database.db", "specifies the path to database.")
	saveCmd.Flags().StringVar(&chosenEncryptor, "encryptor", "aes", "specifies encryption algorithm.")

	saveCmd.Flags().StringVar(&user, "user", "", "specifies the owner of the saving data.")
	saveCmd.MarkFlagRequired("user")

	saveCmd.Flags().StringVar(&masterPassword, "password", "", "specifies the master password.")
	saveCmd.MarkFlagRequired("password")

	saveCmd.Flags().StringVar(&service, "service", "", "specifies the service of the saving data.")
	saveCmd.MarkFlagRequired("service")

	saveCmd.Flags().StringVar(&data, "data", "", "specifies the saving data. It can be login, or password, or both. Or something else.")
	saveCmd.MarkFlagRequired("data")
}
