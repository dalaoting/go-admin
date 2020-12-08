package commentdemandassign

import "github.com/gin-gonic/gin"

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

}
