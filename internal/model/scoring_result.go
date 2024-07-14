package model

import "goAiproject/pkg/app"

type ScoringResult struct {
	*Model
	ResultName       string `json:"result_name"`
	ResultDesc       string `json:"result_desc"`
	ResultPicture    string `json:"result_picture"`
	ResultProp       string `json:"result_prop"`
	ResultScoreRange uint32 `json:"result_score_range"`
	AppId            uint64 `json:"app_id"`
	UserId           uint64 `json:"user_id"`
}

type ScoringResultSwagger struct {
	List  []*ScoringResult
	Pager *app.Pager
}

func (s ScoringResult) TableName() string {
	return "scoring_result"
}
