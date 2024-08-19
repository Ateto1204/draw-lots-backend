package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/we-we-Web/draw-lots-backend/model"
	"github.com/we-we-Web/draw-lots-backend/repository"
	"github.com/we-we-Web/draw-lots-backend/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&model.Admin{}, &model.Senior{}, &model.Junior{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	adminRepo := repository.NewAdminRepo(db)
	seniorRepo := repository.NewSeniorRepo(db)
	juniorRepo := repository.NewJuniorRepo(db)
	service := service.NewService(adminRepo, seniorRepo, juniorRepo)

	router := SetUpRouter(service)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func SetUpRouter(service *service.Service) *gin.Engine {
	router := gin.Default()

	router.POST("/api/admin", service.CreateAdmin)
	router.GET("/api/admin/:id", service.GetAdmin)

	router.POST("/api/senior", service.CreateSenior)
	router.GET("/api/senior/:id", service.GetSenior)

	router.POST("/api/junior", service.CreateJunior)
	router.GET("/api/junior/:id", service.GetJunior)

	return router
}
