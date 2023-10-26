package update

import (
	defaultErrors "errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/src/cloudflare"
	"github.com/totsumaru/md-profile-backend/src/profile/app"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/shared/errors/api_err"
	"github.com/totsumaru/md-profile-backend/src/shared/verify"
	"gorm.io/gorm"
)

// プロフィールを更新します
func UpdateProfile(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/profile/update", func(c *gin.Context) {
		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		// Tx
		err := db.Transaction(func(tx *gorm.DB) error {
			// ファイルが添付されていない場合はエラーにならない
			avatarFile, err := c.FormFile("avatar")
			if err != nil && !defaultErrors.Is(err, http.ErrMissingFile) {
				return errors.NewError("ファイルを取得できません")
			}

			// 画像をCloudflareにアップロードします
			avatarURL, err := cloudflare.Upload(c, avatarFile)
			if err != nil {
				return errors.NewError("画像をアプロードできません", err)
			}

			req := app.UpdateProfileReq{
				ID:           verifyRes.ID,
				Slug:         c.PostForm("slug"),
				AvatarURL:    avatarURL,
				DisplayName:  c.PostForm("display_name"),
				Introduction: c.PostForm("introduction"),
				X:            c.PostForm("x"),
				Instagram:    c.PostForm("instagram"),
				Github:       c.PostForm("github"),
				Website:      c.PostForm("website"),
			}
			_, err = app.UpdateProfile(tx, req)
			if err != nil {
				return errors.NewError("プロフィールを更新できません", err)
			}

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, nil)
	})
}
