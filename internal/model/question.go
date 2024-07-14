package model

import "goAiproject/pkg/app"

type Question struct {
	*Model
	QuestionContent string `json:"question_content"`
	AppId           string `json:"app_id"`
	UserId          string `json:"user_id"`
}

type QuestionSwagger struct {
	List  []*Question
	Pager *app.Pager
}

func (q Question) TableName() string {
	return "question"
}
