package models

import (
	"time"

	"gorm.io/gorm"
)

type Sertifikasi struct {
	ID                 int            `json:"id" gorm:"primaryKey"`
	NamaSertifikasi    string         `json:"nama_sertifikasi" gorm:"not null"`
	Deskripsi          string         `json:"deskripsi"`
	JenisSertifikasi   string         `json:"jenis_sertifikasi" gorm:"not null"`
	TanggalSertifikasi time.Time      `json:"tanggal_sertifikasi" gorm:"not null"`
	Biaya              float64        `json:"biaya" gorm:"not null"`
	VendorID           int            `json:"vendor_id" gorm:"not null"`
	Vendor             *Vendor        `json:"vendor" gorm:"foreignKey:VendorID;onDelete:CASCADE"`
	User               []*User        `json:"user_ids" gorm:"many2many:user_sertifikasi;onDelete:CASCADE"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
