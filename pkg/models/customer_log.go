package models

import (
	"errors"
	orm "go-admin/common/global"
)

type CustomerLog struct {
	Id       int    `json:"id" gorm:"primary_key;auto_increment;"` // 客户ID
	DeptId   int    `json:"dept_id" gorm:"size:11"`                // 企业ID
	Name     string `json:"name" gorm:"size:64"`                   // 客户名称
	Amount   int    `json:"amount" gorm:"size:11"`                 // 余额,单位/人民币(分)
	Prepay   int    `json:"prepay" gorm:"size:11"`                 // 预付费,单位/人民币(分)
	Status   int    `json:"status" gorm:"size:1"`                  // 状态,
	CreateBy string `json:"createBy" gorm:"size:64;"`
	UpdateBy string `json:"updateBy" gorm:"size:64;"`
	BaseModel

	DataScope string `json:"dataScope" gorm:"-"`
}

func (*CustomerLog) TableName() string {
	return "customer_log"
}

//添加
func (e *CustomerLog) Insert() (id int, err error) {

	// check 用户名
	var count int64
	orm.Eloquent.Table(e.TableName()).Where("dept_id = ? AND name = ?", e.DeptId, e.Name).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = orm.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.Id
	return
}
