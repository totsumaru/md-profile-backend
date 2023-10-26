package database

// プロフィールのスキーマです
type Profile struct {
	ID   string `gorm:"type:uuid;primary_key;"`
	Slug string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Data []byte `gorm:"type:jsonb"`
}
