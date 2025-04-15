package models

import (
	"time"

	"gorm.io/gorm"
)

type MataKuliah struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	NamaMataKuliah string         `json:"nama_matakuliah" gorm:"not null"`
	KodeMataKuliah string         `json:"kode_matakuliah" gorm:"not null"`
	Deskripsi      string         `json:"deskripsi"`
	User           []*User        `gorm:"many2many:user_matakuliah;"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
