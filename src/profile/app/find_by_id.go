package app

import (
	"github.com/totsumaru/md-profile-backend/src/profile/domain"
	"github.com/totsumaru/md-profile-backend/src/profile/gateway"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"gorm.io/gorm"
)

// IDでプロフィールを取得します
func FindByID(db *gorm.DB, id string) (Res, error) {
	uID, err := domain.RestoreUUID(id)
	if err != nil {
		return Res{}, errors.NewError("IDを作成できません", err)
	}

	gw, err := gateway.NewGateway(db)
	if err != nil {
		return Res{}, errors.NewError("Gatewayを作成できません", err)
	}

	profile, err := gw.FindByID(uID)
	if err != nil {
		return Res{}, errors.NewError("idでプロフィールを取得できません", err)
	}

	return CreateRes(profile), nil
}
