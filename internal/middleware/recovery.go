package middleware

import (
	"github.com/gin-gonic/gin"
	"goAiproject/global"
	"goAiproject/pkg/app"
	"goAiproject/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover: %v"
				global.Logger.WithCallersFrames().Fatalf(s, err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
