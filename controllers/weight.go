package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liulixiang1988/WIM-System/models/datalineinfo"
	"github.com/oal/beego-pongo2"
	"time"
)

type WeightController struct {
	beego.Controller
}

func (this *WeightController) Get() {
	last, err := datalineinfo.Last()
	var msg string
	if err != nil {
		msg = "error"
	}
	pongo2.Render(this.Ctx, "weight/home.html", pongo2.Context{
		"title":        "首页",
		"msg":          msg,
		"datalineinfo": last,
	})
}

func (this *WeightController) DailyStatistics() {
	pongo2.Render(this.Ctx, "weight/daily_statistics.html", pongo2.Context{
		"title": "日统计",
	})
}

func (this *WeightController) WorkShift() {
	var msg string
	dayStr := this.GetString(":day")
	day, _ := time.Parse("2006-01-02", dayStr)
	workshift, _ := this.GetInt(":workshift")
	results, err := datalineinfo.WorkShift(day, int8(workshift))
	if err != nil {
		msg = "err"
	}
	this.Data["json"] = map[string]interface{}{
		"msg":     msg,
		"results": &results,
	}
	this.ServeJson()
	// pongo2.Render(this.Ctx, "weight/workshift.html", pongo2.Context{
	// 	"title":   "班次明细查询",
	// 	"msg":     msg,
	// 	"results": results,
	// })
}
