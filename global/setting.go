package global

import (
	"goAiproject/pkg/logger"
	"goAiproject/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettings
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	AiSetting       *setting.AiSettingS
)
