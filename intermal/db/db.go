package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

// postgresql://localhost:26257?sslmode=disable

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ docs_simplecrud_gorm"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Image{})
	return db
}
