package link

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/shared/errors"
)

const (
	// githubのアカウントの最大文字数
	GithubMaxLen = 39
)

// githubのアカウントです
type Github struct {
	value string
}

// githubのアカウントを作成します
func NewGithub(value string) (Github, error) {
	res := Github{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// githubのアカウントを取得します
func (g Github) String() string {
	return g.value
}

// githubのアカウントが存在しているか確認します
func (g Github) IsEmpty() bool {
	return g.value == ""
}

// githubのアカウントを検証します
//
// 空を許容します。
func (g Github) validate() error {
	if g.value == "" {
		return nil
	}

	if len(g.value) > GithubMaxLen {
		return errors.NewError("githubのアカウントの最大文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (g Github) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: g.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (g *Github) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	g.value = data.Value

	return nil
}
