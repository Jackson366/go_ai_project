package app

import (
	"github.com/dgrijalva/jwt-go"
	"goAiproject/global"
	"time"
)

type Claims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(userId uint64) (string, error) {
	nowTime := time.Now()
	// 无法读取到配置文件的数据
	expireTime := nowTime.Add(global.JWTSetting.TokenExpireTime)
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 计算该token还有多久过期
func GetTokenRemainingTime(token string) (time.Duration, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return 0, err
	}
	return time.Unix(claims.ExpiresAt, 0).Sub(time.Now()), nil
}
