// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 23:35:00
// @Desc 初始化翻译器
package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func InitTranslate(local string) (trans ut.Translator, err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(filed reflect.StructField) string {
			name := strings.SplitN(filed.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		uni := ut.New(enT, zhT, enT)

		trans, ok = uni.GetTranslator(local)

		if !ok {
			return nil, fmt.Errorf("could not find translator for %s", local)
		}

		//register translate
		// 注册翻译器
		switch local {
		case "zh", "zh-CN":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
