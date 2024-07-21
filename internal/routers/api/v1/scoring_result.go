package v1

import "github.com/gin-gonic/gin"

type ScoringResult struct {
}

func NewScoringResult() ScoringResult {
	return ScoringResult{}
}

func (sr ScoringResult) AddScoringResult(c *gin.Context) {

}

func (sr ScoringResult) DeleteScoringResult(c *gin.Context) {

}

func (sr ScoringResult) UpdateScoringResult(c *gin.Context) {

}

func (sr ScoringResult) GetScoringResultInfo(c *gin.Context) {

}

func (sr ScoringResult) ListScoringResultByPage(c *gin.Context) {

}

func (sr ScoringResult) ListScoringResultInfoByPage(c *gin.Context) {

}

func (sr ScoringResult) ListMyScoringResultInfoByPage(c *gin.Context) {

}

func (sr ScoringResult) EditScoringResult(c *gin.Context) {

}
