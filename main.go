package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/we-we-Web/draw-lots-backend/db"
	"github.com/we-we-Web/draw-lots-backend/repository"
	"github.com/we-we-Web/draw-lots-backend/service"
)

func main() {
	godotenv.Load()

	database := db.InitDB()
	rdb := db.InitRedis()

	adminRepo := repository.NewAdminRepo(database)
	seniorRepo := repository.NewSeniorRepo(database, rdb)
	juniorRepo := repository.NewJuniorRepo(database, rdb)
	s := service.NewService(adminRepo, seniorRepo, juniorRepo)

	router := service.SetUpRouter(s)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
