package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/api/res"
	"github.com/totsumaru/md-profile-backend/src/profile/app"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/shared/errors/api_err"
	"github.com/totsumaru/md-profile-backend/src/shared/verify"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Profile res.ProfileAPIRes `json:"profile"`
}

// プロフィールを作成します
func CreateProfile(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/profile/create", func(c *gin.Context) {
		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		// Tx
		apiRes := Res{}
		err := db.Transaction(func(tx *gorm.DB) error {
			req := app.CreateProfileReq{
				SupabaseID:  verifyRes.ID,
				AvatarURL:   verifyRes.AvatarURL,
				DisplayName: verifyRes.DisplayName,
				X:           verifyRes.X,
			}
			ctxRes, err := app.CreateProfile(tx, req)
			if err != nil {
				return errors.NewError("プロフィールを作成できません", err)
			}

			apiRes.Profile = res.CastToProfileAPIRes(ctxRes)

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, apiRes)
	})
}
