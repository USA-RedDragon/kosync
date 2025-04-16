package gorm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/USA-RedDragon/kosync/internal/config"
	storeErrs "github.com/USA-RedDragon/kosync/internal/store/errors"
	"github.com/USA-RedDragon/kosync/internal/store/models"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func NewGormStore(cfg *config.Config) (*Gorm, error) {
	var dialect gorm.Dialector
	switch cfg.Storage.Type {
	case config.StorageTypeSQLite:
		dialect = sqlite.Open(cfg.Storage.DSN)
	case config.StorageTypePostgres:
		dialect = postgres.Open(cfg.Storage.DSN)
	case config.StorageTypeMySQL:
		dialect = mysql.Open(cfg.Storage.DSN)
	default:
		return nil, config.ErrBadStorageType
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Progress{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Gorm{
		db: db,
	}, nil
}

func (s *Gorm) CreateUser(username, password string) error {
	if err := s.db.Create(&models.User{
		Username: username,
		Password: password,
	}).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *Gorm) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, storeErrs.ErrUserNotFound
		}
		return models.User{}, fmt.Errorf("failed to get user by username: %w", err)
	}

	return user, nil
}

func (s *Gorm) GetProgress(username, document string) (models.Progress, error) {
	var progress models.Progress
	if err := s.db.Where("user = ? AND document = ?", strings.ToLower(username), strings.ToLower(document)).First(&progress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Progress{}, storeErrs.ErrProgressNotFound
		}
		return models.Progress{}, fmt.Errorf("failed to get progress: %w", err)
	}

	return progress, nil
}

func (s *Gorm) UpdateProgress(progress models.Progress) error {
	if err := s.db.Save(&progress).Error; err != nil {
		return fmt.Errorf("failed to update progress: %w", err)
	}

	return nil
}
