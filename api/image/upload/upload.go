package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/src/cloudflare"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/shared/errors/api_err"
	"github.com/totsumaru/md-profile-backend/src/shared/verify"
)

// レスポンスです
type Res struct {
	ImageURL string `json:"image_url"`
}

// 画像をアップロードします
func UploadImage(e *gin.Engine) {
	e.POST("/api/image/upload", func(c *gin.Context) {
		// 認証
		isLogin, _ := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		res := Res{}

		err := func() error {
			imageFile, err := c.FormFile("image")
			if err != nil {
				return errors.NewError("ファイルを取得できません")
			}

			// 画像をCloudflareにアップロードします
			imageURL, err := cloudflare.Upload(c, imageFile)
			if err != nil {
				return errors.NewError("画像をアプロードできません", err)
			}

			res.ImageURL = imageURL

			return nil
		}()
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, res)
	})
}
