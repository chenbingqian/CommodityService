package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	driverName := "mysql"
	dataSourceName := "root:123456@/kiki_product?charset=utf8"
	Engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	err = Engine.Ping()
	if err != nil {
		panic(err)

	}

}
