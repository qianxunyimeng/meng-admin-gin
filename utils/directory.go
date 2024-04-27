// @Author shiqingliang
// @Date 2024/3/26 00:14:00
// @Desc 文件目录工具函数
package utils

import (
	"errors"
	"os"
)

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
