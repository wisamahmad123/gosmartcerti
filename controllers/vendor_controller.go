package controllers

import (
	"go-smartcerti/database"
	"go-smartcerti/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllVendors(c *fiber.Ctx) error {
	var vendors []models.Vendor
	err := database.DB.Preload("Pelatihan").Find(&vendors).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to get all vendors",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all vendors",
		"vendors":   vendors,
	})
}

func GetVendorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var vendor models.Vendor
	if err := database.DB.Preload("Pelatihan").Preload("Sertifikasi").Where("id = ?", id).First(&vendor).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "vendor not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get vendor by id",
		"vendor":    vendor,
	})
}

func CreateVendor(c *fiber.Ctx) error {
	vendor := new(models.Vendor)

	if err := c.BodyParser(vendor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create vendor",
		})
	}

	if err := database.DB.Create(&vendor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create vendor",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success create vendor",
		"vendor": vendor,
	})
}

func UpdateVendor(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil vendor dari DB
	var vendor models.Vendor
	if err := database.DB.First(&vendor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "vendor not found",
		})
	}

	// Struct sementara untuk menerima body (optional fields)
	var body struct {
		NamaVendor  string `json:"nama_vendor"`
		Alamat      string `json:"alamat"`
		Telepon     string `json:"telepon"`
		Email       string `json:"email"`
		Website     string `json:"website"`
		Deskripsi   string `json:"deskripsi"`
		JenisVendor string `json:"jenis_vendor"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	if body.NamaVendor != "" {
		vendor.NamaVendor = body.NamaVendor
	}
	if body.Alamat != "" {
		vendor.Alamat = body.Alamat
	}
	if body.Telepon != "" {
		vendor.Telepon = body.Telepon
	}
	if body.Email != "" {
		vendor.Email = body.Email
	}
	if body.Website != "" {
		vendor.Website = body.Website
	}
	if body.Deskripsi != "" {
		vendor.Deskripsi = body.Deskripsi
	}
	if body.JenisVendor != "" {
		vendor.JenisVendor = body.JenisVendor
	}

	// Simpan perubahan
	if err := database.DB.Save(&vendor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update vendor",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update vendor",
		"vendor":  vendor,
	})
}


func DeleteVendor(c *fiber.Ctx) error {
	id := c.Params("id")
	var vendor models.Vendor
	if err := database.DB.Where("id = ?", id).Delete(&vendor).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete vendor",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete vendor",
	})
}