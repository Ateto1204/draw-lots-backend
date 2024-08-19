package repository

import (
	"fmt"

	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/gorm"
)

type JuniorRepository struct {
	Database *gorm.DB
}

func NewJuniorRepo(db *gorm.DB) *JuniorRepository {
	return &JuniorRepository{Database: db}
}

func (repo *JuniorRepository) CreateJunior(junior *model.Junior) error {
	if err := repo.Database.Create(&junior).Error; err != nil {
		return err
	}
	return nil
}

func (repo *JuniorRepository) GetJunior(id string) *model.Junior {
	var junior model.Junior

	if err := repo.Database.First(&junior, "student_number = ?", id).Error; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Junior found:", junior)
	}
	return &junior
}
