package verify

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/md-profile-backend/src/shared/seeker"
)

// Verifyのレスポンスです
type Res struct {
	ID          string
	AvatarURL   string
	DisplayName string
	X           string
}

// セッションを検証します
func VerifyToken(c *gin.Context) (bool, Res) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
		return false, Res{} // ヘッダーが不正またはトークンが存在しない場合は、空文字列を返します
	}

	tokenString := bearerToken[1]

	secret := os.Getenv("SUPABASE_JWT_SECRET")

	jwt.TimeFunc = func() time.Time {
		return time.Now().UTC().Add(time.Second * 20)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(secret), nil
	})

	// トークンのパースに失敗した場合、またはトークンが無効な場合は、falseと空のResを返します
	if err != nil || !token.Valid {
		return false, Res{}
	}

	// Claimsの型が期待どおりでない場合は、falseと空のResを返します
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, Res{}
	}

	// トークンが有効期限切れの場合は、falseと空のResを返します
	expiredAt := int64(claims["exp"].(float64))
	if expiredAt <= time.Now().Unix() {
		return false, Res{}
	}

	id := seeker.Str(claims, []string{"sub"})
	avatarURL := seeker.Str(claims, []string{"user_metadata", "avatar_url"})
	displayName := seeker.Str(claims, []string{"user_metadata", "full_name"})
	x := seeker.Str(claims, []string{"user_metadata", "user_name"})

	return true, Res{
		ID:          id,
		AvatarURL:   avatarURL,
		DisplayName: displayName,
		X:           x,
	}
}
