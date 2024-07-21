package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goAiproject/docs"
	"goAiproject/global"
	"goAiproject/internal/middleware"
	"goAiproject/internal/routers/api"
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
	//url := ginSwagger.URL("http://localhost:8089/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ping := v1.NewPing()
	user := v1.NewUser()
	app := v1.NewApp()
	question := v1.NewQuestion()
	userAnswer := v1.NewUserAnswer()
	scoringResult := v1.NewScoringResult()
	post := v1.NewPost()
	postFavour := v1.NewPostFavour()
	postThumb := v1.NewPostThumb()

	upload := api.NewUpload()

	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	r.POST("/file/upload", upload.UploadFile)
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

		apiv1.POST("/app/add", app.AddApp)
		apiv1.DELETE("/app/delete/:id", app.DeleteApp)
		apiv1.GET("/app/get/info/:id", app.GetAppInfo)
		apiv1.GET("/app/list/info", app.ListAppInfoByPage)
		apiv1.GET("/app/list/my/info", app.ListMyAppInfoByPage)
		apiv1.PUT("/app/edit", app.EditApp)

		apiv1.POST("/question/add", question.AddQuestion)
		apiv1.DELETE("/question/delete/:id", question.DeleteQuestion)
		apiv1.GET("/question/get/info/:id", question.GetQuestionInfo)
		apiv1.GET("/question/list/info", question.ListQuestionInfoByPage)
		apiv1.GET("/question/list/my/info", question.ListMyQuestionInfoByPage)
		apiv1.PUT("/question/edit", question.EditQuestion)
		apiv1.POST("/question/ai/generate", question.AiGenerateQuestion)
		apiv1.POST("/question/ai/generate/sse", question.AiGenerateQuestionSSE)

		apiv1.POST("/userAnswer/add", userAnswer.AddUserAnswer)
		apiv1.DELETE("/userAnswer/delete/:id", userAnswer.DeleteUserAnswer)
		apiv1.GET("/userAnswer/get/info/:id", userAnswer.GetUserAnswerInfo)
		apiv1.GET("/userAnswer/list/info", userAnswer.ListUserAnswerInfoByPage)
		apiv1.GET("/userAnswer/list/my/info", userAnswer.ListMyUserAnswerInfoByPage)
		apiv1.PUT("/userAnswer/edit", userAnswer.EditUserAnswer)

		apiv1.POST("/scoringResult/add", scoringResult.AddScoringResult)
		apiv1.DELETE("/scoringResult/delete/:id", scoringResult.DeleteScoringResult)
		apiv1.GET("/scoringResult/get/info/:id", scoringResult.GetScoringResultInfo)
		apiv1.GET("/scoringResult/list/info", scoringResult.ListScoringResultInfoByPage)
		apiv1.GET("/scoringResult/list/my/info", scoringResult.ListMyScoringResultInfoByPage)
		apiv1.PUT("/scoringResult/edit", scoringResult.EditScoringResult)

		apiv1.POST("/post/add", post.AddPost)
		apiv1.DELETE("/post/delete/:id", post.DeletePost)
		apiv1.GET("/post/get/info/:id", post.GetPostInfo)
		apiv1.GET("/post/list/info", post.ListPostInfoByPage)
		apiv1.GET("/post/list/my/info", post.ListMyPostInfoByPage)
		apiv1.PUT("/post/edit", post.EditPost)

		apiv1.POST("/postFavour", postFavour.DoPostFavour)
		apiv1.GET("/postFavour/list/my", postFavour.ListMyFavourPostByPage)
		apiv1.GET("/postFavour/list", postFavour.ListFavourPostByPage)

		apiv1.GET("/postThumb", postThumb.DoPostThumb)
		apiv1.Use(middleware.Auth(AdminRole))
		{
			apiv1.GET("/ping", ping.Ping)
			apiv1.GET("/user/get/:id", user.GetUserById)
			apiv1.POST("/user/add", user.AddUser)
			apiv1.PUT("/user/update", user.UpdateUser)
			apiv1.DELETE("/user/delete/:id", user.DeleteUser)
			apiv1.GET("/user/list", user.ListUserByPage)

			apiv1.PUT("/app/update", app.UpdateApp)
			apiv1.GET("/app/list", app.ListAppByPage)
			apiv1.PUT("/app/review", app.ReviewApp)

			apiv1.PUT("/question/update", question.UpdateQuestion)
			apiv1.GET("/question/list", question.ListQuestionByPage)

			apiv1.PUT("/userAnswer/update", userAnswer.UpdateUserAnswer)
			apiv1.GET("/userAnswer/list", userAnswer.ListUserAnswerByPage)

			apiv1.PUT("/scoringResult/update", scoringResult.UpdateScoringResult)
			apiv1.GET("/scoringResult/list", scoringResult.ListScoringResultByPage)

			apiv1.PUT("/post/update", post.UpdatePost)
			apiv1.GET("/post/list", post.ListPostByPage)
		}

	}
	return r
}
