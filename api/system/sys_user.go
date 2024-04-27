package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"meng-admin-gin/common/dto"
	"meng-admin-gin/common/respcode"
	"meng-admin-gin/core/api"
	sysService "meng-admin-gin/service/system"
)

type SysUserApi struct {
	api.Api
}

// 用户注册
func (r *SysUserApi) Register(c *gin.Context) {
	fmt.Println("用户注册...")
	s := sysService.SysUserService{}
	req := dto.SysUserRegisterReq{}
	// 校验信息
	err := r.MakeContext(c).MakeOrm().Bind(&req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(respcode.ErrorParam, err)
		return
	}
	// 设置创建人
	req.SetCreateBy(1)
	err = s.Insert(&req)
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(respcode.Error, err)
		return
	}

	r.OK("注册cg", "注册成功")
}
