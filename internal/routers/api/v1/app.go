package v1

import "github.com/gin-gonic/gin"

type App struct {
}

func NewApp() App {
	return App{}
}

func (a App) AddApp(c *gin.Context) {

}

func (a App) DeleteApp(c *gin.Context) {

}

func (a App) UpdateApp(c *gin.Context) {

}

func (a App) GetAppInfo(c *gin.Context) {

}

func (a App) ListAppByPage(c *gin.Context) {

}

func (a App) ListAppInfoByPage(c *gin.Context) {

}

func (a App) ListMyAppInfoByPage(c *gin.Context) {

}

func (a App) EditApp(c *gin.Context) {

}

func (a App) ReviewApp(c *gin.Context) {

}
