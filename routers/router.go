package routers

import (
	"github.com/astaxie/beego"
	"github.com/liulixiang1988/WIM-System/controllers"
)

func init() {
	beego.Router("/", &controllers.WeightController{}, "get:WorkShift")
	beego.Router("/weight", &controllers.WeightController{}, "get:DailyStatistics")
	beego.Router("/workshift", &controllers.WeightController{}, "get:WorkShift")
	beego.Router("/details", &controllers.WeightController{}, "get:GetDetails")
	beego.Router("/statistics", &controllers.WeightController{}, "get:GetStatics")
}
