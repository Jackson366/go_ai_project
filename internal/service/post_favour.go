package service

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
)

func (svc *Service) DoPostFavour(postId uint64, loginUser *model.User) (int32, error) {
	return 0, nil
}

func (svc *Service) DoPostFavourInner(userId, postId uint64) (int32, error) {
	return 0, nil
}

func (svc *Service) ListFavourPostByPage(post *model.Post, pager *app.Pager, favourUserId uint64) ([]*model.Post, error) {
	return nil, nil
}
