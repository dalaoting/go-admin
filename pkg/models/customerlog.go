package models

import (
	"go-admin/common/models"
	"gorm.io/gorm"
)

type CustomerLog struct {
	gorm.Model
	models.ControlBy

	CustomerId   int64  `json:"customerId" gorm:"type:bigint;comment:客户名"` //
	Amount       int64  `json:"amount"`                                    // 本次操作金额
	OpType       int    `json:"opType"`                                    // 操作类型 1增加、2减少
	BsType       int    `json:"bsType"`                                    // 业务类型
	Detail       string `json:"detail"`                                    // 描述
	BeforeAmount int64  `json:"beforeAmount"`                              // 操作前的金额
	AfterAmount  int64  `json:"afterAmount"`                               // 操作后的金额
	Ext          string `json:"ext"`                                       // 拓展信息

	CreateTime string `json:"createdAt" gorm:"-"`
	CreateUser string `json:"createBy" gorm:"-"`
	UpdateUser string `json:"updateBy" gorm:"-"`
}

func (CustomerLog) TableName() string {
	return "customer_log"
}

func (e *CustomerLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CustomerLog) GetId() interface{} {
	return e.ID
}
