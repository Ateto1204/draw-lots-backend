package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
	"github.com/we-we-Web/draw-lots-backend/repository"
	"github.com/we-we-Web/draw-lots-backend/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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

func SetUpRouter(s *service.Service) *gin.Engine {
	router := gin.Default()

	router.GET("/", service.Home)
	router.GET("/api", service.Api)

	router.POST("/api/admin", s.CreateAdmin)
	router.GET("/api/admin/:id", s.GetAdmin)

	router.POST("/api/senior", s.CreateSenior)
	router.GET("/api/senior/:id", s.GetSenior)

	router.POST("/api/junior", s.CreateJunior)
	router.GET("/api/junior/:id", s.GetJunior)

	return router
}
