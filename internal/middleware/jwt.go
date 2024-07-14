package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goAiproject/pkg/app"
	"goAiproject/pkg/errcode"
	"goAiproject/pkg/redis"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token  string
			ecode  = errcode.Success
			claim  *app.Claims
			userId uint64
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}
		if token == "" {
			ecode = errcode.UnauthorizedTokenError
		} else {
			var err error
			claim, err = app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				userId = claim.UserId
				client := redis.NewRedisClient()
				userIdStr := fmt.Sprintf("%v", userId)
				val, _ := client.Get(context.Background(), userIdStr).Result()
				if val == token {
					response := app.NewResponse(c)
					response.ToErrorResponse(errcode.UnauthorizedTokenError)
					c.Abort()
					return
				}
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
