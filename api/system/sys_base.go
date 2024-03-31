// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 21:23:00
// @Desc
package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"meng-admin-gin/common/models/system"
	"meng-admin-gin/common/respcode"
	"meng-admin-gin/core/api"
	"meng-admin-gin/core/captcha"
	"meng-admin-gin/global"
	"meng-admin-gin/model/base/request"
	"meng-admin-gin/model/common/response"
	"meng-admin-gin/service/dto"
	sysService "meng-admin-gin/service/system"
	"meng-admin-gin/utils"
	"strconv"
	"time"
)

type SysBaseApi struct {
	api.Api
}

var jwtService = sysService.JwtService{}

// 用户登录
func (r *SysBaseApi) Login(c *gin.Context) {
	//r.MakeContext(c)
	var login request.Login
	s := sysService.SysBaseService{}
	err := r.MakeContext(c).MakeOrm().Bind(&login, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		transError, ok := err.(validator.ValidationErrors)
		if !ok {
			// 转换失败，返回原始错误信息
			r.Error(respcode.ErrorParam, err.Error())
			return
		}
		fmt.Println(utils.RemoveTopStruct(transError.Translate(global.MA_TRANS)))
		r.ErrorTrans(respcode.ErrorParam, utils.RemoveTopStruct(transError.Translate(global.MA_TRANS)))
		return
	}
	key := c.ClientIP()
	// 判断验证码是否开启
	openCaptcha := global.MA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.MA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, err := s.Cache.Get(key)
	if err != nil {
		expire := time.Second * time.Duration(openCaptchaTimeOut)
		s.Cache.Set(key, 1, int(expire))
	}
	cacheValue, _ := strconv.Atoi(v)
	var oc bool = openCaptcha == 0 || openCaptcha < cacheValue

	if !oc || (login.CaptchaId != "" && login.Captcha != "" && captcha.Verify(login.CaptchaId, login.Captcha, true)) {
		user, err := s.Login(&login)
		if err != nil {
			s.Log.Error("登录失败！用户名不存在或密码错误")
			s.Cache.Increase(key)
			r.Error(respcode.Error, "用户名不存在或密码错误")
			return
		}

		if user.Status != 1 {
			s.Log.Error("登陆失败! 用户被禁止登录!")
			s.Cache.Increase(key)
			r.Error(respcode.Error, "用户被禁止登录!")
			return
		}
		// 签发token
		r.TokenNext(*user)
		return
	}
	r.Cache.Increase(key)

	r.Error(respcode.Error, "验证码错误")
	return
}

func (r *SysBaseApi) TokenNext(user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.MA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(utils.BaseClaims{
		UserId:   user.UserId,
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)

	if err != nil {
		global.MA_LOG.Error("获取token失败!", zap.Error(err))
		r.Error(respcode.Error, "获取token失败!")
		return
	}

	if !global.MA_CONFIG.System.UseMultipoint {
		utils.SetCookie(r.Context, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		r.OK(dto.LoginResp{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功")
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.MA_LOG.Error("设置登录状态失败!", zap.Error(err))
			r.Error(respcode.Error, "设置登录状态失败!")
			return
		}
		utils.SetCookie(r.Context, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		r.OK(dto.LoginResp{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功")
		return
	} else if err != nil {
		global.MA_LOG.Error("设置登录状态失败!", zap.Error(err))
		r.Error(respcode.Error, "设置登录状态失败!")
		return
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			r.Error(respcode.Error, "jwt作废失败!")
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			r.Error(respcode.Error, "设置登录状态失败!")
			return
		}
		utils.SetCookie(r.Context, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		r.OK(dto.LoginResp{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功")
		return
	}

}

func (r *SysBaseApi) GenerateCaptcha(c *gin.Context) {
	err := r.MakeContext(c).Errors
	if err != nil {
		r.Error(500, "服务初始化失败！")
		return
	}

	key := c.ClientIP()
	// 判断验证码是否开启
	openCaptcha := global.MA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.MA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, err := global.MA_CACHE.Get(key)
	if err != nil {
		expire := time.Second * time.Duration(openCaptchaTimeOut)
		global.MA_CACHE.Set(key, 1, int(expire))
	}

	var vf bool
	cacheValue, _ := strconv.Atoi(v)
	if openCaptcha == 0 || openCaptcha < cacheValue {
		vf = true
	}

	if !vf {
		r.Error(500, "操作过于频繁，请稍后再试！")
	}

	id, b64s, _, err := captcha.DriverDigitCaptcha()
	if err != nil {
		r.Logger.Error("DriverDigitFunc error ", zap.String("err", err.Error()))
		r.Error(500, "验证码获取失败")
		return
	}
	r.Custom(gin.H{
		"code": respcode.Success,
		"data": gin.H{
			"captchaId":     id,
			"captchaBase64": b64s,
		},
		"msg": "验证码获取成功",
	})
}

func (r *SysBaseApi) Logout(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		global.MA_LOG.Error("登出失败!", zap.Error(err))
		response.Error(c, respcode.Error, "登出失败!")
		return
	}
	utils.ClearCookie(c)
	response.OK(c, nil, "登出成功")
}
