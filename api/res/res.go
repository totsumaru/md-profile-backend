package api_res

import "github.com/totsumaru/md-profile-backend/src/profile/app"

// プロフィールのレスポンス
type ProfileAPIRes struct {
	ID           string `json:"id"`
	Slug         string `json:"slug"`
	Avatar       string `json:"avatar"`
	DisplayName  string `json:"display_name"`
	Introduction string `json:"introduction"`
	Link         struct {
		X         string `json:"x"`
		Instagram string `json:"instagram"`
		Github    string `json:"github"`
		Website   string `json:"website"`
	} `json:"link"`
	Markdown string `json:"markdown"`
}

// コンテキストのレスポンスをAPIのレスポンスに変換します
func CastToProfileAPIRes(profileRes app.Res) ProfileAPIRes {
	res := ProfileAPIRes{}
	res.ID = profileRes.ID
	res.Slug = profileRes.Slug
	res.Avatar = profileRes.Avatar
	res.DisplayName = profileRes.DisplayName
	res.Introduction = profileRes.Introduction
	res.Link.X = profileRes.Link.X
	res.Link.Instagram = profileRes.Link.Instagram
	res.Link.Github = profileRes.Link.Github
	res.Link.Website = profileRes.Link.Website
	res.Markdown = profileRes.Markdown

	return res
}
