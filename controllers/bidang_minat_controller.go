package controllers

import (
	"go-smartcerti/database"
	"go-smartcerti/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllBidangMinats(c *fiber.Ctx) error {
	var bidangMinats []*models.BidangMinat
	if err := database.DB.Preload("User").Find(&bidangMinats).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to get all bidangMinats",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "success get all bidangMinats",
		"bidangMinats": bidangMinats,
	})
}

func GetBidangMinatByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var bidangMinat models.BidangMinat
	if err := database.DB.Preload("User").Where("id = ?", id).First(&bidangMinat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "bidangMinat not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "success get bidangMinat by id",
		"bidangMinat": bidangMinat,
	})
}

func CreateBidangMinat(c *fiber.Ctx) error {
	bidangMinat := new(models.BidangMinat)

	if err := c.BodyParser(bidangMinat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create bidangMinat",
		})
	}

	if err := database.DB.Create(&bidangMinat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create bidangMinat",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "success create bidangMinat",
		"bidangMinat": bidangMinat,
	})
}

func UpdateBidangMinat(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil bidangMinat dari DB
	var bidangMinat models.BidangMinat
	if err := database.DB.First(&bidangMinat, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "bidang minat not found",
		})
	}

	// Struct sementara untuk menerima input body
	var body struct {
		NamaBidangMinat string `json:"nama_bidang_minat"`
		Deskripsi       string `json:"deskripsi"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	// Update hanya field yang tidak kosong
	if body.NamaBidangMinat != "" {
		bidangMinat.NamaBidangMinat = body.NamaBidangMinat
	}
	if body.Deskripsi != "" {
		bidangMinat.Deskripsi = body.Deskripsi
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&bidangMinat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update bidang minat",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "success update bidang minat",
		"bidangMinat": bidangMinat,
	})
}


func DeleteBidangMinat(c *fiber.Ctx) error {
	id := c.Params("id")
	var bidangMinat models.BidangMinat
	if err := database.DB.Where("id = ?", id).Delete(&bidangMinat).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete bidangMinat",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete bidangMinat",
	})
}
