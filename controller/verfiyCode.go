package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goAdmin/util"
	"net/http"
)

// @Summary 获取验证码
// @Description 获取验证码
// @Tags 验证码
// @Accept json
// @Produce json
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/getCode [get]
func GenerateCaptchaHandler(ctx *gin.Context)  {
	// 获取二维码配置
	captchaConfig := util.GetCaptchaConfig()

	// 创建base64图像验证码
	config := captchaConfig.ConfigCharacter

	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId,digitCap := base64Captcha.GenerateCaptcha(captchaConfig.Id,config)

	// 生成base64的png图片数据
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	ctx.JSON(http.StatusOK,gin.H{
		"data":gin.H{
			"img":base64Png,
			"captchaId":captchaId,
		},
	})
}
