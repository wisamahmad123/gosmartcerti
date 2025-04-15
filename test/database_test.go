package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func OpenConnection() *gorm.DB{
	var err error
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/gosmartcerti?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db, "Database connection should not be nil")
}