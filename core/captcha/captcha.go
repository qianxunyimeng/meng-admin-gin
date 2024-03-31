// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/30 22:06:00
// @Desc
package captcha

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"meng-admin-gin/global"
)

func SetStore(s base64Captcha.Store) {
	base64Captcha.DefaultMemStore = s
}

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// 字符串验证码
func DriverStringCaptcha() (id, b64s string, answer string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	// 去掉 1和小些字母l
	e.DriverString = base64Captcha.NewDriverString(global.MA_CONFIG.Captcha.ImgHeight, global.MA_CONFIG.Captcha.ImgWidth, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, nil, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return captcha.Generate()
}

// 指定位数对数字验证码
func DriverDigitCaptcha() (id, b64s string, answer string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(global.MA_CONFIG.Captcha.ImgHeight, global.MA_CONFIG.Captcha.ImgWidth, global.MA_CONFIG.Captcha.KeyLong, 0.7, 80)
	driver := e.DriverDigit
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return captcha.Generate()
}

// 数字运算验证码
func DriverMathCaptcha() (id, b64s string, answer string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	rgbaColor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	// 生成driver,g高，宽 背景文字的，画线的调试，背景颜色的指针
	e.DriverMath = base64Captcha.NewDriverMath(global.MA_CONFIG.Captcha.ImgHeight, global.MA_CONFIG.Captcha.ImgWidth, 3, 1, &rgbaColor, nil, fonts)
	driver := e.DriverMath
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return captcha.Generate()
}

// Verify 校验验证码
func Verify(id, code string, clear bool) bool {
	return base64Captcha.DefaultMemStore.Verify(id, code, clear)
}
