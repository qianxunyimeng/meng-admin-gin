// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 00:51:00
// @Desc
package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"meng-admin-gin/global"
	"net"
	"strings"
)

const COOKIE_TOKEN_KEY = "mg-token"
const TOKEN_FROM = "header: Authorization, query: token, param: token, cookie: mg-token"
const TOKEN_HEADER_NAME = "Bearer"

var ErrEmptyAuthHeader = errors.New("auth header is empty")
var ErrInvalidAuthHeader = errors.New("auth header is invalid")
var ErrEmptyQueryToken = errors.New("query token is empty")
var ErrEmptyParamToken = errors.New("param token is empty")
var ErrEmptyCookieToken = errors.New("cookie token is empty")

func ClearCookie(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	// 清除cookie httponly要设置为false，	不然客户端删除失败
	if net.ParseIP(host) != nil {
		c.SetCookie(COOKIE_TOKEN_KEY, "", -1, "/", "", false, false)
	} else {
		c.SetCookie(COOKIE_TOKEN_KEY, "", -1, "/", host, false, false)
	}
}

func SetCookie(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie(COOKIE_TOKEN_KEY, token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie(COOKIE_TOKEN_KEY, token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {

	var token string = ""
	// token 可能来自 header,query, param, cookie,循环来判断
	methods := strings.Split(TOKEN_FROM, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, _ = TokenFromHeader(c, v)
		case "query":
			token, _ = TokenFromQuery(c, v)
		case "cookie":
			token, _ = TokenFromCookie(c, v)
		case "param":
			token, _ = TokenFromParam(c, v)
		}
	}
	return token
}

// 从请求头解析token
func TokenFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == TOKEN_HEADER_NAME) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func TokenFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

func TokenFromCookie(c *gin.Context, key string) (string, error) {
	fmt.Println("token from cookie: ", key)
	cookie, err := c.Cookie(key)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if cookie == "" {
		return "", ErrEmptyCookieToken
	}

	return cookie, nil
}

func TokenFromParam(c *gin.Context, key string) (string, error) {
	token := c.Param(key)

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.MA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.UserId
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.BaseClaims.UserId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.Username
	}
}
