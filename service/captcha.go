package service

import (
	"fangwu-backend/global"
	"fangwu-backend/util"

	"github.com/mojocn/base64Captcha"
)

type CaptchaGet struct{}

type CaptchaVerify struct {
	CaptchaId string `json:"captcha_id,omitempty"`
	Captcha   string `json:"captcha,omitempty"`
}

type CaptchaResult struct {
	Id           string `json:"id,omitempty"`
	Base64String string `json:"base64_string,omitempty"`
}

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

func (c *CaptchaGet) Get() (result *CaptchaResult, resCode int, errDetail *util.ErrDetail) {
	var tmpRes CaptchaResult
	id, b64s, _, err := global.Captcha.Generate()
	if err != nil {
		return nil, util.ErrorFailToCreateCaptcha, util.GetErrDetail(err)
	}

	tmpRes = CaptchaResult{
		Id:           id,
		Base64String: b64s,
	}

	return &tmpRes, util.Success, nil
}

// 校验用户的验证码是否正确
func (c *CaptchaVerify) Verify() (permitted bool) {
	// 使用默认的内存存储来存储和验证验证码
	store := base64Captcha.DefaultMemStore
	// 调用 Verify 方法来验证验证码
	// 第一个参数是验证码的 ID，第二个参数是用户输入的验证码
	// 第三个参数表示验证成功后是否从存储中清除该验证码
	permitted = store.Verify(c.CaptchaId, c.Captcha, true)
	// 返回验证结果
	return permitted
}
