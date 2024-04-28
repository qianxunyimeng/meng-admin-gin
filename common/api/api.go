// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/28 21:33:00
// @Desc
package api

import (
	"errors"
	"fmt"
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"meng-admin-gin/common/response"
	"meng-admin-gin/common/service"
	"meng-admin-gin/core/storage"
	"meng-admin-gin/global"
)

type Api struct {
	Context *gin.Context
	Logger  *zap.Logger
	Orm     *gorm.DB
	Cache   storage.AdapterCache
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Logger.Error(err.Error())
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = global.MA_LOG
	//e.Errors = nil // 防止错误累计
	return e
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = response.Constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			e.AddError(err)
			break
		}
	}
	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err1 := vd.Validate(d); err1 != nil {
		e.AddError(err1)
	}
	return e
}

// MakeOrm 设置Orm DB
func (e *Api) MakeOrm() *Api {
	var err error
	if e.Logger == nil {
		err = errors.New("at MakeOrm logger is nil")
		e.AddError(err)
		return e
	}
	db := global.MA_DB
	e.Orm = db
	return e
}

func (e *Api) MakeService(c *service.Service) *Api {
	c.Log = e.Logger
	c.Orm = e.Orm
	c.Cache = global.MA_CACHE
	e.Cache = c.Cache
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error) {
	//switch err.(type) {
	//case error:
	//	errs, ok := err.(validator.ValidationErrors)
	//	if !ok {
	//		// 转换失败，返回原始错误信息
	//		response.Error(e.Context, code, err.(error).Error())
	//	} else {
	//		fmt.Println("err2::", utils.RemoveTopStruct2(errs.Translate(global.MA_TRANS)))
	//		response.Error(e.Context, code, utils.RemoveTopStruct2(errs.Translate(global.MA_TRANS)))
	//	}
	//case string:
	//	response.Error(e.Context, code, err)
	//
	//default:
	//	response.Error(e.Context, code, err)
	//}
	response.Error(e.Context, code, err, err.Error())
}

// Error 通常错误数据处理
func (e Api) ErrorWithMsg(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custum(e.Context, data)
}
