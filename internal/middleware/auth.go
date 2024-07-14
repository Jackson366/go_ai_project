package middleware

import (
	"github.com/gin-gonic/gin"
	"goAiproject/global"
	"goAiproject/internal/service"
	"goAiproject/pkg/app"
	"goAiproject/pkg/errcode"
)

func Auth(mustRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUserId, exists := c.Get("userId")
		response := app.NewResponse(c)
		if !exists {
			response.ToErrorResponse(errcode.ErrorUserNotLogin)
			c.Abort()
			return
		}
		userId := currentUserId.(uint64)
		svc := service.New(c)
		// 每次都要对数据库进行查询，以确保用户的角色信息是最新的，可否用redis代替，减少数据库查询次数
		user, err := svc.GetUserById(userId)
		if err != nil {
			global.Logger.Fatalf("svc.GetLoginUser err: %v", err)
			response.ToErrorResponse(errcode.ErrorUserNotExist)
			c.Abort()
			return
		}
		if user.UserRole != mustRole {
			response.ToErrorResponse(errcode.ErrorNoPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}
