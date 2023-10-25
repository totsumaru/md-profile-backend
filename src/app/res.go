package app

import "github.com/totsumaru/md-profile-backend/src/domain"

// レスポンスです
type Res struct {
	ID           string
	Slug         string
	DisplayName  string
	Introduction string
	Link         struct {
		X         string
		Instagram string
		Github    string
		Website   string
	}
	Markdown string
}

// プロフィールをレスポンスに変換します
func CreateRes(p domain.Profile) Res {
	res := Res{}
	res.ID = p.ID().String()
	res.Slug = p.Slug().String()
	res.DisplayName = p.DisplayName().String()
	res.Introduction = p.Introduction().String()
	res.Link.X = p.Link().X().String()
	res.Link.Instagram = p.Link().Instagram().String()
	res.Link.Github = p.Link().Github().String()
	res.Link.Website = p.Link().Website().String()
	res.Markdown = p.Markdown().String()

	return res
}
