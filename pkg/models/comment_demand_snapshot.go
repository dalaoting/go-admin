package models

import (
	"gorm.io/gorm"
	"time"
)

type CommentDemandSnapshot struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ID           uint   `gorm:"primarykey"`
	DeptId       int    `json:"deptId" gorm:"type:bigint;comment:企业"`            //
	SerialNumber string `json:"serialNumber" gorm:"type:varchar(64);comment:编号"` //
	CustomerId   int    `json:"customerId" gorm:"type:bigint;comment:客户"`        //
	Type         int    `json:"type" gorm:"type:int;comment:类型，1-亚马逊"`
	Title        string `json:"title" gorm:"type:varchar(128);comment:标题"`       //
	ShopName     string `json:"shopName" gorm:"type:varchar(128);comment:店铺名"`   //
	ProductCode  string `json:"productCode" gorm:"type:varchar(32);comment:商品码"` //
	Reward       int    `json:"reward" gorm:"type:int;comment:佣金(分)"`            //
	ProductPrice int    `json:"productPrice" gorm:"type:int;comment:价格"`         //
	CommentNum   int    `json:"commentNum" gorm:"type:int;comment:测评数"`          //
	Desc         string `json:"desc" gorm:"type:varchar(255);comment:说明"`        //
	Remark       string `json:"remark" gorm:"type:varchar(255);comment:备注"`      //
	Status       string `json:"status" gorm:"type:int;comment:状态"`               //

	CreateBy uint `json:"createBy" gorm:"index;comment:'创建者'"`
	UpdateBy uint `json:"updateBy" gorm:"index;comment:'更新者'"`
}

func (CommentDemandSnapshot) TableName() string {
	return "comment_demand_snapshot"
}
