package config

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	ConnectionString = ""
	ConstAlias       = "default"
	ConstDriver      = "mysql"
)

func Connect() {
	_, err := orm.GetDB(ConstAlias)

	if err != nil {
		errInRegister := orm.RegisterDataBase(ConstAlias, ConstDriver, ConnectionString)
		if errInRegister != nil {
			panic(errInRegister)
		}
	}
}
