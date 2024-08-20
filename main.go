package main

import (
	"log"
	"os"

	"github.com/we-we-Web/draw-lots-backend/model"
	"github.com/we-we-Web/draw-lots-backend/repository"
	"github.com/we-we-Web/draw-lots-backend/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// godotenv.Load()

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

	adminRepo := repository.NewAdminRepo(db)
	seniorRepo := repository.NewSeniorRepo(db)
	juniorRepo := repository.NewJuniorRepo(db)
	s := service.NewService(adminRepo, seniorRepo, juniorRepo)

	router := SetUpRouter(s)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
