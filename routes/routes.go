package routes

import (
	"go-smartcerti/controllers"
	"go-smartcerti/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// Public routes
	r.Post("/login", controllers.Login)
	r.Get("/validate", middleware.JwtAuth, controllers.Validate)

	// Protected routes
	users := r.Group("/users", middleware.JwtAuth)

	users.Get("/", controllers.GetAllUsers)
	users.Post("/", controllers.CreateUser)
	users.Get("/:id", controllers.GetUserByID)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)

	// Protected routes
	vendors := r.Group("/vendors", middleware.JwtAuth)

	vendors.Get("/", controllers.GetAllVendors)
	vendors.Post("/", controllers.CreateVendor)
	vendors.Get("/:id", controllers.GetVendorByID)
	vendors.Put("/:id", controllers.UpdateVendor)
	vendors.Delete("/:id", controllers.DeleteVendor)

	// Protected routes
	bidangMinat := r.Group("/bidangMinats", middleware.JwtAuth)
	bidangMinat.Get("/", controllers.GetAllBidangMinats)
	bidangMinat.Post("/", controllers.CreateBidangMinat)
	bidangMinat.Get("/:id", controllers.GetBidangMinatByID)
	bidangMinat.Put("/:id", controllers.UpdateBidangMinat)
	bidangMinat.Delete("/:id", controllers.DeleteBidangMinat)

	mataKuliah := r.Group("/mataKuliahs", middleware.JwtAuth)
	mataKuliah.Get("/", controllers.GetAllMatKul)
	mataKuliah.Post("/", controllers.CreateMataKuliah)
	mataKuliah.Get("/:id", controllers.GetMatKulByID)
	mataKuliah.Put("/:id", controllers.UpdateMataKuliah)
	mataKuliah.Delete("/:id", controllers.DeleteMataKuliah)

	pelatihan := r.Group("/pelatihans", middleware.JwtAuth)
	pelatihan.Get("/", controllers.GetAllPelatihan)
	pelatihan.Post("/", controllers.CreatePelatihan)
	pelatihan.Get("/:id", controllers.GetPelatihanByID)
	pelatihan.Put("/:id", controllers.UpdatePelatihan)
	pelatihan.Delete("/:id", controllers.DeletePelatihan)

	sertifikasi := r.Group("/sertifikasis", middleware.JwtAuth)
	sertifikasi.Get("/", controllers.GetAllSertifikasi)
	sertifikasi.Post("/", controllers.CreateSertifikasi)
	sertifikasi.Get("/:id", controllers.GetSertifikasiByID)
	sertifikasi.Put("/:id", controllers.UpdateSertifikasi)
	sertifikasi.Delete("/:id", controllers.DeleteSertifikasi)

}
