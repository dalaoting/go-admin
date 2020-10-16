package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type Customer struct {
    gorm.Model
    models.ControlBy
    
    DeptId string `json:"deptId" gorm:"type:bigint;comment:企业ID"` // 
    Name string `json:"name" gorm:"type:varchar(128);comment:客户名称"` // 
    AvailAmount string `json:"availAmount" gorm:"type:int;comment:当前余额，包含预扣冻结"` // 
    TotalAmount string `json:"totalAmount" gorm:"type:int;comment:历史入账总数"` // 
    PrepayAmount string `json:"prepayAmount" gorm:"type:int;comment:预扣冻结金额，avail_amount减去prepay_amount等于实际可用余额"` // 
    Status string `json:"status" gorm:"type:varchar(4);comment:状态"` // 
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
