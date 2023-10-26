package link

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// Xのアカウント名の最大文字数(twitterは14)
	XMaxLen = 14
)

// Xのアカウント名です
type X struct {
	value string
}

// Xのアカウント名を作成します
func NewX(value string) (X, error) {
	res := X{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// Xのアカウント名を取得します
func (x X) String() string {
	return x.value
}

// Xのアカウント名が存在しているか確認します
func (x X) IsEmpty() bool {
	return x.value == ""
}

// Xのアカウント名を検証します
//
// 空を許容します。
func (x X) validate() error {
	if x.value == "" {
		return nil
	}

	if len(x.value) > XMaxLen {
		return errors.NewError("Xのアカウント名の最大文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (x X) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: x.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (x *X) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	x.value = data.Value

	return nil
}
