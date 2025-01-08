package util

import (
	"fangwu-backend/global"

	"github.com/mojocn/base64Captcha"
)

func LoadCaptcha() {
	store := base64Captcha.DefaultMemStore
	driver := base64Captcha.NewDriverString(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.NoiseCount,
		base64Captcha.OptionShowHollowLine,
		global.Config.Captcha.Length,
		"123456789abcdefghjkmnprstuvwxyz",
		nil, nil, nil)
	global.Captcha = base64Captcha.NewCaptcha(driver, store)
}
