// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/27 00:04:00
// @Desc
package utils

import "strings"

// gin 翻译器 返回数据 修正
func RemoveTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
