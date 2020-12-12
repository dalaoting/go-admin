package commentdemandassign

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/actions"
	"go-admin/common/log"
	"go-admin/constant"
	"go-admin/pkg/models"
	"go-admin/tools"
	"net/http"
	"strconv"
	"time"
)

type PushIssueRequest struct {
	// 任务编号
	// required: true
	AssignSerial string `json:"assign_serial" binding:"required"`
	// 内容类型，1-文本 2-图片 3-视频
	// required: true
	ContentType int `json:"content_type" binding:"required"`
	// 内容
	// required: true
	Content string `json:"content" binding:"required"`
}

func (e *CommentDemandAssign) PushIssue(c *gin.Context) {
	req := &PushIssueRequest{}
	if err := c.BindJSON(req); err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数错误")
		return
	}

	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusUnprocessableEntity, err, "服务异常")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	// 切换企业ID
	de := &models.SysDept{DeptId: p.DeptId}
	dept, _ := de.Get()
	if dept.ParentId > 0 {
		p.DeptId = dept.ParentId
	}

	// 先查询该记录是否还存在
	var (
		record = &models.CommentDemandAssign{}
	)

	if err := db.Debug().Where("assign_serial=?", req.AssignSerial).First(record).Error; err != nil {
		log.Error(err)
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	status, _ := strconv.Atoi(record.Status)
	if status == constant.DemandAssignCancel || status == constant.DemandAssignFinish || status == constant.DemandAssignRefund {
		e.Error(c, http.StatusUnprocessableEntity, err, "当前状态已不可以提交工单")
		return
	}

	issue := &models.AssignIssue{
		Identity:     record.Commentator,
		AssignSerial: req.AssignSerial,
		DeptId:       p.DeptId,
		Content:      req.Content,
		ContentType:  req.ContentType,
		SendType:     2,
		UserRead:     0,
		DeptRead:     1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := db.Debug().Save(issue).Error; err != nil {
		log.Error(err)
		e.Error(c, http.StatusUnprocessableEntity, err, "发送失败")
		return
	}

	e.OK(c, issue.ID, "发送成功")
}
