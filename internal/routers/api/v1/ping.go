package v1

import "github.com/gin-gonic/gin"

type Ping struct{}

func NewPing() Ping {
	return Ping{}
}

func (p Ping) Ping(c *gin.Context) {
	userId, exists := c.Get("userId")
	if exists {
		c.JSON(200, gin.H{
			"userId": userId,
		})
		return
	}
	c.String(200, "pong")
	return
}
