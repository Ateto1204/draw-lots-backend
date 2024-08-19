package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
	"github.com/we-we-Web/draw-lots-backend/repository"
	"github.com/we-we-Web/draw-lots-backend/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=ewr1.clusters.zeabur.com user=root password=ARh6JwzaM27Q1Xe35um8KprB0f4sV9UH dbname=zeabur port=31718"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Admin{})
	fmt.Println("Database connected successfully!")

	adminRepo := repository.NewAdminRepo(db)
	seniorRepo := &repository.SeniorRepository{}
	juniorRepo := &repository.JuniorRepository{}
	service := service.NewService(adminRepo, seniorRepo, juniorRepo)

	router := SetUpRouter(service)
	router.Run(":8080")
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
