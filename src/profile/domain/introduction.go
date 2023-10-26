package domain

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// 自己紹介の最大文字数(twitterは160)
	IntroductionMaxLen = 200
)

// 自己紹介です
type Introduction struct {
	value string
}

// 自己紹介を作成します
func NewIntroduction(value string) (Introduction, error) {
	res := Introduction{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 自己紹介を取得します
func (i Introduction) String() string {
	return i.value
}

// 自己紹介が存在しているか確認します
func (i Introduction) IsEmpty() bool {
	return i.value == ""
}

// 自己紹介を検証します
//
// 空を許容します。
func (i Introduction) validate() error {
	if i.value == "" {
		return nil
	}

	if len([]rune(i.value)) > IntroductionMaxLen {
		return errors.NewError("自己紹介の最大文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (i Introduction) MarshalJSON() ([]byte, error) {
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
func (i *Introduction) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	i.value = data.Value

	return nil
}
