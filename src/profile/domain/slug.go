package domain

import (
	"encoding/json"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// slugの最大文字数
	SlugMaxLen = 40
)

// slugです
type Slug struct {
	value string
}

// slugを作成します
func NewSlug(value string) (Slug, error) {
	res := Slug{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// slugを取得します
func (s Slug) String() string {
	return s.value
}

// slugが存在しているか確認します
func (s Slug) IsEmpty() bool {
	return s.value == ""
}

// slugを検証します
func (s Slug) validate() error {
	if len([]rune(s.value)) > SlugMaxLen {
		return errors.NewError("slugの最大文字数を超えています")
	}

	if s.IsEmpty() {
		return errors.NewError("slugが空です")
	}

	return nil
}

// 構造体からJSONに変換します
func (s Slug) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: s.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (s *Slug) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	s.value = data.Value

	return nil
}
