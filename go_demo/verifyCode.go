package go_demo

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// captcha包实现 验证码的简单demo
// 参考： https://mp.weixin.qq.com/s/25gmXIKqbiWED2JhpRna7Q
// 		 https://github.com/GoLangStackDev/captcha-demo

// Captcha 方便后期扩展
type Captcha struct {}

// 单例
var captchaInstance *Captcha
func Instance() *Captcha {
	if captchaInstance==nil {
		captchaInstance = &Captcha{}
	}
	return captchaInstance
}

// CreateImage 创建图片验证码
func (this *Captcha) CreateImage() string {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	return captchaId
}

// Reload 重载
func (this *Captcha) Reload(captchaId string) bool {
	return captcha.Reload(captchaId)
}

// Verify 验证
func (this *Captcha) Verify(captchaId,val string) bool {
	return captcha.VerifyString(captchaId, val)
}

// GetImageByte 获取图片二进制流
func (this *Captcha) GetImageByte(captchaId string) []byte {
	var content bytes.Buffer
	err := captcha.WriteImage(&content, captchaId, captcha.StdWidth, captcha.StdHeight)
	if err!=nil {
		log.Println(err)
		return nil
	}
	return content.Bytes()
}

func main_verifyCode() {
	r := gin.New()
	// 创建
	// 这里方便看到效果 我用的 GET 请求，实际生产最好不要用 GET
	r.Handle("GET", "/captcha/create", func(c *gin.Context) {
		imgId := Instance().CreateImage()
		c.JSON(http.StatusOK,
			gin.H{
				"code": 200,
				"key": imgId,
				"url": "/captcha/img/"+imgId,
			})
	})
	// 现实图片
	r.Handle("GET", "/captcha/img/:key", func(c *gin.Context) {
		captchaId := c.Param("key")
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")
		c.Writer.Header().Set("Content-Type", "image/png")
		// 重载一次
		Instance().Reload(captchaId)
		// 输出图片
		c.Writer.Write(Instance().GetImageByte(captchaId))
	})
	// 校验
	r.Handle("GET", "/captcha/verify/:key/:val", func(c *gin.Context) {
		captchaId := c.Param("key")
		val := c.Param("val")
		if Instance().Verify(captchaId,val) {
			c.JSON(http.StatusOK, gin.H{"code": 200})
		}else{
			c.JSON(http.StatusOK, gin.H{"code": 400})
		}
	})

	r.Run(":8083")
}