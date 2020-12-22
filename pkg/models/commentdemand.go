package models

import (
	"gorm.io/gorm"

	"go-admin/common/models"
)

type CommentDemand struct {
	gorm.Model
	models.ControlBy

	DeptId       string `json:"deptId" gorm:"type:bigint;comment:企业"`            //
	SerialNumber string `json:"serialNumber" gorm:"type:varchar(64);comment:编号"` //
	CustomerId   string `json:"customerId" gorm:"type:bigint;comment:客户"`        //
	Title        string `json:"title" gorm:"type:varchar(128);comment:标题"`       //
	ShopName     string `json:"shopName" gorm:"type:varchar(128);comment:店铺名"`   //
	ProductCode  string `json:"productCode" gorm:"type:varchar(32);comment:商品码"` //
	Reward       string `json:"reward" gorm:"type:int;comment:佣金(分)"`            //
	ProductPrice string `json:"productPrice" gorm:"type:int;comment:价格"`         //
	DemandPrice  string `json:"demandPrice" gorm:"type:int;comment:需求总费用"`       //
	CommentNum   string `json:"commentNum" gorm:"type:int;comment:测评数"`          //
	Desc         string `json:"desc" gorm:"type:varchar(255);comment:说明"`        //
	Remark       string `json:"remark" gorm:"type:varchar(255);comment:备注"`      //
	Status       string `json:"status" gorm:"type:int;comment:状态"`               //

	CreateTime string `json:"createdAt" gorm:"-"`
	CreateUser string `json:"createBy" gorm:"-"`
	UpdateUser string `json:"updateBy" gorm:"-"`
}

func (CommentDemand) TableName() string {
	return "comment_demand"
}

func (e *CommentDemand) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CommentDemand) GetId() interface{} {
	return e.ID
}
