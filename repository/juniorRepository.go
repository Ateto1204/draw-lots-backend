package repository

import (
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/we-we-Web/draw-lots-backend/db"
	"github.com/we-we-Web/draw-lots-backend/model"
	"gorm.io/gorm"
)

type JuniorRepository struct {
	Database *gorm.DB
	Cache    *redis.Client
}

// MARK: - NewJuniorRepo -
func NewJuniorRepo(db *gorm.DB, rdb *redis.Client) *JuniorRepository {
	return &JuniorRepository{
		Database: db,
		Cache:    rdb,
	}
}

// MARK: - CreateJunior -
func (repo *JuniorRepository) CreateJunior(junior *model.Junior) error {
	if err := repo.Database.Create(&junior).Error; err != nil {
		return err
	}
	return nil
}

// MARK: - GetAllJuniors -
func (repo *JuniorRepository) GetAllJuniors() (*[]model.Junior, error) {
	var juniors []model.Junior
	if err := repo.Database.Find(&juniors).Error; err != nil {
		return nil, err
	}
	return &juniors, nil
}

// MARK: - GetJunior -
func (repo *JuniorRepository) GetJunior(id string) (*model.Junior, error) {
	val, err := repo.Cache.Get(db.Ctx, id).Result()

	if err == redis.Nil {
		var junior model.Junior
		if err := repo.Database.First(&junior, "student_number = ?", id).Error; err != nil {
			return nil, err
		}
		juniorJSON, err := json.Marshal(junior)
		if err != nil {
			return nil, err
		}
		repo.Cache.Set(db.Ctx, id, juniorJSON, 0)

		return &junior, nil
	} else if err != nil {
		return nil, err
	}

	var junior model.Junior
	if err := json.Unmarshal([]byte(val), &junior); err != nil {
		return nil, err
	}
	return &junior, nil
}

// MARK: - UpdateJunior -
func (repo *JuniorRepository) UpdateJunior(junior *model.Junior) error {
	errCh := make(chan error, 2)

	go func() {
		if err := repo.Database.Save(junior).Error; err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()
	go func() {
		juniorJSON, err := json.Marshal(junior)
		if err != nil {
			errCh <- err
			return
		}
		id := junior.StudentNumber
		errCh <- repo.Cache.Set(db.Ctx, id, juniorJSON, 0).Err()
	}()
	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}
	return nil
}
