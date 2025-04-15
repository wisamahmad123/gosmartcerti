package controllers

import (
	"go-smartcerti/database"
	"go-smartcerti/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllMatKul(c *fiber.Ctx) error {
	var matkuls []*models.MataKuliah
	if err := database.DB.Preload("User").Find(&matkuls).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to get all matkuls",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all matkuls",
		"matkuls": matkuls,
	})
}

func GetMatKulByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var matkul models.MataKuliah
	if err := database.DB.Preload("User").Where("id = ?", id).First(&matkul).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "matkul not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get matkul by id",
		"matkul":  matkul,
	})
}

func CreateMataKuliah(c *fiber.Ctx) error {
	mataKuliah := new(models.MataKuliah)

	if err := c.BodyParser(mataKuliah); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create mataKuliah",
		})
	}

	if err := database.DB.Create(&mataKuliah).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create mataKuliah",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "success create mataKuliah",
		"mataKuliah": mataKuliah,
	})
}

func UpdateMataKuliah(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil mataKuliah dari DB
	var mataKuliah models.MataKuliah
	if err := database.DB.First(&mataKuliah, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "mata kuliah not found",
		})
	}

	// Struct sementara untuk menerima input body
	var body struct {
		NamaMataKuliah string `json:"nama_matakuliah"`
		KodeMataKuliah string `json:"kode_matakuliah"`
		Deskripsi      string `json:"deskripsi"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse body",
			"error":   err.Error(),
		})
	}

	// Update hanya field yang tidak kosong
	if body.NamaMataKuliah != "" {
		mataKuliah.NamaMataKuliah = body.NamaMataKuliah
	}
	if body.KodeMataKuliah != "" {
		mataKuliah.KodeMataKuliah = body.KodeMataKuliah
	}
	if body.Deskripsi != "" {
		mataKuliah.Deskripsi = body.Deskripsi
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&mataKuliah).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update mata kuliah",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "success update mata kuliah",
		"mataKuliah": mataKuliah,
	})
}

func DeleteMataKuliah(c *fiber.Ctx) error {
	id := c.Params("id")
	var mataKuliah models.MataKuliah
	if err := database.DB.Where("id = ?", id).Delete(&mataKuliah).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete mataKuliah",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete mataKuliah",
	})
}
