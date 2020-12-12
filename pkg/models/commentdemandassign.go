package models

import (
	"gorm.io/gorm"

	"go-admin/common/models"
)

type CommentDemandAssign struct {
	gorm.Model

	Serial             string `json:"serial" gorm:"column:assign_serial;ype:varchar(64);comment:任务编号"` //
	DemandSerial       string `json:"demandSerial" gorm:"type:varchar(64);comment:需求编号"`               //
	DemandSnapshotCode string `json:"demandSnapshotCode" gorm:"type:varchar(32);comment:快照ID"`         //
	DeptId             string `json:"deptId" gorm:"type:bigint;comment:企业"`                            //
	Commentator        string `json:"commentator" gorm:"type:varchar(16);comment:测评人"`                 //
	CommentAccount     string `json:"commentAccount" gorm:"type:varchar(32);comment:测评账号"`             //
	CommentName        string `json:"commentName" gorm:"type:varchar(32);comment:测评账号名称"`              //
	Status             string `json:"status" gorm:"type:int;comment:状态"`                               //
	TipsStatus         string `json:"tipsStatus" gorm:"type:int;comment:提示"`                           //
	IncomeAccount      string `json:"incomeAccount" gorm:"type:varchar(64);comment:收款账号"`              //
	IncomeType         string `json:"incomeType" gorm:"type:int;comment:收款类型"`                         //
	IncomeName         string `json:"incomeName" gorm:"type:varchar(64);comment:收款名称"`                 //
	OrderMedias        string `json:"orderMedias" gorm:"type:varchar(32);comment:下单截图"`                //
	CommentMedias      string `json:"commentMedias" gorm:"type:varchar(32);comment:留评截图"`              //
	SettleMedias       string `json:"settleMedias" gorm:"type:varchar(32);comment:结算截图"`               //
}

func (CommentDemandAssign) TableName() string {
	return "comment_demand_assign"
}

func (e *CommentDemandAssign) SetCreateBy(createBy uint) {

}

func (e *CommentDemandAssign) SetUpdateBy(updateBy uint) {

}

func (e *CommentDemandAssign) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CommentDemandAssign) GetId() interface{} {
	return e.ID
}
