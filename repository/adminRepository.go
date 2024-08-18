package repository

import (
	"fmt"

	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/gorm"
)

type AdminRepository struct {
	Database *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{Database: db}
}

func (repo *AdminRepository) CreateAdmin(admin *model.Admin) error {
	if err := repo.Database.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) GetAdmin(id string) *model.Admin {
	var admin model.Admin

	if err := repo.Database.First(&admin, "student_number = ?", id).Error; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Admin found:", admin)
	}
	return &admin
}
