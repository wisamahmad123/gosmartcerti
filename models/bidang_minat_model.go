package models

import (
	"time"

	"gorm.io/gorm"
)

type BidangMinat struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	NamaBidangMinat string         `json:"nama_bidang_minat" gorm:"not null"`
	User          []*User        `gorm:"many2many:user_bidangminat;"`
	Deskripsi       string         `json:"deskripsi"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
