package models

import (
	"time"
)

//明细表
type Tb_weigh_datalineinfo_detail struct {
	Id          int64
	BatchNumber int64  `xorm:"varchar(64) BatchNumber"`
	CarNumber   string `xorm:"varchar(64) CarNumber"`
	Weight      float64
	WeighTime   time.Time `xorm:"WeighTime"`
	OperateBit  int8      `xorm:"OperateBit"`
	Remark      string
	Attribute1  string
	Attribute2  string
	Attribute3  string
	Attribute4  string
	Attribute5  string
}
