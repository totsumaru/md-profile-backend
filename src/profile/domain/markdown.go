package domain

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// マークダウンの最大文字数
	MarkdownMaxLen = 3000
)

const InitialMarkdown = `
## 自己紹介

ここに自己紹介や実績を書きます。

- リスト1
  - リスト2
`

// マークダウンです
type Markdown struct {
	value string
}

// マークダウンを作成します
func NewMarkdown(value string) (Markdown, error) {
	res := Markdown{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// マークダウンを取得します
func (w Markdown) String() string {
	return w.value
}

// マークダウンが存在しているか確認します
func (w Markdown) IsEmpty() bool {
	return w.value == ""
}

// マークダウンを検証します
//
// 空を許容します。
func (w Markdown) validate() error {
	if w.value == "" {
		return nil
	}

	if len(w.value) > MarkdownMaxLen {
		return errors.NewError("マークダウンの最大文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (w Markdown) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: w.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (w *Markdown) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	w.value = data.Value

	return nil
}
