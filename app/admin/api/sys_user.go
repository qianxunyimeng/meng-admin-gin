package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"meng-admin-gin/app/admin/dto"
	"meng-admin-gin/app/admin/models"
	"meng-admin-gin/app/admin/service"
	"meng-admin-gin/common/api"
	"meng-admin-gin/common/code"
	"meng-admin-gin/core/jwtauth/user"
	"meng-admin-gin/utils"
)

type SysUserApi struct {
	api.Api
}

// 用户注册
func (r SysUserApi) Register(c *gin.Context) {
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

func (r SysUserApi) GetInfo(c *gin.Context) {
	s := service.SysUserService{}
	err := r.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(code.INVALID_PARAM, err)
		return
	}

	userId := utils.GetUserID(c)
	sysUser := models.SysUser{}
	fmt.Println(userId)
	err = s.GetById(userId, &sysUser)

	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(code.ERROR, err)
		return
	}

	r.OK(gin.H{
		"user": gin.H{
			"userId":   sysUser.UserId,
			"userName": sysUser.UserName,
			"nickName": sysUser.NickName,
			"avatar":   sysUser.Avatar,
			"roleId":   sysUser.RoleId,
		},
		"roles":       []string{"admin"},
		"permissions": []string{"*:*:*"},
	}, "获取用户信息成功")
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [get]
// @Security Bearer
func (r SysUserApi) GetPage(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.SysUserGetPageReq{}
	err := r.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(code.INVALID_PARAM, err)
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]models.SysUser, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		r.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}

	r.PageOK(list, int(count), req.GetPageNum(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [get]
// @Security Bearer
func (r SysUserApi) Get(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.SysUserById{}
	err := r.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		r.Logger.Error(err.Error())
		r.Error(code.INVALID_PARAM, err)
		return
	}
	var object models.SysUser
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, &object)
	if err != nil {
		r.ErrorWithMsg(code.ERROR, err, "查询失败")
		return
	}
	r.OK(object, "查询成功")
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [put]
// @Security Bearer
func (e SysUserApi) Update(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [delete]
// @Security Bearer
func (e SysUserApi) Delete(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// UpdateStatus 修改用户状态
// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateSysUserStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/status [put]
// @Security Bearer
func (e SysUserApi) UpdateStatus(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.UpdateSysUserStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.UpdateStatus(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// ResetPwd 重置用户密码
// @Summary 重置用户密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.ResetSysUserPwdReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/reset [put]
// @Security Bearer
func (e SysUserApi) ResetPwd(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.ResetSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.ResetPwd(&req)
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// UpdatePwd
// @Summary 修改密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/set [put]
// @Security Bearer
func (e SysUserApi) UpdatePwd(c *gin.Context) {
	s := service.SysUserService{}
	req := dto.PassWord{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost); err != nil {
		req.NewPassword = string(hash)
	}

	err = s.UpdatePwd(user.GetUserId(c), req.OldPassword, req.NewPassword)
	if err != nil {
		e.Logger.Error(err.Error())
		e.ErrorWithMsg(code.ERROR, err, err.Error())
		return
	}

	e.OK(nil, "密码修改成功")
}
