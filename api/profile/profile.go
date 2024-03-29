package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/api/internal"
	"github.com/totsumaru/md-profile-backend/src/profile/app"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/shared/errors/api_err"
	"github.com/totsumaru/md-profile-backend/src/shared/verify"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	internal.ProfileAPIRes
}

// tokenからプロフィールを取得します
func FindProfileByAccessToken(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/profile", func(c *gin.Context) {
		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		res := Res{}

		// プロフィールを取得します
		err := func() error {
			backendProfile, err := app.FindByID(db, verifyRes.ID)
			if err != nil {
				return errors.NewError("idでプロフィールを取得できません", err)
			}

			res.ProfileAPIRes = internal.CastToProfileAPIRes(backendProfile)

			return nil
		}()
		if err != nil {
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", err))
			return
		}

		c.JSON(200, res)
	})
}
