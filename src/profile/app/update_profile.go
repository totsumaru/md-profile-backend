package app

import (
	"github.com/totsumaru/md-profile-backend/src/profile/domain"
	"github.com/totsumaru/md-profile-backend/src/profile/domain/link"
	"github.com/totsumaru/md-profile-backend/src/profile/gateway"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"gorm.io/gorm"
)

// プロフィール情報を更新するリクエストです
type UpdateProfileReq struct {
	ID           string
	Slug         string
	AvatarURL    string
	DisplayName  string
	Introduction string
	X            string
	Instagram    string
	Github       string
	Website      string
}

// プロフィール情報を更新します
func UpdateProfile(tx *gorm.DB, req UpdateProfileReq) (Res, error) {
	id, err := domain.RestoreUUID(req.ID)
	if err != nil {
		return Res{}, errors.NewError("IDを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return Res{}, errors.NewError("Gatewayを作成できません", err)
	}

	// FOR UPDATEで取得します
	profile, err := gw.FindByIDForUpdate(id)
	if err != nil {
		return Res{}, errors.NewError("プロフィールのレコードを作成できません", err)
	}

	slug, err := domain.NewSlug(req.Slug)
	if err != nil {
		return Res{}, errors.NewError("slugを作成できません", err)
	}

	avatar, err := domain.NewAvatar(req.AvatarURL)
	if err != nil {
		return Res{}, errors.NewError("アバターを作成できません", err)
	}

	displayName, err := domain.NewDisplayName(req.DisplayName)
	if err != nil {
		return Res{}, errors.NewError("表示名を作成できません", err)
	}

	introduction, err := domain.NewIntroduction(req.Introduction)
	if err != nil {
		return Res{}, errors.NewError("自己紹介を作成できません", err)
	}

	x, err := link.NewX(req.X)
	if err != nil {
		return Res{}, errors.NewError("自己紹介を作成できません", err)
	}

	instagram, err := link.NewInstagram(req.Instagram)
	if err != nil {
		return Res{}, errors.NewError("Instagramを作成できません", err)
	}

	github, err := link.NewGithub(req.Github)
	if err != nil {
		return Res{}, errors.NewError("Githubを作成できません", err)
	}

	website, err := link.NewWebsite(req.Website)
	if err != nil {
		return Res{}, errors.NewError("websiteを作成できません", err)
	}

	l, err := link.NewLink(x, instagram, github, website)
	if err != nil {
		return Res{}, errors.NewError("リンクを作成できません", err)
	}

	if err = profile.UpdateProfile(slug, avatar, displayName, introduction, l); err != nil {
		return Res{}, errors.NewError("プロフィールを更新できません", err)
	}

	if err = gw.Update(profile); err != nil {
		return Res{}, errors.NewError("プロフィールを更新できません", err)
	}

	return CreateRes(profile), nil
}
