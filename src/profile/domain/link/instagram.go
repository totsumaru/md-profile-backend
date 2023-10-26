package link

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// インスタグラムのアカウントの最大文字数
	InstagramMaxLen = 30
)

// インスタグラムのアカウントです
type Instagram struct {
	value string
}

// インスタグラムのアカウントを作成します
func NewInstagram(value string) (Instagram, error) {
	res := Instagram{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// インスタグラムのアカウントを取得します
func (i Instagram) String() string {
	return i.value
}

// インスタグラムのアカウントが存在しているか確認します
func (i Instagram) IsEmpty() bool {
	return i.value == ""
}

// インスタグラムのアカウントを検証します
//
// 空を許容します。
func (i Instagram) validate() error {
	if i.value == "" {
		return nil
	}

	if len([]rune(i.value)) > InstagramMaxLen {
		return errors.NewError("インスタグラムのアカウントの最大文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (i Instagram) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: i.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (i *Instagram) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	i.value = data.Value

	return nil
}
