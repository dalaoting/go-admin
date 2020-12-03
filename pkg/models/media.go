package models

// Media 媒体表
type Media struct {
	ID            int64  `gorm:"primary_key;column:id;AUTO_INCREMENT" json:"-"` // 主键id
	MediaType     int    `gorm:"column:media_type" json:"media_type"`           // 媒体类型
	OriginalURL   string `gorm:"column:original_url" json:"original_url"`       // 媒体源地址
	OriginalMd5   string `gorm:"column:original_md5" json:"original_md5"`       // 源文件MD5
	NewMd5        string `gorm:"column:new_md5" json:"new_md5"`                 // 新文件MD5
	Status        int    `gorm:"column:status" json:"status"`                   // 状态 1: 通过(有效) -1: 无效 0: 待审核 2: 审核中
	Width         int    `gorm:"column:width" json:"width"`                     // 文件长度
	Height        int    `gorm:"column:height" json:"height"`                   // 文件宽度
	MediaTime     int    `gorm:"column:media_time" json:"media_time"`           // 媒体时长，单位: 毫秒
	MediaSize     int    `gorm:"column:media_size" json:"media_size"`           // 媒体大小
	MarkStatus    int    `gorm:"column:mark_status" json:"mark_status"`         // 水印标记, 1: 有印记 0: 无印记
	UnmarkedURL   string `gorm:"column:unmarked_url" json:"unmarked_url"`       // 净化后的文件地址
	CodeRate      int    `gorm:"column:code_rate" json:"code_rate"`             // 媒体码率
	FirstFrameURL string `gorm:"column:first_frame_url" json:"first_frame_url"` // 视频首帧图
	MarkPath      string `gorm:"column:mark_path" json:"mark_path"`             // 水印图片路径
	IP            string `gorm:"column:ip" json:"ip"`                           // ip地址
	CreateTime    string `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime    string `gorm:"column:update_time;default:null" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *Media) TableName() string {
	return "media"
}
