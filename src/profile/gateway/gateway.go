package gateway

import (
	"encoding/json"

	domain2 "github.com/totsumaru/md-profile-backend/src/profile/domain"
	"github.com/totsumaru/md-profile-backend/src/shared/database"
	"github.com/totsumaru/md-profile-backend/src/shared/errors"
	"gorm.io/gorm"
)

type Gateway struct {
	tx *gorm.DB
}

// gatewayを作成します
func NewGateway(tx *gorm.DB) (Gateway, error) {
	if tx == nil {
		return Gateway{}, errors.NewError("引数が空です")
	}

	res := Gateway{
		tx: tx,
	}

	return res, nil
}

// プロフィールを新規作成します
//
// 同じIDのレコードが存在する場合はエラーを返します。
func (g Gateway) Create(u domain2.Profile) error {
	dbProfile, err := castToDBProfile(u)
	if err != nil {
		return errors.NewError("ドメインモデルをDBの構造体に変換できません", err)
	}

	// 新しいレコードをデータベースに保存
	result := g.tx.Create(&dbProfile)
	if result.Error != nil {
		return errors.NewError("レコードを保存できませんでした", result.Error)
	}
	// 主キー制約違反を検出（同じIDのレコードが既に存在する場合）
	if result.RowsAffected == 0 {
		return errors.NewError("既存のレコードが存在しています")
	}

	return nil
}

// 更新します
func (g Gateway) Update(u domain2.Profile) error {
	dbProfile, err := castToDBProfile(u)
	if err != nil {
		return errors.NewError("ドメインモデルをDBの構造体に変換できません", err)
	}

	// IDに基づいてレコードを更新
	result := g.tx.Model(&database.Profile{}).Where(
		"id = ?",
		dbProfile.ID,
	).Updates(&dbProfile)
	if result.Error != nil {
		return errors.NewError("更新できません", result.Error)
	}

	// 主キー制約違反を検出（指定されたIDのレコードが存在しない場合）
	if result.RowsAffected == 0 {
		return errors.NewError("レコードが存在しません")
	}

	return nil
}

// IDでプロフィールを取得します
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindByID(id domain2.UUID) (domain2.Profile, error) {
	res := domain2.Profile{}

	var dbProfile database.Profile
	if err := g.tx.First(&dbProfile, "id = ?", id.String()).Error; err != nil {
		return res, errors.NewError("IDでプロフィールを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModel(dbProfile)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// slugでプロフィールを取得します
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindBySlug(slug domain2.Slug) (domain2.Profile, error) {
	res := domain2.Profile{}

	var dbProfile database.Profile
	if err := g.tx.First(&dbProfile, "slug = ?", slug.String()).Error; err != nil {
		return res, errors.NewError("slugでプロフィールを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModel(dbProfile)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// FOR UPDATEでプロフィールを取得します
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindByIDForUpdate(id domain2.UUID) (domain2.Profile, error) {
	res := domain2.Profile{}

	var dbProfile database.Profile
	if err := g.tx.Set("gorm:query_option", "FOR UPDATE").First(
		&dbProfile, "id = ?", id.String(),
	).Error; err != nil {
		return res, errors.NewError("IDでプロフィールを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModel(dbProfile)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// =============
// private
// =============

// ドメインモデルをDBの構造体に変換します
func castToDBProfile(domainProfile domain2.Profile) (database.Profile, error) {
	res := database.Profile{}

	b, err := json.Marshal(&domainProfile)
	if err != nil {
		return res, errors.NewError("Marshalに失敗しました", err)
	}

	res.ID = domainProfile.ID().String()
	res.Slug = domainProfile.Slug().String()
	res.Data = b

	return res, nil
}

// DBの構造体からドメインモデルに変換します
func castToDomainModel(dbProfile database.Profile) (domain2.Profile, error) {
	res := domain2.Profile{}
	if err := json.Unmarshal(dbProfile.Data, &res); err != nil {
		return res, errors.NewError("Unmarshalに失敗しました", err)
	}

	return res, nil
}
