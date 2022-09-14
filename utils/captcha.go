package utils

import (
	"log"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

// CaptchaID 获取验证ID
func CaptchaID(len int) string {
	return captcha.NewLen(len)
}

// CaptchaImg 验证图片
func CaptchaImg(c *gin.Context, id string) {
	if err := captcha.WriteImage(c.Writer, id, 240, 80); err != nil {
		log.Println("show captcha error", err)
	}
}

// CaptchaAudio 验证语音
func CaptchaAudio(c *gin.Context, id string) {
	if err := captcha.WriteAudio(c.Writer, id, "zh"); err != nil {
		log.Println("show captcha error", err)
	}
}

// CaptchaVerify 验证内容
func CaptchaVerify(id string, challenge string) bool {
	return captcha.VerifyString(id, challenge)
}
