package models

import (
	"go-admin/common/models"
	"time"
)

type CommentAccount struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
	models.ControlBy

	Identity    string    `json:"identity" gorm:"type:varchar(32);comment:用户id"`   //
	Type        string    `json:"type" gorm:"type:int;comment:类型"`                 //
	Account     string    `json:"account" gorm:"type:varchar(128);comment:账号"`     //
	AccountName string    `json:"accountName" gorm:"type:varchar(64);comment:用户名"` //
	Status      string    `json:"status" gorm:"type:int;comment:状态"`               //
	IsDelete    string    `json:"isDelete" gorm:"type:int;comment:已删除"`            //
	CreateTime  time.Time `gorm:"-"`                                               //
	UpdateTime  time.Time `gorm:"-"`                                               //
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
