package domain

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/shared/errors"
	"github.com/totsumaru/md-profile-backend/src/domain/link"
)

// プロフィールです
type Profile struct {
	id           UUID
	slug         Slug
	displayName  DisplayName
	introduction Introduction
	link         link.Link
	markdown     Markdown
}

// プロフィールを作成します
func NewProfile(
	id UUID,
	slug Slug,
	displayName DisplayName,
	introduction Introduction,
	link link.Link,
	markdown Markdown,
) (Profile, error) {
	res := Profile{
		id:           id,
		slug:         slug,
		displayName:  displayName,
		introduction: introduction,
		link:         link,
		markdown:     markdown,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// プロフィール情報を更新します
func (p *Profile) UpdateProfile(
	slug Slug,
	displayName DisplayName,
	introduction Introduction,
	link link.Link,
) error {
	p.slug = slug
	p.displayName = displayName
	p.introduction = introduction
	p.link = link

	if err := p.validate(); err != nil {
		return errors.NewError("検証に失敗しました")
	}

	return nil
}

// マークダウンを更新します
func (p *Profile) UpdateMarkdown(markdown Markdown) error {
	p.markdown = markdown

	if err := p.validate(); err != nil {
		return errors.NewError("検証に失敗しました")
	}

	return nil
}

// IDを取得します
func (p Profile) ID() UUID {
	return p.id
}

// slugを取得します
func (p Profile) Slug() Slug {
	return p.slug
}

// 表示名を取得します
func (p Profile) DisplayName() DisplayName {
	return p.displayName
}

// 自己紹介を取得します
func (p Profile) Introduction() Introduction {
	return p.introduction
}

// リンクを取得します
func (p Profile) Link() link.Link {
	return p.link
}

// マークダウンを取得します
func (p Profile) Markdown() Markdown {
	return p.markdown
}

// 検証します
func (p Profile) validate() error {
	return nil
}

// 構造体からJSONに変換します
func (p Profile) MarshalJSON() ([]byte, error) {
	data := struct {
		ID           UUID         `json:"id"`
		Slug         Slug         `json:"slug"`
		DisplayName  DisplayName  `json:"display_name"`
		Introduction Introduction `json:"introduction"`
		Link         link.Link    `json:"link"`
		Markdown     Markdown     `json:"markdown"`
	}{
		ID:           p.id,
		Slug:         p.slug,
		DisplayName:  p.displayName,
		Introduction: p.introduction,
		Link:         p.link,
		Markdown:     p.markdown,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (p *Profile) UnmarshalJSON(b []byte) error {
	var data struct {
		ID           UUID         `json:"id"`
		Slug         Slug         `json:"slug"`
		DisplayName  DisplayName  `json:"display_name"`
		Introduction Introduction `json:"introduction"`
		Link         link.Link    `json:"link"`
		Markdown     Markdown     `json:"markdown"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	p.id = data.ID
	p.slug = data.Slug
	p.displayName = data.DisplayName
	p.introduction = data.Introduction
	p.link = data.Link
	p.markdown = data.Markdown

	return nil
}
