package model

import (
	"goAiproject/pkg/app"
	"time"
)

type App struct {
	*Model
	AppName         string    `gorm:"column:appName" json:"app_name"`
	AppDesc         string    `gorm:"column:appDesc" json:"app_desc"`
	AppIcon         string    `gorm:"column:appIcon" json:"app_icon"`
	AppType         uint8     `gorm:"column:appType" json:"app_type"`
	ScoringStrategy uint8     `gorm:"column:scoringStrategy" json:"scoring_strategy"`
	ReviewStatus    uint32    `gorm:"column:reviewStatus" json:"review_status"`
	ReviewMessage   string    `gorm:"column:reviewMessage" json:"review_message"`
	ReviewerId      uint64    `gorm:"column:reviewId" json:"reviewer_id"`
	ReviewTime      time.Time `gorm:"column:reviewTime" json:"review_time"`
	UserId          uint64    `gorm:"column:userId" json:"user_id"`
}

type AppSwagger struct {
	List  []*App
	Pager *app.Pager
}

func (a App) TableName() string {
	return "app"
}
