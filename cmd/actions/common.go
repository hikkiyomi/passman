package actions

import (
	"log"

	"github.com/hikkiyomi/passman/internal/common"
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

func initDatabase(cmd *cobra.Command, args []string) {
	viper.Set("user", common.User)

	if common.Salt == "" {
		common.Salt = getSalt(common.SaltEnv)
	}

	encryptor := encryption.GetEncryptor(common.ChosenEncryptor, common.MasterPassword, common.Salt)
	common.Database = databases.Open(common.User, common.Path, encryptor)
}
