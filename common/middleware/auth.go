package middleware

import (
	"meng-admin-gin/common/middleware/handler"
	jwt "meng-admin-gin/core/jwtauth"
	"meng-admin-gin/global"
	"time"
)

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if global.MA_CONFIG.System.Env == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if global.MA_CONFIG.JWT.ExpiresTime != "" {
			//timeout = time.Duration(global.MA_CONFIG.JWT.ExpiresTime) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test zone",
		// Key:             []byte(config.JwtConfig.Secret),
		Key:             []byte(global.MA_CONFIG.JWT.SigningKey),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		//Authenticator:   handler.Authenticator,
		//Authorizator:    handler.Authorizator,
		//Unauthorized:    handler.Unauthorized,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

}
