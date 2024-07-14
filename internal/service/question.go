package service

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
	"time"
)

type QuestionInfo struct {
	ID              uint64
	QuestionContent []*QuestionContent
	AppId           uint64
	UserId          uint64
	CreateTime      time.Time
	UpdateTime      time.Time
	User            *UserInfo
}

type QuestionContent struct {
	Title   string
	Options []*Option
}

type Option struct {
	Result string
	score  uint32
	value  string
	key    string
}

func (svc *Service) ValidQuestion(question *model.Question) error {
	return nil
}

func (svc *Service) GetQuestionInfo(question *model.Question) (*QuestionInfo, error) {
	return nil, nil
}

func (svc *Service) GetQuestionInfoPage(question model.Question, pager *app.Pager) ([]*QuestionInfo, error) {
	return nil, nil
}
