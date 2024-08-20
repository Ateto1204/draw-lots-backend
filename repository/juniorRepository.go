package repository

import (
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

func (repo *JuniorRepository) GetAllJuniors() (*[]model.Junior, error) {
	var juniors []model.Junior
	if err := repo.Database.Find(&juniors).Error; err != nil {
		return nil, err
	}
	return &juniors, nil
}

func (repo *JuniorRepository) GetJunior(id string) (*model.Junior, error) {
	var junior model.Junior

	if err := repo.Database.First(&junior, "student_number = ?", id).Error; err != nil {
		return nil, err
	}
	return &junior, nil
}

func (repo *JuniorRepository) UpdateParentId(junior *model.Junior) error {
	if err := repo.Database.Save(junior).Error; err != nil {
		return err
	}
	return nil
}

func (repo *JuniorRepository) UpdateLineId(junior *model.Junior) error {
	if err := repo.Database.Save(junior).Error; err != nil {
		return err
	}
	return nil
}
