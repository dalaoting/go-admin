package customer

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	common "go-admin/common/models"
	"go-admin/pkg/constant"
	"net/http"
	"strconv"

	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/pkg/models"
	"go-admin/tools"
)

type Customer struct {
	apis.Api
}

func (e *Customer) GetCustomerList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.CustomerSearch)
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
	// 切换企业ID
	de := &models.SysDept{DeptId: p.DeptId}
	dept, _ := de.Get()
	if dept.ParentId > 0 {
		p.DeptId = dept.ParentId
	}
	list := make([]models.Customer, 0)
	var count int64
	serviceStudent := service.Customer{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetCustomerPage(req, p, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	var (
		createUserId = make([]int, 0)
	)
	for i := range list {
		createUserId = append(createUserId, int(list[i].CreateBy), int(list[i].UpdateBy))
		list[i].DeptId = dept.DeptName
		list[i].CreateTime = list[i].CreatedAt.Format(constant.DefaultTimeFormat)
	}
	user := &models.SysUser{}
	userMap, _ := user.BatchGet(createUserId)
	for i := range list {
		if v, ok := userMap[int(list[i].CreateBy)]; ok {
			list[i].CreateUser = v.NickName
		} else {
			list[i].CreateUser = "用户"
		}

		if v, ok := userMap[int(list[i].UpdateBy)]; ok {
			list[i].UpdateUser = v.NickName
		} else {
			list[i].UpdateUser = ""
		}
	}
	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e *Customer) GetCustomer(c *gin.Context) {
	control := new(dto.CustomerById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//查看详情
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.Customer

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCustomer := service.Customer{}
	serviceCustomer.MsgID = msgID
	serviceCustomer.Orm = db
	err = serviceCustomer.GetCustomer(req, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}

func (e *Customer) InsertCustomer(c *gin.Context) {
	control := new(dto.CustomerControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	control.DeptId = strconv.Itoa(p.DeptId)
	// 切换企业ID
	de := &models.SysDept{DeptId: p.DeptId}
	dept, _ := de.Get()
	if dept.ParentId > 0 {
		control.DeptId = strconv.Itoa(dept.ParentId)
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

	serviceCustomer := service.Customer{}
	serviceCustomer.Orm = db
	serviceCustomer.MsgID = msgID
	err = serviceCustomer.InsertCustomer(object)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

func (e *Customer) UpdateCustomer(c *gin.Context) {
	control := new(dto.CustomerControl)
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

	serviceCustomer := service.Customer{}
	serviceCustomer.Orm = db
	serviceCustomer.MsgID = msgID
	err = serviceCustomer.UpdateCustomer(object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

func (e *Customer) DeleteCustomer(c *gin.Context) {
	control := new(dto.CustomerById)
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

	serviceCustomer := service.Customer{}
	serviceCustomer.Orm = db
	serviceCustomer.MsgID = msgID
	err = serviceCustomer.RemoveCustomer(req, object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}
