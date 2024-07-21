package v1

import "github.com/gin-gonic/gin"

type Post struct {
}

func NewPost() Post {
	return Post{}
}

func (p Post) AddPost(c *gin.Context) {

}

func (p Post) DeletePost(c *gin.Context) {

}

func (p Post) UpdatePost(c *gin.Context) {

}

func (p Post) GetPostInfo(c *gin.Context) {

}

func (p Post) ListPostByPage(c *gin.Context) {

}

func (p Post) ListPostInfoByPage(c *gin.Context) {

}

func (p Post) ListMyPostInfoByPage(c *gin.Context) {

}

func (p Post) EditPost(c *gin.Context) {

}
