package model

import "time"

type PostFavour struct {
	ID         uint64    `json:"id"`
	PostId     uint64    `json:"post_id"`
	UserId     uint64    `json:"user_id"`
	CreateTime time.Time `json:"create_time"`
	UpDateTime time.Time `json:"update_time"`
}

func (p PostFavour) TableName() string {
	return "post_favour"
}
