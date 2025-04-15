package test

import (
	"bytes"
	"encoding/json"
	"go-smartcerti/controllers"
	"go-smartcerti/database"
	"go-smartcerti/models"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func SetupTestDB() {
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/gosmartcerti?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	database.DB = db
}

func TestLoginSuccess(t *testing.T) {
	app := fiber.New()

	// Setup env SECRET_KEY untuk JWT
	os.Setenv("SECRET_KEY", "testsecret")

	// Setup database
	SetupTestDB()

	// Setup route
	app.Post("/login", controllers.Login)

	// Request body
	loginBody := map[string]string{
		"email":    "cb@gmail.com",
		"password": "12345678",
	}
	body, _ := json.Marshal(loginBody)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestLoginWrongPassword(t *testing.T) {
	app := fiber.New()
	os.Setenv("SECRET_KEY", "testsecret")
	SetupTestDB()

	app.Post("/login", controllers.Login)

	loginBody := map[string]string{
		"email":    "cb@gmail.com",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(loginBody)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

func TestLoginEmailNotFound(t *testing.T) {
	app := fiber.New()
	os.Setenv("SECRET_KEY", "testsecret")
	SetupTestDB()

	app.Post("/login", controllers.Login)

	loginBody := map[string]string{
		"email":    "notfound@example.com",
		"password": "whatever",
	}
	body, _ := json.Marshal(loginBody)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}
