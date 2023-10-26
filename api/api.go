package api

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/api/profile"
	"github.com/totsumaru/md-profile-backend/api/profile/create"
	"github.com/totsumaru/md-profile-backend/api/profile/slug"
	"github.com/totsumaru/md-profile-backend/api/profile/update"
	"github.com/totsumaru/md-profile-backend/api/profile/update/markdown"
	"gorm.io/gorm"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine, db *gorm.DB) {
	Route(e)

	// プロフィールを作成します
	create.CreateProfile(e, db)
	// プロフィールを更新します
	update.UpdateProfile(e, db)
	// Markdownを更新します
	markdown.UpdateMarkdown(e, db)
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
