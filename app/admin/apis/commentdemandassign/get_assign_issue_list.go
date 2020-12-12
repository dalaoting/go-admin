package commentdemandassign

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/actions"
	"go-admin/common/log"
	"go-admin/pkg/models"
	"go-admin/tools"
	"net/http"
)

type GetAssignIssueListRequest struct {
	AssignSerial string `json:"assign_serial"`
}

func (e *CommentDemandAssign) GetAssignIssueList(c *gin.Context) {
	req := &GetAssignIssueListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
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

	var (
		list  = make([]*models.AssignIssue, 0)
		count = int64(0)
	)

	if err := db.Debug().Where("assign_serial=? AND dept_id=?", req.AssignSerial, p.DeptId).
		Order("created_at DESC").Find(&list).Count(&count).Error; err != nil {
		log.Error(err)
	}

	e.PageOK(c, list, int(count), 0, 0, "查询成功")
}
