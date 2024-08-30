package repository

import (
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/we-we-Web/draw-lots-backend/db"
	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/gorm"
)

type SeniorRepository struct {
	Database *gorm.DB
	Cache    *redis.Client
}

// MARK: - NewSeniorRepo -
func NewSeniorRepo(db *gorm.DB, rdb *redis.Client) *SeniorRepository {
	return &SeniorRepository{
		Database: db,
		Cache:    rdb,
	}
}

// MARK: - CreateSenior -
func (repo *SeniorRepository) CreateSenior(senior *model.Senior) error {
	if err := repo.Database.Create(&senior).Error; err != nil {
		return err
	}
	return nil
}

// MARK: - GetAllSeniors -
func (repo *SeniorRepository) GetAllSeniors() (*[]model.Senior, error) {
	var seniors []model.Senior
	if err := repo.Database.Find(&seniors).Error; err != nil {
		return nil, err
	}
	return &seniors, nil
}

// MARK: - GetSenior -
func (repo *SeniorRepository) GetSenior(id string) (*model.Senior, error) {
	val, err := repo.Cache.Get(db.Ctx, id).Result()

	if err == redis.Nil {
		var senior model.Senior
		if err := repo.Database.First(&senior, "student_number = ?", id).Error; err != nil {
			return nil, err
		}
		seniorJSON, err := json.Marshal(senior)
		if err != nil {
			return nil, err
		}
		repo.Cache.Set(db.Ctx, id, seniorJSON, 0)

		return &senior, nil
	} else if err != nil {
		return nil, err
	}

	var senior model.Senior
	if err := json.Unmarshal([]byte(val), &senior); err != nil {
		return nil, err
	}
	if senior.ChildrenId == nil {
		return nil, errors.New("omgomgomg")
	}
	return &senior, nil
}

// MARK: - UpdateSenior -
func (repo *SeniorRepository) UpdateSenior(senior *model.Senior) error {
	errCh := make(chan error, 2)
	go func() {
		if err := repo.Database.Save(senior).Error; err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()
	go func() {
		seniorJSON, err := json.Marshal(senior)
		if err != nil {
			errCh <- err
			return
		}
		id := senior.StudentNumber
		repo.Cache.Set(db.Ctx, id, seniorJSON, 0)
		errCh <- nil
	}()
	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}
	return nil
}
