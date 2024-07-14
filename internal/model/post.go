package model

import "goAiproject/pkg/app"

type Post struct {
	*Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tags      string `json:"tags"`
	ThumbNum  uint32 `json:"thumb_num"`
	FavourNum uint32 `json:"comment_num"`
	UserId    uint64 `json:"user_id"`
}

type PostSwagger struct {
	List  []*Post
	Pager *app.Pager
}

func (p Post) TableName() string {
	return "post"
}
