package gate

import (
	mod "github.com/mikeqiao/ant/module"
	ftype "github.com/mikeqiao/common/moudle/functype"
)

var M *mod.Module

func Init() {
	M = mod.NewModule(0, 0, 1, 1)
	Register()
}

func Register() {
	M.Register(ftype.HandleUserLogin, HandleUserLogin)
}
