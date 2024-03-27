package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"meng-admin-gin/common/respcode"
	"meng-admin-gin/core/api"
	"meng-admin-gin/global"
	"meng-admin-gin/model/base/request"
	"meng-admin-gin/utils"
)

type BaseApi struct {
	api.Api
}

func (r *BaseApi) Login(c *gin.Context) {
	r.MakeContext(c)
	var login request.Login

	err := c.ShouldBind(&login)
	if err != nil {
		transError, ok := err.(validator.ValidationErrors)
		if !ok {
			// 转换失败，返回原始错误信息
			r.Error(respcode.ErrorParam, err, err.Error())
			return
		}
		fmt.Println(utils.RemoveTopStruct(transError.Translate(global.MA_TRANS)))
		r.ErrorTrans(respcode.ErrorParam, utils.RemoveTopStruct(transError.Translate(global.MA_TRANS)))
		return

		//c.JSON(500, gin.H{"msg": err.Error()})
		//return
	}

	fmt.Println(login)

	r.OK("token:111", "登录成功")
}
