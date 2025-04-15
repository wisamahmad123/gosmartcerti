package middleware

import (
	"go-smartcerti/database"
	"go-smartcerti/models"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth(c *fiber.Ctx) error {
	// header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized - missing token",
		})
	}

	// cut prefix "Bearer "
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse token MapClaims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized - invalid token",
		})
	}

	// claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration date
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < jwt.NewNumericDate(time.Now()).Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "unauthorized - token expired",
				})
			}
		}
		// find user 
		var user models.User
		database.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized - user not found",
			})
		}

		c.Locals("user", user)
		
	return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized - invalid token",
		})
	}
}

