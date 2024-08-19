package repository

import "gorm.io/gorm"

type SeniorRepository struct {
	Database *gorm.DB
}

func NewSeniorRepo(db *gorm.DB) *SeniorRepository {
	return &SeniorRepository{Database: db}
}
