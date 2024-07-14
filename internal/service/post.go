package service

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
)

type PostInfo struct {
}

func (svc *Service) ValidPost(post *model.Post) error {
	return nil
}

func (svc *Service) GetPostInfo(post *model.Post) (*PostInfo, error) {
	return nil, nil
}

func (svc *Service) GetPostInfoPage(post *model.Post, pager *app.Pager) ([]*model.Post, error) {
	return nil, nil
}
