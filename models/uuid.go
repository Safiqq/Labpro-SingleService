package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UUID struct {
	ID string `gorm:"primaryKey;not null" json:"id"`
}

func (t *UUID) BeforeCreate(tx *gorm.DB) (err error) {
    t.ID = uuid.NewString()
    return
}