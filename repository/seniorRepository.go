package repository

import (
	"fmt"

	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/gorm"
)

type SeniorRepository struct {
	Database *gorm.DB
}

func NewSeniorRepo(db *gorm.DB) *SeniorRepository {
	return &SeniorRepository{Database: db}
}

func (repo *SeniorRepository) CreateSenior(senior *model.Senior) error {
	if err := repo.Database.Create(&senior).Error; err != nil {
		return err
	}
	return nil
}

func (repo *SeniorRepository) GetSenior(id string) *model.Senior {
	var senior model.Senior

	if err := repo.Database.First(&senior, "student_number = ?", id).Error; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Junior found:", senior)
	}
	return &senior
}
