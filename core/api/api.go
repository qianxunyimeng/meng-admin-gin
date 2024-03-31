// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 21:18:00
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
	"meng-admin-gin/core/service"
	"meng-admin-gin/core/storage"
	"meng-admin-gin/global"
	"meng-admin-gin/model/common/response"
)

type Api struct {
	Context *gin.Context
	Logger  *zap.Logger
	Orm     *gorm.DB
	Cache   storage.AdapterCache
	Errors  error
}

// 设置 http 上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = global.MA_LOG
	e.Errors = nil // 防止错误累计
	return e
}

// 设置 gorm
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

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Logger.Error(err.Error())
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	fmt.Println(bindings)
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
	// 自定义错误返回格式
	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err1 := vd.Validate(d); err1 != nil {
		e.AddError(err1)
	}
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, msg string) {
	response.Error(e.Context, code, msg)
}

func (e Api) ErrorTrans(code int, msg map[string]string) {
	response.ErrorTrans(e.Context, code, msg)
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
