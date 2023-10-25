package app

import (
	"github.com/totsumaru/md-profile-backend/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/domain"
	"github.com/totsumaru/md-profile-backend/src/domain/link"
	"github.com/totsumaru/md-profile-backend/src/gateway"
	"gorm.io/gorm"
)

// プロフィールを作成するリクエストです
type CreateProfileReq struct {
	supabaseID   string
	displayName  string
	introduction string
	x            string
}

// プロフィールを作成します
func CreateProfile(tx *gorm.DB, req CreateProfileReq) (Res, error) {
	id, err := domain.RestoreUUID(req.supabaseID)
	if err != nil {
		return Res{}, errors.NewError("IDを復元できません", err)
	}

	// slugは一番最初はIDを入れる
	slug, err := domain.NewSlug(req.supabaseID)
	if err != nil {
		return Res{}, errors.NewError("slugを作成できません", err)
	}

	name, err := domain.NewDisplayName(req.displayName)
	if err != nil {
		return Res{}, errors.NewError("表示名を作成できません", err)
	}

	intro, err := domain.NewIntroduction(req.introduction)
	if err != nil {
		return Res{}, errors.NewError("自己紹介を作成できません", err)
	}

	xAccount, err := link.NewX(req.x)
	if err != nil {
		return Res{}, errors.NewError("Xを作成できません", err)
	}

	l, err := link.NewLink(xAccount, link.Instagram{}, link.Github{}, link.Website{})
	if err != nil {
		return Res{}, errors.NewError("リンクを作成できません", err)
	}

	profile, err := domain.NewProfile(id, slug, name, intro, l, domain.Markdown{})
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
