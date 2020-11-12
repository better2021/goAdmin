package util

import (
	"github.com/mojocn/base64Captcha" // 注意这个版本要用v1.2.2的，1.3有问题
	"sync"
)

type CaptchaConfig struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

var (
	captchaConfig *CaptchaConfig
	captchaConfigOnce sync.Once
)

// 获取base64验证码基本配置
func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			Id:              "",
			CaptchaType:     "character",
			VerifyValue:     "",
			ConfigAudio:     base64Captcha.ConfigAudio{CaptchaLen: 6, Language: "zh"}, //声音验证码配置
			ConfigCharacter: base64Captcha.ConfigCharacter{
				Height: 60,
				Width: 240,
				//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
				Mode: base64Captcha.CaptchaModeNumber,
				ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
				ComplexOfNoiseDot: base64Captcha.CaptchaComplexLower,
				IsShowHollowLine: false,
				IsShowNoiseDot: false,
				IsShowNoiseText: false,
				IsShowSlimeLine: false,
				IsShowSineLine: false,
				CaptchaLen: 4,
			},
			ConfigDigit:     base64Captcha.ConfigDigit{ // 数字验证码配置
				Height: 80,
				Width: 240,
				MaxSkew: 0.7,
				DotCount: 80,
				CaptchaLen: 5,
			},
		}
	})
	return captchaConfig
}

//  验证 验证码是否正确
// captchaId: 存于session中
// verifyValue: 客户端发来的验证码
func VerfiyCaptcha(captchaId , verifyValue string) bool{
	verifyResult := base64Captcha.VerifyCaptcha(captchaId, verifyValue)
	return verifyResult
}
