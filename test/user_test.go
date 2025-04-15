package test

import (
	"go-smartcerti/controllers"
	"go-smartcerti/database"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Setup() *fiber.App {
	app := fiber.New()
	return app
}

var app = Setup()
func TestGetUser(t *testing.T) {
	database.DatabaseInit()
	
	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.GetAllUsers(c)
	})
 
	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "success get all users")
}

func TestGetUserByID(t *testing.T) {
	database.DatabaseInit()
	
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return controllers.GetUserByID(c)
	})

	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "success get user by id")
}

func TestCreateUser(t *testing.T) {
	database.DatabaseInit()
	
	app.Post("/users", func(c *fiber.Ctx) error {
		return controllers.CreateUser(c)
	})

	req := httptest.NewRequest("POST", "/users", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "success create user")
}

func TestUpdateUser(t *testing.T) {
	database.DatabaseInit()
	
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateUser(c)
	})

	req := httptest.NewRequest("PUT", "/users/1", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "success update user")
}

func TestDeleteUser(t *testing.T) {
	database.DatabaseInit()
	
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteUser(c)
	})

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "success delete user")
}