package models

import (
	"time"

	"go-admin/common/models"
)

type CommentAccount struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	models.ControlBy

	Identity    string    `json:"identity" gorm:"type:varchar(32);comment:用户id"`   //
	Type        string    `json:"type" gorm:"type:int;comment:类型"`                 //
	Account     string    `json:"account" gorm:"type:varchar(128);comment:账号"`     //
	AccountName string    `json:"accountName" gorm:"type:varchar(64);comment:用户名"` //
	Status      string    `json:"status" gorm:"type:int;comment:状态"`               //
	IsDelete    string    `json:"isDelete" gorm:"type:int;comment:已删除"`            //
	CreateTime  time.Time `json:"createTime" gorm:"type:timestamp;comment:创建时间"`   //
	UpdateTime  time.Time `json:"updateTime" gorm:"type:timestamp;comment:更新时间"`   //
}

func (CommentAccount) TableName() string {
	return "comment_account"
}

func (e *CommentAccount) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CommentAccount) GetId() interface{} {
	return e.ID
}
