package helper

import (
	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger = logs.NewLogger(1)

func initLogger() {
	Logger.SetLogger("file", `{"filename":"run.log"}`)
}
