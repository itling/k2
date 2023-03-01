package router

import (
	"admin/apis"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerImRouter)
}

// 需认证的路由代码
func registerImRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.ImUserFriend{}
	r := v1.Group("/im").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/friend-req-list", api.GetFriendRequestPage)
		r.GET("/friend-list", api.GetFriendPage)
		r.POST("/friend-req", api.CreateFriendReq)
		r.PUT("/accept-friend-req", api.AcceptFriendReq)
		r.PUT("/reject-friend-req", api.RejectFriendReq)
		r.PUT("/update-friend-remark", api.UpdateFriendRemark)
		r.DELETE("/delete-friend", api.RemoveUserFriend)
		r.POST("/block-friend", api.BlockUserFriend)
		r.DELETE("/cancel-block-friend", api.CancelBlockUserFriend)
	}
}
