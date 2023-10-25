package app

import (
	"github.com/totsumaru/md-profile-backend/shared/errors"
	domain2 "github.com/totsumaru/md-profile-backend/src/profile/domain"
	"github.com/totsumaru/md-profile-backend/src/profile/gateway"
	"gorm.io/gorm"
)

// マークダウンを更新するリクエストです
type UpdateMarkdownReq struct {
	ID       string
	Markdown string
}

// マークダウンを更新します
func UpdateMarkdown(tx *gorm.DB, req UpdateMarkdownReq) (Res, error) {
	id, err := domain2.RestoreUUID(req.ID)
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

	md, err := domain2.NewMarkdown(req.Markdown)
	if err != nil {
		return Res{}, errors.NewError("マークダウンを作成できません", err)
	}

	if err = profile.UpdateMarkdown(md); err != nil {
		return Res{}, errors.NewError("マークダウンを更新できません", err)
	}

	if err = gw.Update(profile); err != nil {
		return Res{}, errors.NewError("プロフィールを更新できません", err)
	}

	return CreateRes(profile), nil
}
