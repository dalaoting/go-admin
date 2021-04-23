package models

type SysUser struct {
	ControlBy
	ModelTime
	UserId   int      `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string   `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Password string   `json:"-" gorm:"type:varchar(128);comment:密码"`
	NickName string   `json:"nickName" gorm:"type:varchar(128);comment:昵称"`
	Phone    string   `json:"phone" gorm:"type:varchar(11);comment:手机号"`
	RoleId   int      `json:"roleId" gorm:"type:bigint(20);comment:角色ID"`
	Salt     string   `json:"-" gorm:"type:varchar(255);comment:加盐"`
	Avatar   string   `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Sex      string   `json:"sex" gorm:"type:varchar(255);comment:性别"`
	Email    string   `json:"email" gorm:"type:varchar(128);comment:邮箱"`
	DeptId   int      `json:"deptId" gorm:"type:bigint(20);comment:部门"`
	PostId   int      `json:"postId" gorm:"type:bigint(20);comment:岗位"`
	Remark   string   `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Status   string   `json:"status" gorm:"type:varchar(4);comment:状态"`
	DeptIds  []int    `json:"deptIds" gorm:"-"`
	PostIds  []int    `json:"postIds" gorm:"-"`
	RoleIds  []int    `json:"roleIds" gorm:"-"`
	Dept     *SysDept `json:"dept"`
}

func (SysUser) TableName() string {
	return "sys_user"
}