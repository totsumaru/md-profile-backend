package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/totsumaru/md-profile-backend/api"
	"github.com/totsumaru/md-profile-backend/src/shared/database"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// .envが存在している場合は読み込み
	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	dialector := postgres.Open(os.Getenv("DB_URL"))
	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(errors.NewError("DBに接続できません", err))
	}

	// テーブルが存在していない場合のみテーブルを作成します
	// 存在している場合はスキーマを同期します
	if err = db.AutoMigrate(&database.Profile{}); err != nil {
		panic(errors.NewError("テーブルのスキーマが一致しません", err))
	}

	// gin
	engine := gin.Default()

	// CORSの設定
	// ここからCorsの設定
	engine.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://profio.jp",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
		},
		ExposeHeaders: []string{"Content-Length"},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		//MaxAge: 24 * time.Hour,
	}))

	// ルートを設定する
	api.RegisterRouter(engine, db)

	if err = engine.Run(":8080"); err != nil {
		log.Fatal("起動に失敗しました", err)
	}
}
