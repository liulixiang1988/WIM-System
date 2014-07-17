package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liulixiang1988/WIM-System/models/datalineinfo"
	"github.com/oal/beego-pongo2"
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
