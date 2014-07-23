package models

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
)

var X *xorm.Engine

func init() {
	var err error
	X, err = xorm.NewEngine("odbc", "driver={sql server};server=172.16.57.198;port=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=tgzljl_czgl")
	if err != nil {
		log.Fatalf("err happened", err)
	}
}
