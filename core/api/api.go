// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 21:18:00
// @Desc
package api

import (
	"github.com/gin-gonic/gin"
	"meng-admin-gin/model/common/response"
)

type Api struct {
	Context *gin.Context
}

// 设置 http 上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
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
