package v1

import "github.com/gin-gonic/gin"

type UserAnswer struct {
}

func NewUserAnswer() UserAnswer {
	return UserAnswer{}
}

func (ua UserAnswer) AddUserAnswer(c *gin.Context) {

}

func (ua UserAnswer) DeleteUserAnswer(c *gin.Context) {

}

func (ua UserAnswer) UpdateUserAnswer(c *gin.Context) {

}

func (ua UserAnswer) GetUserAnswerInfo(c *gin.Context) {

}

func (ua UserAnswer) ListUserAnswerByPage(c *gin.Context) {

}

func (ua UserAnswer) ListUserAnswerInfoByPage(c *gin.Context) {

}

func (ua UserAnswer) ListMyUserAnswerInfoByPage(c *gin.Context) {

}

func (ua UserAnswer) EditUserAnswer(c *gin.Context) {

}
