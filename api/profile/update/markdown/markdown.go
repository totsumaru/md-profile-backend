package markdown

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

// Markdownを更新します
func UpdateMarkdown(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/profile/update/markdown", func(c *gin.Context) {
		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		res := Res{}

		// Tx
		err := db.Transaction(func(tx *gorm.DB) error {
			req := app.UpdateMarkdownReq{
				ID:       verifyRes.ID,
				Markdown: c.PostForm("markdown"),
			}
			backendRes, err := app.UpdateMarkdown(tx, req)
			if err != nil {
				return errors.NewError("Markdownを更新できません", err)
			}

			res.ProfileAPIRes = internal.CastToProfileAPIRes(backendRes)

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, res)
	})
}
