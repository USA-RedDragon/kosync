package gorm

import (
	"fmt"

	"github.com/USA-RedDragon/kosync/internal/config"
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
	case config.StorageTypeRedis:
		return nil, config.ErrBadStorageType
	default:
		return nil, config.ErrBadStorageType
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Gorm{
		db: db,
	}, nil
}

func (s *Gorm) CreateUser(username, password string) error {
	return fmt.Errorf("not implemented")
}

func (s *Gorm) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, fmt.Errorf("not implemented")
}

func (s *Gorm) GetProgress(username, document string) (models.Progress, error) {
	return models.Progress{}, fmt.Errorf("not implemented")
}

func (s *Gorm) UpdateProgress(progress models.Progress) error {
	return fmt.Errorf("not implemented")
}
