package models

import (
	"errors"
	orm "go-admin/common/global"
)

type Customer struct {
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

func (*Customer) TableName() string {
	return "customer"
}

func (e *Customer) GetPage() (customers []Customer, err error) {
	table := orm.Eloquent.Table(e.TableName())
	if e.DeptId <= 0 {
		return nil, errors.New("请先加入企业")
	}

	table = table.Where("dept_id = ?", e.DeptId)

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if err = table.Order("amount DESC").Find(&customers).Error; err != nil {
		return
	}
	return
}

//添加
func (e *Customer) Insert() (id int, err error) {
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

//修改
func (e *Customer) Update(id int) (update SysUser, err error) {

	if err = orm.Eloquent.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if e.DeptId == 0 {
		e.DeptId = update.RoleId
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

//修改余额
func (e *Customer) UpdateAmount(id int) (update Customer, err error) {
	if err = orm.Eloquent.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}
	tx := orm.Eloquent.Set("gorm:query_option", "FOR UPDATE")
	if err = tx.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if err = tx.Exec("UPDATE customer SET amount = amount+? WHERE id = ?", e.Amount, id).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
