package controllers

import (
	"go-smartcerti/database"
	"go-smartcerti/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllPelatihan(c *fiber.Ctx) error {
	var pelatihans []*models.Pelatihan
	if err := database.DB.Preload("User").Preload("Vendor").Find(&pelatihans).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to get all pelatihans",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "success get all pelatihans",
		"pelatihans": pelatihans,
	})
}

func GetPelatihanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var pelatihan models.Pelatihan
	if err := database.DB.Preload("User").Preload("Vendor").Where("id = ?", id).First(&pelatihan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "pelatihan not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "success get pelatihan by id",
		"pelatihan": pelatihan,
	})
}

func CreatePelatihan(c *fiber.Ctx) error {
	// Struct untuk menerima input JSON
	var body struct {
		NamaPelatihan  string `json:"nama_pelatihan"`
		Deskripsi      string `json:"deskripsi"`
		JenisPelatihan string `json:"jenis_pelatihan"`
		TanggalMulai   string `json:"tanggal_mulai"`   // string, akan di-parse ke time.Time
		TanggalSelesai string `json:"tanggal_selesai"` // string, akan di-parse ke time.Time
		Tempat         string `json:"tempat"`
		Biaya          string `json:"biaya"`     // biaya bisa jadi string, kita konversi ke float64
		VendorID       string `json:"vendor_id"` // vendor_id bisa jadi string, kita konversi ke int
		UserIDs        []int  `json:"user_ids"`  // daftar user yang ikut pelatihan
	}

	// Parse body JSON
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	// Konversi tanggal
	tanggalMulai, err := time.Parse(time.RFC3339, body.TanggalMulai)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid start date format",
			"error":   err.Error(),
		})
	}

	tanggalSelesai, err := time.Parse(time.RFC3339, body.TanggalSelesai)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid end date format",
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

	// Konversi vendor_id
	vendorID, err := strconv.Atoi(body.VendorID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid vendor_id format",
			"error":   err.Error(),
		})
	}

	// Buat instance pelatihan baru
	pelatihan := &models.Pelatihan{
		NamaPelatihan:  body.NamaPelatihan,
		Deskripsi:      body.Deskripsi,
		JenisPelatihan: body.JenisPelatihan,
		TanggalMulai:   tanggalMulai,
		TanggalSelesai: tanggalSelesai,
		Tempat:         body.Tempat,
		Biaya:          biaya,
		VendorID:       vendorID,
	}

	// Simpan pelatihan ke database
	if err := database.DB.Create(&pelatihan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create pelatihan",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Preload("Vendor").First(&pelatihan, pelatihan.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to load vendor data",
			"error":   err.Error(),
		})
	}

	// Menambahkan relasi user ke pelatihan
	if len(body.UserIDs) > 0 {
		var users []models.User
		if err := database.DB.Where("id IN ?", body.UserIDs).Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to find users",
				"error":   err.Error(),
			})
		}
		if err := database.DB.Model(&pelatihan).Association("User").Append(&users); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to associate users",
				"error":   err.Error(),
			})
		}
	}

	// Kembalikan response sukses
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "success create pelatihan",
		"pelatihan": pelatihan,
	})
}

func UpdatePelatihan(c *fiber.Ctx) error {
	id := c.Params("id")

	var pelatihan models.Pelatihan
	if err := database.DB.Preload("User").First(&pelatihan, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "pelatihan not found",
		})
	}

	// Body struct
	var body struct {
		NamaPelatihan  string    `json:"nama_pelatihan"`
		Deskripsi      string    `json:"deskripsi"`
		JenisPelatihan string    `json:"jenis_pelatihan"`
		TanggalMulai   time.Time `json:"tanggal_mulai"`
		TanggalSelesai time.Time `json:"tanggal_selesai"`
		Tempat         string    `json:"tempat"`
		Biaya          float64   `json:"biaya"`
		VendorID       int       `json:"vendor_id"`
		UserIDs        []uint    `json:"user_ids"`
		UpdateMode     string    `json:"update_mode"` // "replace" atau "append"
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	// Update data pelatihan
	if body.NamaPelatihan != "" {
		pelatihan.NamaPelatihan = body.NamaPelatihan
	}
	if body.JenisPelatihan != "" {
		pelatihan.JenisPelatihan = body.JenisPelatihan
	}
	if body.Deskripsi != "" {
		pelatihan.Deskripsi = body.Deskripsi
	}
	if !body.TanggalMulai.IsZero() {
		pelatihan.TanggalMulai = body.TanggalMulai
	}
	if !body.TanggalSelesai.IsZero() {
		pelatihan.TanggalSelesai = body.TanggalSelesai
	}
	if body.Tempat != "" {
		pelatihan.Tempat = body.Tempat
	}
	if body.Biaya != 0 {
		pelatihan.Biaya = body.Biaya
	}
	if body.VendorID != 0 {
		pelatihan.VendorID = body.VendorID
	}

	// Update daftar user (jika dikirim)
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
			if err := database.DB.Model(&pelatihan).Association("User").Append(&users); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "failed to append users",
					"error":   err.Error(),
				})
			}
		default: // default = replace
			if err := database.DB.Model(&pelatihan).Association("User").Replace(&users); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "failed to replace users",
					"error":   err.Error(),
				})
			}
		}
	}

	// Simpan perubahan pelatihan
	if err := database.DB.Save(&pelatihan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update pelatihan",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "success update pelatihan",
		"pelatihan": pelatihan,
	})
}

func DeletePelatihan(c *fiber.Ctx) error {
	id := c.Params("id")
	var pelatihan models.Pelatihan
	if err := database.DB.Where("id = ?", id).Delete(&pelatihan).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete pelatihan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete pelatihan",
	})
}
