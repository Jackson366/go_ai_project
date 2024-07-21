package v1

import "github.com/gin-gonic/gin"

type PostThumb struct {
}

func NewPostThumb() PostThumb {
	return PostThumb{}
}

func (pt PostThumb) DoPostThumb(c *gin.Context) {

}
