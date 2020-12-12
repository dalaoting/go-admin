package models

import "time"

type AssignIssue struct {
	ID           uint   `gorm:"primarykey"`
	Identity     string `gorm:"column:identity" json:"identity"`          // 用户id
	AssignSerial string `gorm:"column:assign_serial" json:"assignSerial"` // 任务编号
	DeptId       int    `gorm:"column:dept_id" json:"deptId"`             // 商家ID
	Content      string `gorm:"column:content" json:"content"`            // 内容
	ContentType  int    `gorm:"column:content_type" json:"contentType"`   // 内容类型，1-文本 2-图片 3-视频
	SendType     int    `gorm:"column:send_type" json:"sendType"`         // 消息发送方，1-用户端 2-商家端 3-系统
	UserRead     int    `gorm:"column:user_read" json:"userRead"`         // 用户已读状态，0-未读 1-已读
	DeptRead     int    `gorm:"column:dept_read" json:"deptRead"`         // 商家已读状态，0-未读 1-已读

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*AssignIssue) TableName() string {
	return "assign_issue"
}
