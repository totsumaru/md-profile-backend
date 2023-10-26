package api

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/api/image/upload"
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

	create.CreateProfile(e, db)
	update.UpdateProfile(e, db)
	markdown.UpdateMarkdown(e, db)
	slug.FindProfileBySlug(e, db)
	profile.FindProfileByAccessToken(e, db)
	upload.UploadImage(e)
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
