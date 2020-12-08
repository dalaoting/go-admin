package commentdemandassign

import "github.com/gin-gonic/gin"

type GetAssignIssueListRequest struct {
	AssignSerial string `json:"assign_serial"`
}

func (e *CommentDemandAssign) GetAssignIssueList(c *gin.Context) {

}
