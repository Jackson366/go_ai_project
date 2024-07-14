package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettings struct {
	DB       int
	Host     string
	Password string
	Timeout  time.Duration
}

type JWTSettingS struct {
	Secret          string
	Issuer          string
	TokenExpireTime time.Duration
}

type AiSettingS struct {
	APIKey string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	// k, v对应配置文件中的key和value
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
