package repository

import (
	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/gorm"
)

type AdminRepository struct {
	Database *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepository {
	return &AdminRepository{Database: db}
}

func (repo *AdminRepository) CreateAdmin(admin *model.Admin) error {
	if err := repo.Database.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) GetAdmin(id string) (*model.Admin, error) {
	var admin model.Admin

	if err := repo.Database.First(&admin, "student_number = ?", id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
