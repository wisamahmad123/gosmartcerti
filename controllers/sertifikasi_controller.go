package controllers

import (
	"go-smartcerti/database"
	"go-smartcerti/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllSertifikasi(c *fiber.Ctx) error {
	var sertifikasis []*models.Sertifikasi
	if err := database.DB.Preload("User").Preload("Vendor").Find(&sertifikasis).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to get all sertifikasis",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "success get all sertifikasis",
		"sertifikasis": sertifikasis,
	})
}

func GetSertifikasiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var sertifikasi models.Sertifikasi
	if err := database.DB.Preload("User").Preload("Vendor").Where("id = ?", id).First(&sertifikasi).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "sertifikasi not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "success get sertifikasi by id",
		"sertifikasi": sertifikasi,
	})
}

func CreateSertifikasi(c *fiber.Ctx) error {
	// Body struct sementara
	var body struct {
		NamaSertifikasi    string   `json:"nama_sertifikasi"`
		Deskripsi          string   `json:"deskripsi"`
		JenisSertifikasi   string   `json:"jenis_sertifikasi"`
		TanggalSertifikasi string   `json:"tanggal_sertifikasi"`
		Biaya              string   `json:"biaya"`
		VendorID           string   `json:"vendor_id"`
		UserIDs            []uint   `json:"user_ids"`
	}

	// Parse body JSON
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	// Konversi tanggal sertifikasi
	tanggalSertifikasi, err := time.Parse(time.RFC3339, body.TanggalSertifikasi)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid tanggal_sertifikasi format",
			"error":   err.Error(),
		})
	}

	// Konversi biaya
	biaya, err := strconv.ParseFloat(body.Biaya, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid biaya format",
			"error":   err.Error(),
		})
	}

	// Konversi vendor ID
	vendorID, err := strconv.Atoi(body.VendorID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid vendor_id format",
			"error":   err.Error(),
		})
	}

	// Buat instance sertifikasi baru
	sertifikasi := &models.Sertifikasi{
		NamaSertifikasi:    body.NamaSertifikasi,
		Deskripsi:          body.Deskripsi,
		JenisSertifikasi:   body.JenisSertifikasi,
		TanggalSertifikasi: tanggalSertifikasi,
		Biaya:              biaya,
		VendorID:           vendorID,
	}

	// Simpan sertifikasi ke database
	if err := database.DB.Create(&sertifikasi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create sertifikasi",
			"error":   err.Error(),
		})
	}

	// Preload data vendor
	if err := database.DB.Preload("Vendor").First(&sertifikasi, sertifikasi.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to load vendor data",
			"error":   err.Error(),
		})
	}

	// Menambahkan relasi user ke sertifikasi
	if len(body.UserIDs) > 0 {
		var users []models.User
		if err := database.DB.Where("id IN ?", body.UserIDs).Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to find users",
				"error":   err.Error(),
			})
		}
		if err := database.DB.Model(&sertifikasi).Association("User").Append(&users); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to associate users",
				"error":   err.Error(),
			})
		}
	}

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "success create sertifikasi",
		"sertifikasi": sertifikasi,
	})
}


func UpdateSertifikasi(c *fiber.Ctx) error {
	id := c.Params("id")

	var sertifikasi models.Sertifikasi
	if err := database.DB.Preload("User").First(&sertifikasi, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "sertifikasi not found",
		})
	}

	// Body struct
	var body struct {
		NamaSertifikasi    string    `json:"nama_sertifikasi"`
		Deskripsi          string    `json:"deskripsi"`
		JenisSertifikasi   string    `json:"jenis_sertifikasi"`
		TanggalSertifikasi time.Time `json:"tanggal_sertifikasi"`
		Biaya              float64   `json:"biaya"`
		VendorID           int       `json:"vendor_id"`
		UserIDs            []uint    `json:"user_ids"`
		UpdateMode         string    `json:"update_mode"` // "replace" atau "append"
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	// Update field jika ada perubahan
	if body.NamaSertifikasi != "" {
		sertifikasi.NamaSertifikasi = body.NamaSertifikasi
	}
	if body.Deskripsi != "" {
		sertifikasi.Deskripsi = body.Deskripsi
	}
	if body.JenisSertifikasi != "" {
		sertifikasi.JenisSertifikasi = body.JenisSertifikasi
	}
	if !body.TanggalSertifikasi.IsZero() {
		sertifikasi.TanggalSertifikasi = body.TanggalSertifikasi
	}
	if body.Biaya != 0 {
		sertifikasi.Biaya = body.Biaya
	}
	if body.VendorID != 0 {
		sertifikasi.VendorID = body.VendorID
	}

	// Update relasi user
	if len(body.UserIDs) > 0 {
		var users []models.User
		if err := database.DB.Where("id IN ?", body.UserIDs).Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to find users",
				"error":   err.Error(),
			})
		}

		switch strings.ToLower(body.UpdateMode) {
		case "append":
			if err := database.DB.Model(&sertifikasi).Association("User").Append(&users); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "failed to append users",
					"error":   err.Error(),
				})
			}
		default: // replace
			if err := database.DB.Model(&sertifikasi).Association("User").Replace(&users); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "failed to replace users",
					"error":   err.Error(),
				})
			}
		}
	}

	// Simpan perubahan
	if err := database.DB.Save(&sertifikasi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update sertifikasi",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "success update sertifikasi",
		"sertifikasi": sertifikasi,
	})
}


func DeleteSertifikasi(c *fiber.Ctx) error {
	id := c.Params("id")
	var sertifikasi models.Sertifikasi
	if err := database.DB.Where("id = ?", id).Delete(&sertifikasi).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete sertifikasi",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete sertifikasi",
	})
}
