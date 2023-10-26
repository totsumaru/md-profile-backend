package link

import (
	"encoding/json"
	"strings"

	"github.com/totsumaru/md-profile-backend/src/shared/errors"
)

const (
	// webサイトの最大文字数
	WebsiteMaxLen = 300
)

// webサイトです
type Website struct {
	value string
}

// webサイトを作成します
func NewWebsite(value string) (Website, error) {
	res := Website{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// webサイトを取得します
func (w Website) String() string {
	return w.value
}

// webサイトが存在しているか確認します
func (w Website) IsEmpty() bool {
	return w.value == ""
}

// webサイトを検証します
//
// 空を許容します。
func (w Website) validate() error {
	if w.value == "" {
		return nil
	}

	if len([]rune(w.value)) > WebsiteMaxLen {
		return errors.NewError("webサイトの最大文字数を超えています")
	}

	if !(strings.HasPrefix(w.value, "http://") ||
		strings.HasPrefix(w.value, "https://")) {
		return errors.NewError("URLの形式ではありません")
	}

	return nil
}

// 構造体からJSONに変換します
func (w Website) MarshalJSON() ([]byte, error) {
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
func (w *Website) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	w.value = data.Value

	return nil
}
