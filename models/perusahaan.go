package models

type Perusahaan struct {
	UUID
	Nama   string   `gorm:"not null" json:"nama"`
	Alamat string   `gorm:"not null" json:"alamat"`
	NoTelp string   `gorm:"not null" json:"no_telp"`
	Kode   string   `gorm:"not null" json:"kode"`
}