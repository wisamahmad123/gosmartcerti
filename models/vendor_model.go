package models

import (
	"time"

	"gorm.io/gorm"
)

type Vendor struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	NamaVendor  string         `json:"nama_vendor" gorm:"not null"`
	Alamat      string         `json:"alamat" gorm:"not null"`
	Telepon     string         `json:"telepon" gorm:"not null"`
	Email       string         `json:"email" gorm:"not null"`
	Website     string         `json:"website" gorm:"not null"`
	Deskripsi   string         `json:"deskripsi" gorm:"not null"`
	JenisVendor string         `json:"jenis_vendor" gorm:"not null"`
	Pelatihan   []*Pelatihan   `json:"pelatihan" gorm:"foreignKey:VendorID"`
	Sertifikasi []*Sertifikasi `json:"sertifikasi" gorm:"foreignKey:VendorID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
