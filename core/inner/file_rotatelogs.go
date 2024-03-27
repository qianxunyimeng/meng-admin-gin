// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 00:40:00
// @Desc
package inner

import (
	"go.uber.org/zap/zapcore"
	"meng-admin-gin/global"
	"os"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.MA_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.MA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
