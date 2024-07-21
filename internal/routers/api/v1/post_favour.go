package v1

import "github.com/gin-gonic/gin"

type PostFavour struct {
}

func NewPostFavour() PostFavour {
	return PostFavour{}
}

func (pf PostFavour) DoPostFavour(c *gin.Context) {

}

func (pf PostFavour) ListMyFavourPostByPage(c *gin.Context) {

}

func (pf PostFavour) ListFavourPostByPage(c *gin.Context) {

}
