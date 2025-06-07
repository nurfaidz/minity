package models

type User struct {
	GormModel
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
