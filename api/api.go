package api

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/api/profile"
	"github.com/totsumaru/md-profile-backend/api/profile/create"
	"github.com/totsumaru/md-profile-backend/api/profile/slug"
	"github.com/totsumaru/md-profile-backend/api/profile/update"
	"gorm.io/gorm"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine, db *gorm.DB) {
	Route(e)

	// プロフィールの作成
	create.CreateProfile(e, db)
	// プロフィールの更新
	update.UpdateProfile(e, db)
	// slugでプロフィールを取得します
	slug.FindProfileBySlug(e, db)
	// AccessTokenからプロフィールを取得します
	profile.FindProfileByAccessToken(e, db)
}

// ルートです
//
// Note: この関数は削除しても問題ありません
func Route(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
}
