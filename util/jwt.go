package util

import (
	"fangwu-backend/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

// ParseToken 验证用户token。这部分基本就是参照官方写法。
// 第一个参数是token字符串，第二个参数是结构体，第三个参数是jwt规定的解析函数，包含密钥
func ParseToken(token string) (*customClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &customClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(global.Config.Jwt.SecretKey), nil
		})

	if err != nil {
		return nil, err
	} else if claims, ok := tokenClaims.Claims.(*customClaims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GetUserId(c *gin.Context) (userId int64, resCode int, errDetail *ErrDetail) {
	accessToken := c.GetHeader("access_token")
	if accessToken == "" {
		return 0, ErrorAccessTokenNotFound, nil
	}
	//开始校验access_token
	customClaims, err := ParseToken(accessToken)
	if err != nil {
		return 0, ErrorFailToParseAccessToken, GetErrDetail(err)
	}
	//如果token过期
	if customClaims.ExpiresAt.Unix() < time.Now().Unix() {
		return 0, ErrorAccessTokenExpired, nil
	}

	return customClaims.UserId, Success, nil
}

// 构建载荷
func buildClaims(userId int64) customClaims {
	validityHours := time.Duration(global.Config.Jwt.ValidityDays) * 24 * time.Hour
	return customClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(validityHours)),
		},
	}
}

func GenerateToken(userId int64) (token string, resCode int, errDetail *ErrDetail) {
	claims := buildClaims(userId)
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenStruct.SignedString([]byte(global.Config.Jwt.SecretKey))
	if err != nil {
		return "", ErrorFailToSignToken, GetErrDetail(err)
	}

	return tokenString, Success, nil
}
