package commentdemandassign

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	common "go-admin/common/models"
	"go-admin/constant"
	"go-admin/pkg/service/mediaService"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"

	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/pkg/models"
	"go-admin/tools"
)

type CommentDemandAssign struct {
	apis.Api
}

func (e *CommentDemandAssign) GetCommentDemandAssignList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.CommentDemandAssignSearch)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	req := d.Generate()

	//查询列表
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.CommentDemandAssign, 0)
	var count int64
	serviceStudent := service.CommentDemandAssign{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetCommentDemandAssignPage(req, p, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	for i := range list {
		if list[i].OrderMedias != "" {
			list[i].OrderMedias, err = getMediaUrlByString(db, list[i].OrderMedias)
		}

		if list[i].CommentMedias != "" {
			list[i].CommentMedias, err = getMediaUrlByString(db, list[i].CommentMedias)
		}

		if list[i].SettleMedias != "" {
			list[i].SettleMedias, err = getMediaUrlByString(db, list[i].SettleMedias)
		}
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e *CommentDemandAssign) GetCommentDemandAssign(c *gin.Context) {
	type R struct {
		SerialNumber string `uri:"serial"`
	}
	control := new(R)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//查看详情
	err = c.ShouldBindUri(control)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.CommentDemandAssign

	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.MsgID = msgID
	serviceCommentDemandAssign.Orm = db
	err = serviceCommentDemandAssign.GetCommentDemandAssign(control.SerialNumber, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	if object.OrderMedias != "" {
		object.OrderMedias, err = getMediaUrlByString(db, object.OrderMedias)
	}

	if object.CommentMedias != "" {
		object.CommentMedias, err = getMediaUrlByString(db, object.CommentMedias)
	}

	if object.SettleMedias != "" {
		object.SettleMedias, err = getMediaUrlByString(db, object.SettleMedias)
	}

	var (
		result   = make(map[string]interface{})
		snapshot = &models.CommentDemandSnapshot{}
	)

	db.Where("serial_number=?", object.DemandSnapshotCode).First(snapshot)
	buf, _ := json.Marshal(object)
	_ = json.Unmarshal(buf, &result)
	result["snapshot"] = snapshot
	e.OK(c, result, "查看成功")
}

func (e *CommentDemandAssign) InsertCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//新增操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(tools.GetUserIdUint(c))

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.Orm = db
	serviceCommentDemandAssign.MsgID = msgID
	err = serviceCommentDemandAssign.InsertCommentDemandAssign(object)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

func (e *CommentDemandAssign) UpdateCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	req := control.Generate()
	//更新操作
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(tools.GetUserIdUint(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.Orm = db
	serviceCommentDemandAssign.MsgID = msgID
	err = serviceCommentDemandAssign.UpdateCommentDemandAssign(object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

func (e *CommentDemandAssign) DeleteCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//删除操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		log.Errorf("MsgID[%s] Bind error: %s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}

	// 设置编辑人
	object.SetUpdateBy(tools.GetUserIdUint(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.Orm = db
	serviceCommentDemandAssign.MsgID = msgID
	err = serviceCommentDemandAssign.RemoveCommentDemandAssign(req, object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}

func getMediaUrlByString(db *gorm.DB, mediaIdsStr string) (string, error) {
	mediaIds := make([]int64, 0)
	if err := json.Unmarshal([]byte(mediaIdsStr), &mediaIds); err != nil {
		return "", err
	}
	urlArr, err := mediaService.GetMediaUrlArr(db, mediaIds)
	if err != nil {
		return "", err
	}
	return strings.Join(urlArr, ","), nil
}

type UpdateStatusRequest struct {
	// 任务编号
	// required: true
	AssignSerial string `json:"assign_serial" binding:"required"`
	// 媒体Id列表
	MediaIds []int64 `json:"media_ids" binding:"required"`
}

func (e *CommentDemandAssign) UpdateStatus(c *gin.Context) {
	req := &UpdateStatusRequest{}
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

	if err := tx.Where("assign_serial=?", req.AssignSerial).First(record).Error; err != nil {
		tx.Rollback()
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	status, _ := strconv.Atoi(record.Status)
	if status != constant.DemandAssignComment && status != constant.DemandAssignOrderSettle {
		tx.Rollback()
		e.Error(c, http.StatusUnprocessableEntity, err, "当前状态不可更新")
		return
	}

	if status == constant.DemandAssignComment {
		record.Status = strconv.Itoa(constant.DemandAssignOrderSettle)
		mediaIdsBuf, _ := json.Marshal(req.MediaIds)

		record.SettleMedias = string(mediaIdsBuf)
	} else {
		record.Status = strconv.Itoa(constant.DemandAssignRewardSettle)
		mediaIds := make([]int64, 0)
		_ = json.Unmarshal([]byte(record.SettleMedias), &mediaIds)
		mediaIds = append(mediaIds, req.MediaIds...)
		buf, _ := json.Marshal(mediaIds)
		record.SettleMedias = string(buf)
	}
	// 开始更新
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		e.Error(c, http.StatusUnprocessableEntity, err, "更新失败")
		return
	}
	tx.Commit()
	e.OK(c, record, "更新成功")
}
