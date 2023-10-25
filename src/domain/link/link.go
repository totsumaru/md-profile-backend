package link

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/shared/errors"
)

// リンクです
type Link struct {
	x         X
	instagram Instagram
	github    Github
	website   Website
}

// リンクを作成します
func NewLink(
	x X,
	instagram Instagram,
	github Github,
	website Website,
) (Link, error) {
	res := Link{
		x:         x,
		instagram: instagram,
		github:    github,
		website:   website,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// Xを取得します
func (l Link) X() X {
	return l.x
}

// instagramを取得します
func (l Link) Instagram() Instagram {
	return l.instagram
}

// githubを取得します
func (l Link) Github() Github {
	return l.github
}

// websiteを取得します
func (l Link) Website() Website {
	return l.website
}

// 検証します
func (l Link) validate() error {
	return nil
}

// 構造体からJSONに変換します
func (l Link) MarshalJSON() ([]byte, error) {
	data := struct {
		X         X         `json:"x"`
		Instagram Instagram `json:"instagram"`
		Github    Github    `json:"github"`
		Website   Website   `json:"website"`
	}{
		X:         l.x,
		Instagram: l.instagram,
		Github:    l.github,
		Website:   l.website,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (l *Link) UnmarshalJSON(b []byte) error {
	var data struct {
		X         X         `json:"x"`
		Instagram Instagram `json:"instagram"`
		Github    Github    `json:"github"`
		Website   Website   `json:"website"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	l.x = data.X
	l.instagram = data.Instagram
	l.github = data.Github
	l.website = data.Website

	return nil
}
