package models

// MediaType 媒体类型记录表
type MediaType struct {
	ID         int64  `gorm:"primary_key;column:id;AUTO_INCREMENT" json:"-"` // 主键id
	Status     int    `gorm:"column:status" json:"status"`                   // 状态， 1: 有效 -1: 无效
	CreateTime string `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;default:null" json:"update_time"`
	Comment    string `gorm:"column:comment" json:"comment"` // 说明
	Suffix     string `gorm:"column:suffix" json:"suffix"`   // 文件类型后缀名
	Folder     string `gorm:"column:folder" json:"folder"`   // 文件类型所属的文件夹
}

// TableName get sql table name.获取数据库表名
func (m *MediaType) TableName() string {
	return "media_type"
}
