package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goAiproject/docs"
	"goAiproject/global"
	"goAiproject/internal/middleware"
	v1 "goAiproject/internal/routers/api/v1"
)

const (
	AdminRole = "admin"
	UserRole  = "user"
	BanRole   = "ban"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// 启动Debug模式
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	//url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ping := v1.NewPing()
	user := v1.NewUser()
	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	r.Use(middleware.JWT()).POST("/logout", user.Logout)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//apiv1.Use(middleware.Auth("admin")).GET("/ping", ping.Ping)
		//apiv1.Use(middleware.Auth("admin")).GET("/user/get/:id", user.GetLoginUser)
		apiv1.GET("/user/get/login", user.GetLoginUser)
		apiv1.GET("/user/get/info", user.GetUserInfoById)
		apiv1.PUT("/user/update/my", user.UpdateMyUser)
		apiv1.GET("/user/list/info", user.ListUserInfoByPage)
		apiv1.Use(middleware.Auth(AdminRole))
		{
			apiv1.GET("/ping", ping.Ping)
			apiv1.GET("/user/get/:id", user.GetUserById)
			apiv1.POST("/user/add", user.AddUser)
			apiv1.PUT("/user/update", user.UpdateUser)
			apiv1.DELETE("/user/delete/:id", user.DeleteUser)
			apiv1.GET("/user/list", user.ListUserByPage)
		}
	}
	return r
}
