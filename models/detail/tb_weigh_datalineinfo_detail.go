package detail

import (
	"github.com/liulixiang1988/WIM-System/models"
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

func GetDetails(workarea int64, batchNumber int64) ([]*Tb_weigh_datalineinfo_detail, error) {
	results := make([]*Tb_weigh_datalineinfo_detail, 0)
	err := models.GetOrm(workarea).Where("BatchNumber = ? or BatchNumber = ?", batchNumber, batchNumber+1).Asc("Id").Find(&results)
	return results, err
}
