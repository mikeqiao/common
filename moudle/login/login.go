package login

import (
	ftype "github.com/mikeqiao/common/moudle/functype"
	mod "github.com/mikeqiao/newant/module"
)

var M *mod.Module

func Init() {
	M = mod.NewModule(0, 0, 1, 1)
	Register()
}

func Register() {
	M.Register(ftype.HandleVerifyAccount, HandleVerifyAccount)
	M.Register(ftype.HandCreateAccount, HandCreateAccount)
	M.Register(ftype.HandCreateUser, HandCreateUser)
}
