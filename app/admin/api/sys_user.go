package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/api"
	"meng-admin-gin/common/code"
)

type SysUserApi struct {
	api.Api
}

// 用户注册
func (r *SysUserApi) Register(c *gin.Context) {
	fmt.Println("用户注册...")
	s := service.SysUserService{}
	req := dto.SysUserRegisterReq{}
	// 校验信息
	err := r.MakeContext(c).MakeOrm().Bind(&req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(code.INVALID_PARAM, err)
		return
	}
	// 设置创建人
	req.SetCreateBy(1)
	err = s.Insert(&req)
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(code.ERROR, err)
		return
	}

	r.OK("注册cg", "注册成功")
}
