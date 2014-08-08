package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/liulixiang1988/WIM-System/helper"
	"github.com/liulixiang1988/WIM-System/models/datalineinfo"
	"github.com/liulixiang1988/WIM-System/models/detail"
	"github.com/oal/beego-pongo2"
	"time"
)

type WeightController struct {
	beego.Controller
}

func (this *WeightController) Get() {
	last, err := datalineinfo.Last(1)
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

//按班次查询
func (this *WeightController) WorkShift() {
	data := pongo2.Context{"title": "班次明细查询"}
	var msg []string = make([]string, 0)

	dayStr := this.GetString("day")
	if dayStr == "" {
		dayStr = time.Now().Format("2006-01-02")
	}
	data["day"] = dayStr

	workshift, err := this.GetInt("workshift")
	if err != nil {
		workshift = 1
	}
	data["workshift"] = workshift

	workarea, err := this.GetInt("workarea")
	if err != nil {
		workarea = 1
	}
	data["workarea"] = workarea

	valid := validation.Validation{}
	valid.Match(dayStr, helper.DatePatten, "day")
	valid.Required(workshift, "workshift")
	valid.Required(workarea, "workarea")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			msg = append(msg, err.String())
		}
		data["msg"] = msg
		pongo2.Render(this.Ctx, "weight/workshift.html", data)
		return
	}

	day, _ := time.Parse("2006-01-02", dayStr)

	results, err := datalineinfo.WorkShift(day, int8(workshift), workarea)
	if err != nil {
		msg = append(msg, "请选择日期与班次")
		msg = append(msg, err.Error())
		data["msg"] = msg
	}
	data["results"] = results

	pongo2.Render(this.Ctx, "weight/workshift.html", data)
}

func (this *WeightController) GetDetails() {
	data := pongo2.Context{"title": "批次明细"}
	var msg []string = make([]string, 0)
	batchNumber, err := this.GetInt("batch")
	if err != nil {
		batchNumber = 0
	}
	data["batch"] = batchNumber

	workarea, err := this.GetInt("workarea")
	if err != nil {
		workarea = 1
	}
	data["workarea"] = workarea

	results, err := detail.GetDetails(workarea, batchNumber)
	if err != nil {
		msg = append(msg, err.Error())
		data["msg"] = msg
	}
	data["results"] = results
	pongo2.Render(this.Ctx, "weight/details.html", data)
}

func (this *WeightController) GetStatics() {
	data := pongo2.Context{"title": "时间范围统计"}
	var msg []string = make([]string, 0)

}
