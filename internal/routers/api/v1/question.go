package v1

import "github.com/gin-gonic/gin"

type Question struct {
}

func NewQuestion() Question {
	return Question{}
}

func (q Question) AddQuestion(c *gin.Context) {

}

func (q Question) DeleteQuestion(c *gin.Context) {

}

func (q Question) UpdateQuestion(c *gin.Context) {

}

func (q Question) GetQuestionInfo(c *gin.Context) {

}

func (q Question) ListQuestionByPage(c *gin.Context) {

}

func (q Question) ListQuestionInfoByPage(c *gin.Context) {

}

func (q Question) ListMyQuestionInfoByPage(c *gin.Context) {

}

func (q Question) EditQuestion(c *gin.Context) {

}

func (q Question) AiGenerateQuestion(c *gin.Context) {

}

func (q Question) AiGenerateQuestionSSE(c *gin.Context) {

}
