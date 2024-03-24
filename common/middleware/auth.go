package middleware

import (
	"meng-admin-gin/common/middleware/handler"
	"meng-admin-gin/core"
	jwt "meng-admin-gin/core/jwtauth"
	"time"
)

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if core.MG_CONGIG.System.Env == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if core.MG_CONGIG.JWT.ExpiresTime != 0 {
			timeout = time.Duration(core.MG_CONGIG.JWT.ExpiresTime) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test zone",
		// Key:             []byte(config.JwtConfig.Secret),
		Key:             []byte(core.MG_CONGIG.JWT.SigningKey),
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
