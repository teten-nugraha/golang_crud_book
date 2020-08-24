package models

import "time"

type Kategori struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nama      string    `gorm:"size:255;not null;unique" json:"nama"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
