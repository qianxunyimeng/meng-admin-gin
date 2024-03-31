// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 20:47:00
// @Desc
package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meng-admin-gin/common/respcode"
	"meng-admin-gin/utils"
	"net/http"
)

var Default = &response{}

// Error 失败数据处理
func Error(c *gin.Context, code int, msg string) {
	res := Default.Clone()
	if msg != "" {
		res.SetMsg(msg)
	}
	//res.SetTraceID(utils.GetTranceId(c))
	res.SetCode(int32(code))
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ErrorTrans(c *gin.Context, code int, msg map[string]string) {
	res := Default.Clone()
	res.SetMsg(msg)
	res.SetCode(int32(code))
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(utils.GetTranceId(c))
	res.SetCode(respcode.Success)

	fmt.Println(res)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

// Custum 兼容函数
func Custum(c *gin.Context, data gin.H) {
	data["requestId"] = utils.GetTranceId(c)
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}
