package models

type User struct {
	UUID
	Username string `json:"username"`
	Password string `gorm:"not null" json:"password"`
	Nama     string `gorm:"not null" json:"name"`
	Tipe     string `gorm:"default:user" json:"type"`
}