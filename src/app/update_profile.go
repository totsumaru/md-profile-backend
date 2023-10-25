package app

import (
	"github.com/totsumaru/md-profile-backend/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/gateway"
	"gorm.io/gorm"
)

// プロフィール情報を更新するリクエストです
type UpdateProfileReq struct {
}

// プロフィール情報を更新します
func UpdateProfile(tx *gorm.DB, req UpdateProfileReq) (Res, error) {
	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return Res{}, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.FindByIDForUpdate(profile); err != nil {
		return Res{}, errors.NewError("プロフィールのレコードを作成できません", err)
	}

	// FOR UPDATEで取得します
}
