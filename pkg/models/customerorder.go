package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type CustomerOrder struct {
    gorm.Model
    models.ControlBy
    
    CustomerId int64 `json:"customerId" gorm:"type:bigint;comment:客户名"` // 
    Name string `json:"name" gorm:"type:varchar(255);comment:客户名-冗余字段"` // 
    Transaction string `json:"transaction" gorm:"type:varchar(255);comment:交易单号"` // 
    TransactionType int64 `json:"transactionType" gorm:"type:int;comment:交易渠道"` // 
    Amount int64 `json:"amount" gorm:"type:int;comment:交易金额"` // 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:备注"` // 
}

func (CustomerOrder) TableName() string {
    return "customer_order"
}

func (e *CustomerOrder) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CustomerOrder) GetId() interface{} {
	return e.ID
}
