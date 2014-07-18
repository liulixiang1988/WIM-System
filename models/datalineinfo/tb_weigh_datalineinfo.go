package datalineinfo

import (
	"fmt"
	"github.com/liulixiang1988/WIM-System/models"
	"time"
)

type Tb_weigh_datalineinfo struct {
	Id                       int64
	WeighRecordNumber        string    `xorm:"varchar(64) WeighRecordNumber"`
	PlanCode                 string    `xorm:"varchar(64) PlanCode"`
	PlanNumber               string    `xorm:"varchar(64) PlanNumber"`
	BatchNumber              string    `xorm:"varchar(64) BatchNumber"`
	SupplierCode             string    `xorm:"varchar(64) SupplierCode"`
	SupplierName             string    `xorm:"varchar(128) SupplierName"`
	RecipientCode            string    `xorm:"varchar(64) RecipientCode"`
	RecipientName            string    `xorm:"varchar(128) RecipientName"`
	MaterialCode             string    `xorm:"varchar(64) MaterialCode"`
	MaterialName             string    `xorm:"varchar(128) MaterialName"`
	Specification            string    `xorm:"varchar(64)"`
	Model                    string    `xorm:"varchar(64)"`
	CarNumber                string    `xorm:"varchar(64) CarNumber"`
	GrossWeighHouseCode      string    `xorm:"varchar(64) GrossWeighHouseCode"`
	GrossWeighMachineCode    string    `xorm:"varchar(64) GrossWeighMachineCode"`
	GrossWeight              float64   `xorm:"GrossWeight"`
	GrossWeighTime           time.Time `xorm:"GrossWeighTime"`
	GrossWeighManCode        string    `xorm:"varchar(64) GrossWeighManCode"`
	GrossWeighManName        string    `xorm:"varchar(64) GrossWeighManName"`
	GrossWeighSupervisorCode string    `xorm:"varchar(64) GrossWeighSupervisorCode"`
	GrossWeighSupervisor     string    `xorm:"varchar(64) GrossWeighSupervisor"`
	GrossWeighShift          string    `xorm:"varchar(64) GrossWeighShift"`
	TareWeighHouseCode       string    `xorm:"varchar(64) TareWeighHouseCode"`
	TareWeighMachineCode     string    `xorm:"varchar(64) TareWeighMachineCode"`
	TareWeight               float64   `xorm:"TareWeight"`
	TareWeighTime            time.Time `xorm:"TareWeighTime"`
	TareWeighManCode         string    `xorm:"varchar(64) TareWeighManCode"`
	TareWeighManName         string    `xorm:"varchar(64) TareWeighManName"`
	TareWeighSupervisorCode  string    `xorm:"varchar(64) TareWeighSupervisorCode"`
	TareWeighSupervisor      string    `xorm:"varchar(64) TareWeighSupervisor"`
	TareWeighShift           string    `xorm:"varchar(64) TareWeighShift"`
	IsManualInputTare        bool      `xorm:"IsManualInputTare"`
	Suttle                   float64
	MeasureUnit              string    `xorm:"varchar(64) MeasureUnit"`
	WeighTime                time.Time `xorm:"WeighTime"`
	Deduction                float64
	DeductionPercent         float64   `xorm:"DeductionPercent"`
	TwiceDeduction           float64   `xorm:"TwiceDeduction"`
	TwiceDeductionPercent    float64   `xorm:"TwiceDeductionPercent"`
	MoistureContent          float64   `xorm:"MoistureContent"`
	IsManualInput            bool      `xorm:"IsManualInput"`
	OperateBit               int8      `xorm:"OperateBit"`
	IsSurplus                bool      `xorm:"IsSurplus"`
	IsReturnPurchase         bool      `xorm:"IsReturnPurchase"`
	CreateUserCode           string    `xorm:"varchar(64) CreateUserCode"`
	CreateUserName           string    `xorm:"varchar(64) CreateUserName"`
	CreateTime               time.Time `xorm:"CreateTime"`
	LastModifiedUserCode     string    `xorm:"varchar(64) LastModifiedUserCode"`
	LastModifiedUserName     string    `xorm:"varchar(64) LastModifiedUserName"`
	LastModifiedTime         time.Time `xorm:"LastModifiedTime"`
	Remark                   string
	Attribute1               string
	Attribute2               string
	Attribute3               string
	Attribute4               string
	Attribute5               string
	Attribute6               string
	Attribute7               string
	Attribute8               string
	Attribute9               string
	Attribute10              string
	Attribute11              string
	Attribute12              string
	Attribute13              string
	Attribute14              string
	Attribute15              string
}

func Last() (*Tb_weigh_datalineinfo, error) {
	last := &Tb_weigh_datalineinfo{}
	_, err := models.X.Sql("select top 1 * from tb_weigh_datalineinfo where operatebit<>2 order by id desc").Get(last)
	return last, err
}

func WorkShift(day time.Time, workshift int8) ([]*Tb_weigh_datalineinfo, error) {
	results := make([]*Tb_weigh_datalineinfo, 0)
	var beginTime, endTime time.Time
	if workshift == 0 { //早
		beginTime = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
		endTime = time.Date(day.Year(), day.Month(), day.Day(), 8, 0, 0, 0, time.UTC)
	} else if workshift == 1 { //中
		beginTime = time.Date(day.Year(), day.Month(), day.Day(), 8, 0, 0, 0, time.UTC)
		endTime = time.Date(day.Year(), day.Month(), day.Day(), 16, 0, 0, 0, time.UTC)
	} else { //晚
		beginTime = time.Date(day.Year(), day.Month(), day.Day(), 16, 0, 0, 0, time.UTC)
		endTime = time.Date(day.Year(), day.Month(), day.Day()+1, 0, 0, 0, 0, time.UTC)
	}
	fmt.Println(beginTime, endTime)
	err := models.X.Where("GrossWeighTime between ? and ?", beginTime, endTime).Find(&results)
	return results, err
}
