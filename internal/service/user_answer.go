package service

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
	"time"
)

type UserAnswerInfo struct {
	ID              uint64
	AppId           uint64
	AppType         uint8
	ScoringStrategy uint8
	Choices         []*string
	ResultId        uint64
	ResultName      string
	ResultDesc      string
	ResultPicture   string
	ResultScore     uint8
	UserId          uint64
	CreateTime      time.Time
	UpdateTime      time.Time
	User            *UserInfo
}

func (svc *Service) ValidUserAnswer(answer *model.UserAnswer, add bool) error {
	return nil
}

func (svc *Service) GetUserAnswerInfo(answer *model.UserAnswer) (*UserAnswerInfo, error) {
	return nil, nil
}

func (svc *Service) GetUserAnswerInfoPage(answer *model.UserAnswer, pager *app.Pager) ([]*UserAnswerInfo, error) {
	return nil, nil
}
