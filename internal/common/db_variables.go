package common

import "github.com/hikkiyomi/passman/internal/databases"

var (
	User            string
	SaltEnv         string
	Salt            string
	Path            string
	MasterPassword  string
	Service         string
	Data            string
	ChosenEncryptor string = "aes"
	Database        *databases.Database
	UpdateId        int64
)
