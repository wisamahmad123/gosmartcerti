package models

import (
	"time"

	"gorm.io/gorm"
)

type Pelatihan struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	NamaPelatihan  string         `json:"nama_pelatihan" gorm:"not null"`
	Deskripsi      string         `json:"deskripsi"`
	JenisPelatihan string         `json:"jenis_pelatihan" gorm:"not null"`
	TanggalMulai   time.Time      `json:"tanggal_mulai" gorm:"not null"`
	TanggalSelesai time.Time      `json:"tanggal_selesai" gorm:"not null"`
	Tempat         string         `json:"tempat" gorm:"not null"`
	Biaya          float64        `json:"biaya" gorm:"not null"`
	VendorID       int            `json:"vendor_id" gorm:"not null"`
	Vendor         *Vendor        `json:"vendor" gorm:"foreignKey:VendorID;onDelete:CASCADE"`
	User           []*User        `json:"user_ids" gorm:"many2many:user_pelatihan;onDelete:CASCADE"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
