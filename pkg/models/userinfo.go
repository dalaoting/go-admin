package models

import (
	"gorm.io/gorm"
	"time"

	"go-admin/common/models"
)

type UserInfo struct {
	gorm.Model
	models.ControlBy

	Identity       string    `json:"identity" gorm:"type:varchar(64);comment:用户id"`     //
	Account        string    `json:"account" gorm:"type:varchar(32);comment:账号"`        //
	Password       string    `json:"password" gorm:"type:varchar(64);comment:密码"`       //
	NickName       string    `json:"nickName" gorm:"type:varchar(20);comment:昵称"`       //
	Avatar         string    `json:"avatar" gorm:"type:varchar(200);comment:头像"`        //
	Profession     string    `json:"profession" gorm:"type:varchar(30);comment:职业"`     //
	Sex            string    `json:"sex" gorm:"type:int;comment:性别"`                    //
	Remark         string    `json:"remark" gorm:"type:varchar(100);comment:备注"`        //
	Status         string    `json:"status" gorm:"type:int;comment:状态"`                 //
	ForbidDeadline time.Time `json:"forbidDeadline" gorm:"type:timestamp;comment:禁用时间"` //
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp;comment:创建时间"`     //
	UpdateTime     time.Time `json:"updateTime" gorm:"type:timestamp;comment:更新时间"`     //
}

func (UserInfo) TableName() string {
	return "user_info"
}

func (e *UserInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserInfo) GetId() interface{} {
	return e.ID
}
