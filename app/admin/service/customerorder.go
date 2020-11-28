package service

import (
	"errors"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/common/service"
	"go-admin/pkg/models"
	"gorm.io/gorm"
)

type CustomerOrder struct {
	service.Service
}

// GetCustomerOrderPage 获取CustomerOrder列表
func (e *CustomerOrder) GetCustomerOrderPage(c cDto.Index, p *actions.DataPermission, list *[]models.CustomerOrder, count *int64) error {
	var err error
	var data models.CustomerOrder
	msgID := e.MsgID

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// GetCustomerOrder 获取CustomerOrder对象
func (e *CustomerOrder) GetCustomerOrder(d cDto.Control, p *actions.DataPermission, model *models.CustomerOrder) error {
	var err error
	var data models.CustomerOrder
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	if db.Error != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// InsertCustomerOrder 创建CustomerOrder对象
func (e *CustomerOrder) InsertCustomerOrder(model common.ActiveRecord) error {
	var err error
	var data models.CustomerOrder
	msgID := e.MsgID
	err = e.Orm.Model(&data).
		Create(model).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// UpdateCustomerOrder 修改CustomerOrder对象
func (e *CustomerOrder) UpdateCustomerOrder(c common.ActiveRecord, p *actions.DataPermission) error {
	var err error
	var data models.CustomerOrder
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if db.Error != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// RemoveCustomerOrder 删除CustomerOrder
func (e *CustomerOrder) RemoveCustomerOrder(d cDto.Control, c common.ActiveRecord, p *actions.DataPermission) error {
	var err error
	var data models.CustomerOrder
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where(d.GetId()).Delete(c)
	if db.Error != nil {
		err = db.Error
		log.Errorf("MsgID[%s] Delete error: %s", msgID, err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
