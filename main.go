package main

import (
	"github.com/gin-gonic/gin"
	"goAiproject/global"
	"goAiproject/internal/model"
	"goAiproject/internal/routers"
	"goAiproject/pkg/logger"
	"goAiproject/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		panic(err)
	}
	err = setupDBEngine()
	if err != nil {
		panic(err)
	}
	err = setupLogger()
	if err != nil {
		panic(err)
	}
}

// @title AI_Project_DEMO
// @version 1.0
// @description jackson366 使用chatglm4的能力所开发的AI项目
// @termsOfService http://www.github.com
// @contact.name API Support
// @contact.url http://www.github.com
// @contact.email 2391815999@qq.com
// @license.name The MIT License (MIT)
// @license.url http://www.github.com
// @host localhost:8089
// @BasePath /api/v1
func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Ai", &global.AiSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.TokenExpireTime *= time.Second
	global.RedisSetting.Timeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngin(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
