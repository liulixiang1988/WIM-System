package routers

import (
	"github.com/astaxie/beego"
	"github.com/liulixiang1988/WIM-System/wim-system/controllers"
)

func init() {
	beego.Router("/", &controllers.WeightController{})
	beego.Router("/weight", &controllers.WeightController{}, "get:DailyStatistics")
}
