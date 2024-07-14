package service

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
	"time"
)

type ScoringResultInfo struct {
	ID               uint64
	ResultName       string
	ResultDesc       string
	ResultPicture    string
	ResultProp       []*string
	ResultScoreRange uint32
	AppId            uint64
	UserId           uint64
	CreateTime       time.Time
	UpdateTime       time.Time
	User             *UserInfo
}

func (svc *Service) ValidScoringResult(result *model.ScoringResult, add bool) error {
	return nil
}

func (svc *Service) GetScoringResultInfo(result *model.ScoringResult) (*ScoringResultInfo, error) {
	return nil, nil
}

func (svc *Service) GetScoringResultInfoPage(result *ScoringResultInfo, pager *app.Pager) ([]*ScoringResultInfo, error) {
	return nil, nil
}
