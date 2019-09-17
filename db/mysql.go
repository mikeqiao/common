package db

import (
	"github.com/mikeqiao/newant/db/mysql"
)

var DB *mysql.STypeMap

func InitDB() {
	DB = mysql.NewDBMoudle()
	if nil != DB {
		DB.InitData()
	}
}

func OnClose() {
	if nil != DB {
		DB.OnClose()
	}

}
