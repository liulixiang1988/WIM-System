package routers

import (
	"github.com/liulixiang1988/WIM-System/wim-system/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
