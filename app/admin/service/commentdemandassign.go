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

type CommentDemandAssign struct {
	service.Service
}

// GetCommentDemandAssignPage 获取CommentDemandAssign列表
func (e *CommentDemandAssign) GetCommentDemandAssignPage(c cDto.Index, p *actions.DataPermission, list *[]models.CommentDemandAssign, count *int64) error {
	var err error
	var data models.CommentDemandAssign
	msgID := e.MsgID

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			//actions.PermissionDeptId(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// GetCommentDemandAssign 获取CommentDemandAssign对象
func (e *CommentDemandAssign) GetCommentDemandAssign(serialNumber string, p *actions.DataPermission, model *models.CommentDemandAssign) error {
	var err error
	var data models.CommentDemandAssign
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Where("assign_serial=?", serialNumber).
		First(model)
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

// InsertCommentDemandAssign 创建CommentDemandAssign对象
func (e *CommentDemandAssign) InsertCommentDemandAssign(model common.ActiveRecord) error {
	var err error
	var data models.CommentDemandAssign
	msgID := e.MsgID

	err = e.Orm.Model(&data).
		Create(model).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// UpdateCommentDemandAssign 修改CommentDemandAssign对象
func (e *CommentDemandAssign) UpdateCommentDemandAssign(c common.ActiveRecord, p *actions.DataPermission) error {
	var err error
	var data models.CommentDemandAssign
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

// RemoveCommentDemandAssign 删除CommentDemandAssign
func (e *CommentDemandAssign) RemoveCommentDemandAssign(d cDto.Control, c common.ActiveRecord, p *actions.DataPermission) error {
	var err error
	var data models.CommentDemandAssign
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
