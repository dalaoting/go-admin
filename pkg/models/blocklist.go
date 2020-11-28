package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type BlockList struct {
    gorm.Model
    models.ControlBy
    
    Account string `json:"account" gorm:"type:varchar(64);comment:账号"` // 
    Type string `json:"type" gorm:"type:int;comment:类型"` // 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:备注，封禁原因"` // 
    Origin string `json:"origin" gorm:"type:int;comment:来源"` // 
}

func (BlockList) TableName() string {
    return "block_list"
}

func (e *BlockList) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BlockList) GetId() interface{} {
	return e.ID
}
