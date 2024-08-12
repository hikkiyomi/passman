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
	UpdateId        int64 = -1
	ImporterType    string
	ExporterType    string
	ImportFrom      string
	ExportInto      string
	Browser         string
)
