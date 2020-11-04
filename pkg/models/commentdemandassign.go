package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type CommentDemandAssign struct {
    gorm.Model
    models.ControlBy
    
    DemandSerial string `json:"demandSerial" gorm:"type:varchar(64);comment:需求编号"` // 
    DemandSnapshotCode string `json:"demandSnapshotCode" gorm:"type:varchar(32);comment:快照编号"` // 
    DeptId string `json:"deptId" gorm:"type:bigint;comment:任务所属企业ID"` // 
    Commentator string `json:"commentator" gorm:"type:varchar(16);comment:测评人"` // 
    CommentAccount string `json:"commentAccount" gorm:"type:varchar(32);comment:测评账号"` // 
    CommentName string `json:"commentName" gorm:"type:varchar(32);comment:测评账号名称"` // 
    Status string `json:"status" gorm:"type:int;comment:状态"` // 
    TipsStatus string `json:"tipsStatus" gorm:"type:int;comment:提示"` // 
    IncomeAccouunt string `json:"incomeAccouunt" gorm:"type:varchar(64);comment:收款账号"` // 
    IncomeType string `json:"incomeType" gorm:"type:int;comment:收款账号类型"` // 
    IncomeName string `json:"incomeName" gorm:"type:varchar(64);comment:收款名称"` // 
    OrderMedias string `json:"orderMedias" gorm:"type:varchar(32);comment:下单截图"` // 
    CommentMedias string `json:"commentMedias" gorm:"type:varchar(32);comment:留评截图"` // 
    SettleMedias string `json:"settleMedias" gorm:"type:varchar(32);comment:结算截图"` // 
}

func (CommentDemandAssign) TableName() string {
    return "comment_demand_assign"
}

func (e *CommentDemandAssign) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CommentDemandAssign) GetId() interface{} {
	return e.ID
}
