package models

type Barang struct {
	UUID
	Nama         string     `gorm:"not null" json:"nama"`
	Harga        int        `gorm:"not null;check:harga > 0" json:"harga"`
	Stok         int        `gorm:"not null;check:stok >= 0" json:"stok"`
	Kode         string     `gorm:"unique;not null" json:"kode"`
	PerusahaanID string     `json:"perusahaan_id"`
	Perusahaan   Perusahaan `gorm:"foreignKey:PerusahaanID;references:ID;constraint:OnDelete:CASCADE;"`
}