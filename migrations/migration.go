package migrations

import (
	"fmt"
	"go-smartcerti/database"
	"go-smartcerti/models"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database...")
	}
	fmt.Println("Database migrated successfully...")
}