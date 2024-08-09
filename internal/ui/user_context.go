package ui

import "github.com/hikkiyomi/passman/internal/common"

type UserContext struct {
	login    string
	password string
	salt     string
	path     string
	data     string
	service  string
}

func MapUserContextToDatabaseVariables(ctx UserContext) {
	common.User = ctx.login
	common.MasterPassword = ctx.password
	common.Salt = ctx.salt
	common.Path = ctx.path
	common.Data = ctx.data
	common.Service = ctx.service
}
