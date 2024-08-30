package db

import (
	"log"
	"os"

	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		log.Println("Opened database successfully")
	}
	if err := db.AutoMigrate(&model.Admin{}, &model.Senior{}, &model.Junior{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	} else {
		log.Println("Migrated database successfully")
	}
	return db
}
