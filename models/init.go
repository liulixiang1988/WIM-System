package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"log"
)

var X360 *xorm.Engine
var X240 *xorm.Engine
var X40 *xorm.Engine

func init() {
	var err error
	server := beego.AppConfig.String("dev::server")
	db1 := beego.AppConfig.String("dev::db1")
	db2 := beego.AppConfig.String("dev::db2")
	db3 := beego.AppConfig.String("dev::db3")
	conStr1 := fmt.Sprintf("driver={sql server};server=%s;prot=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=%s", server, db1)
	conStr2 := fmt.Sprintf("driver={sql server};server=%s;prot=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=%s", server, db2)
	conStr3 := fmt.Sprintf("driver={sql server};server=%s;prot=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=%s", server, db3)
	X360, err = xorm.NewEngine("odbc", conStr1)
	if err != nil {
		log.Fatalf("360 err happened", err)
	}

	X240, err = xorm.NewEngine("odbc", conStr2)
	if err != nil {
		log.Fatalf("240 err happened", err)
	}
	X40, err = xorm.NewEngine("odbc", conStr3)
	if err != nil {
		log.Fatalf("40 err happened", err)
	}
}

func GetOrm(workarea int64) *xorm.Engine {
	switch workarea {
	case 0:
		return X360
	case 1:
		return X240
	case 2:
		return X40
	default:
		return X240
	}
	return X240
}
