package app

import (
	"github.com/totsumaru/md-profile-backend/src/profile/domain"
	"github.com/totsumaru/md-profile-backend/src/profile/domain/link"
	"github.com/totsumaru/md-profile-backend/src/profile/gateway"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"gorm.io/gorm"
)

// プロフィールを作成するリクエストです
type CreateProfileReq struct {
	SupabaseID   string
	AvatarURL    string
	DisplayName  string
	Introduction string
	X            string
}

// プロフィールを作成します
func CreateProfile(tx *gorm.DB, req CreateProfileReq) (Res, error) {
	id, err := domain.RestoreUUID(req.SupabaseID)
	if err != nil {
		return Res{}, errors.NewError("IDを復元できません", err)
	}

	// slugは一番最初はIDを入れる
	slug, err := domain.NewSlug(req.SupabaseID)
	if err != nil {
		return Res{}, errors.NewError("slugを作成できません", err)
	}

	avatar, err := domain.NewAvatar(req.AvatarURL)
	if err != nil {
		return Res{}, errors.NewError("アバターを作成できません", err)
	}

	name, err := domain.NewDisplayName(req.DisplayName)
	if err != nil {
		return Res{}, errors.NewError("表示名を作成できません", err)
	}

	intro, err := domain.NewIntroduction(req.Introduction)
	if err != nil {
		return Res{}, errors.NewError("自己紹介を作成できません", err)
	}

	xAccount, err := link.NewX(req.X)
	if err != nil {
		return Res{}, errors.NewError("Xを作成できません", err)
	}

	l, err := link.NewLink(xAccount, link.Instagram{}, link.Github{}, link.Website{})
	if err != nil {
		return Res{}, errors.NewError("リンクを作成できません", err)
	}

	profile, err := domain.NewProfile(
		id, slug, avatar, name, intro, l, domain.Markdown{},
	)
	if err != nil {
		return Res{}, errors.NewError("プロフィールを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return Res{}, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(profile); err != nil {
		return Res{}, errors.NewError("プロフィールのレコードを作成できません", err)
	}

	return CreateRes(profile), nil
}
