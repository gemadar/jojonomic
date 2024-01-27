package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var db *gorm.DB
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	return db
}
