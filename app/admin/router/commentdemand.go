package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/commentdemand"
	"go-admin/pkg/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCommentDemandRouter)
}

// 需认证的路由代码
func registerCommentDemandRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &commentdemand.CommentDemand{}
	r := v1.Group("/commentdemand").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetCommentDemandList)
		r.GET("/:id", api.GetCommentDemand)
		r.POST("", api.InsertCommentDemand)
		r.PUT("/:id", api.UpdateCommentDemand)
		r.DELETE("", api.DeleteCommentDemand)
	}
}
