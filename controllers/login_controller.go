package controllers

import (
	"go-smartcerti/database"
	"go-smartcerti/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	// Implement login logic here
	var body struct {
		Email	string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}

	// check email
	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "email not found",
		})
	}

	// comparing password with hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid password or email",
		})
	}

	// Generate JWT token 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
	})

	// sign token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 30), // 30 days
		HTTPOnly: true,  
		Secure:   true,  
		SameSite: "Lax", 
	})
	

	// Return the token in the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successful",
	})
}

func Validate(c *fiber.Ctx) error {
	if user, ok := c.Locals("user").(models.User); ok {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "token valid",
			"user":    user,
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "failed to retrieve user from context",
	})
}
