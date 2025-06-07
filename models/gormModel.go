package models

type GormModel struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_att"`
}
