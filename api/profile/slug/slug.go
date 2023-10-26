package slug

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/api/internal"
	"github.com/totsumaru/md-profile-backend/src/profile/app"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/shared/errors/api_err"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	internal.ProfileAPIRes
}

// slugでプロフィールを取得します
func FindProfileBySlug(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/profile/slug/:slug", func(c *gin.Context) {
		slug := c.Param("slug")

		res := Res{}

		// プロフィールを取得します
		err := func() error {
			backendProfile, err := app.FindBySlug(db, slug)
			if err != nil {
				return errors.NewError("slugでプロフィールを取得できません", err)
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
