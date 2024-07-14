package model

type UserAnswer struct {
	*Model
	AppId           string `json:"app_id"`
	AppType         uint8  `json:"app_type"`
	ScoringStrategy uint8  `json:"scoring_strategy"`
	Choices         string `json:"choices"`
	ResultId        uint64 `json:"result_id"`
	ResultName      string `json:"result_name"`
	ResultDesc      string `json:"result_desc"`
	ResultPicture   string `json:"result_picture"`
	ResultScore     uint32 `json:"result_score"`
	UserId          uint64 `json:"user_id"`
}

func (u UserAnswer) TableName() string {
	return "user_answer"
}
