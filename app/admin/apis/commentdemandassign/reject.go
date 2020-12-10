package commentdemandassign

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-admin/common/actions"
	"go-admin/common/log"
	"go-admin/constant"
	"go-admin/pkg/models"
	"go-admin/tools"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RejectRequest struct {
	AssignSerial string `json:"assignSerial" binding:"required"`
	Remark       string `json:"remark" binding:"required"`
}

func (e *CommentDemandAssign) Reject(c *gin.Context) {
	req := &RejectRequest{}
	if err := c.BindJSON(req); err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数错误")
		return
	}

	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}
	tx := db.Begin().Set("gorm:query_option", "FOR UPDATE")

	msgID := tools.GenerateMsgIDFromContext(c)
	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	var data models.CommentDemandAssign

	err = tx.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where("assign_serial=?", req.AssignSerial).First(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "查看对象不存在或无权查看")
		return
	}
	if db.Error != nil {
		tx.Rollback()
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	status, _ := strconv.Atoi(data.Status)
	if status != constant.DemandAssignComment && status != constant.DemandAssignOrder {
		tx.Rollback()
		e.Error(c, http.StatusUnprocessableEntity, err, "当前状态不可驳回")
		return
	}

	if status == constant.DemandAssignOrder {
		data.Status = strconv.Itoa(constant.DemandAssignProcessing)
	} else if status == constant.DemandAssignComment {
		data.Status = strconv.Itoa(constant.DemandAssignOrder)
	}
	data.TipsStatus = "1"
	if err := tx.Save(data).Error; err != nil {
		tx.Rollback()
		e.Error(c, http.StatusUnprocessableEntity, err, "驳回失败")
		return
	}

	// TODO 添加一句反馈消息

}
