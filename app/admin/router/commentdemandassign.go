package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/commentdemandassign"
	"go-admin/common/actions"
	jwt "go-admin/pkg/jwtauth"
	"go-admin/pkg/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCommentDemandAssignRouter)
}

// 需认证的路由代码
func registerCommentDemandAssignRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &commentdemandassign.CommentDemandAssign{}
	r := v1.Group("/commentdemandassign").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.GET("", api.GetCommentDemandAssignList)
		r.GET("/:id", api.GetCommentDemandAssign)
		r.POST("", api.InsertCommentDemandAssign)
		r.PUT("/:id", api.UpdateCommentDemandAssign)
		r.DELETE("", api.DeleteCommentDemandAssign)
	}
}
