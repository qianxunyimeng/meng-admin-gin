// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 00:16:00
// @Desc
package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"meng-admin-gin/global"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired") // token 已过期
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token") // token 无效
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserId   int
	RoleCode string
	Username string
	NickName string
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.MA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := ParseDuration(global.MA_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.MA_CONFIG.JWT.ExpiresTime)
	fmt.Println("token 过期时间： ", ep)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"MA"},                    // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.MA_CONFIG.JWT.Issuer,               // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := global.MA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, TokenInvalid
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenExpired

	} else {
		return nil, TokenInvalid
	}
}
