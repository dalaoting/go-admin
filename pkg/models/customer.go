package models

import (
	"gorm.io/gorm"

	"go-admin/common/models"
)

type Customer struct {
	gorm.Model
	models.ControlBy

	DeptId       string `json:"deptId" gorm:"type:bigint;comment:企业"`       //
	Name         string `json:"name" gorm:"type:varchar(128);comment:客户名称"` //
	AvailAmount  int64  `json:"availAmount" gorm:"type:int;comment:余额"`     //
	TotalAmount  int64  `json:"totalAmount" gorm:"type:int;comment:总金额"`    //
	PrepayAmount int64  `json:"prepayAmount" gorm:"type:int;comment:冻结金额"`  //
	Status       string `json:"status" gorm:"type:varchar(4);comment:状态"`   //
	CreateTime   string `json:"createdAt" gorm:"-"`
	CreateUser   string `json:"createBy" gorm:"-"`
	UpdateUser   string `json:"updateBy" gorm:"-"`
}

func (Customer) TableName() string {
	return "customer"
}

func (e *Customer) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Customer) GetId() interface{} {
	return e.ID
}
