package service

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
	"time"
)

type AppInfo struct {
	ID              uint64
	AppName         string
	AppDesc         string
	AppIcon         string
	AppType         uint8
	ScoringStrategy uint8
	ReviewStatus    uint8
	ReviewMessage   string
	ReviewerId      uint64
	ReviewTime      time.Time
	UserId          uint64
	CreateTime      time.Time
	UpdateTime      time.Time
	UserInfo        *UserInfo
}

func (svc *Service) GetAppInfo(app *model.App) (*AppInfo, error) {
	return nil, nil
}

func (svc *Service) GetAppInfoPage(app *model.App, pager *app.Pager) ([]*model.App, error) {
	return nil, nil
}
