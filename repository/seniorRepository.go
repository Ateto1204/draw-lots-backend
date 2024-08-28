package repository

import (
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

func (repo *SeniorRepository) GetAllSeniors() (*[]model.Senior, error) {
	var seniors []model.Senior
	if err := repo.Database.Find(&seniors).Error; err != nil {
		return nil, err
	}
	return &seniors, nil
}

func (repo *SeniorRepository) GetSenior(id string) (*model.Senior, error) {
	var senior model.Senior

	if err := repo.Database.First(&senior, "student_number = ?", id).Error; err != nil {
		return nil, err
	}
	return &senior, nil
}

func (repo *SeniorRepository) UpdateSenior(senior *model.Senior) error {
	if err := repo.Database.Save(senior).Error; err != nil {
		return err
	}
	return nil
}
