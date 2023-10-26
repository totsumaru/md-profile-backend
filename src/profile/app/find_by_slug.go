package app

import (
	"github.com/totsumaru/md-profile-backend/src/profile/domain"
	"github.com/totsumaru/md-profile-backend/src/profile/gateway"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"gorm.io/gorm"
)

// slugでプロフィールを取得します
func FindBySlug(db *gorm.DB, slug string) (Res, error) {
	sl, err := domain.NewSlug(slug)
	if err != nil {
		return Res{}, errors.NewError("slugを作成できません", err)
	}

	gw, err := gateway.NewGateway(db)
	if err != nil {
		return Res{}, errors.NewError("Gatewayを作成できません", err)
	}

	profile, err := gw.FindBySlug(sl)
	if err != nil {
		return Res{}, errors.NewError("slugでプロフィールを取得できません", err)
	}

	return CreateRes(profile), nil
}
