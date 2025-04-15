package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Level       string         `json:"level" validate:"required,oneof=admin pimpinan dosen" gorm:"type:ENUM('admin','pimpinan','dosen');not null"`
	Name        string         `json:"name" form:"name" validate:"gte=6, lte=32" gorm:"not null"`
	Email       string         `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Password    string         `json:"password" form:"password" validate:"required,gte=8" gorm:"not null,column:password"`
	Phone       string         `json:"phone"`
	BidangMinat []*BidangMinat `json:"bidang_minat" gorm:"many2many:user_bidangminat;"`
	MataKuliah  []*MataKuliah  `json:"mata_kuliah" gorm:"many2many:user_matakuliah;"`
	Sertifikasi []*Sertifikasi `json:"sertifikasi" gorm:"many2many:user_sertifikasi;"`
	Pelatihan   []*Pelatihan   `json:"pelatihan" gorm:"many2many:user_pelatihan;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
