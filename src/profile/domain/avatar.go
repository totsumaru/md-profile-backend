package domain

import (
	"encoding/json"
	"strings"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// アバターURLの最大
	AvatarMaxLen = 300
)

// アバターのURLです
type Avatar struct {
	value string
}

// アバターを作成します
func NewAvatar(value string) (Avatar, error) {
	// Xの元のサイズだと48x48で小さいので、URLをリプレイスします
	value = strings.Replace(value, "normal", "400x400", -1)

	res := Avatar{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// アバターを取得します
func (a Avatar) String() string {
	return a.value
}

// アバターが存在しているか確認します
func (a Avatar) IsEmpty() bool {
	return a.value == ""
}

// アバターを検証します
//
// 空を許容します。
func (a Avatar) validate() error {
	if len([]rune(a.value)) > AvatarMaxLen {
		return errors.NewError("アバターの最大文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (a Avatar) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: a.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (a *Avatar) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	a.value = data.Value

	return nil
}
