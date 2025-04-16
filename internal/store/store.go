package store

import (
	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/store/gorm"
	"github.com/USA-RedDragon/kosync/internal/store/models"
)

type Store interface {
	CreateUser(username, password string) error
	GetUserByUsername(username string) (models.User, error)
	GetProgress(username, document string) (models.Progress, error)
	UpdateProgress(progress models.Progress) error
}

func NewStore(cfg *config.Config) (Store, error) {
	switch cfg.Storage.Type {
	case config.StorageTypeSQLite:
		return gorm.NewGormStore(cfg)
	case config.StorageTypePostgres:
		return gorm.NewGormStore(cfg)
	case config.StorageTypeMySQL:
		return gorm.NewGormStore(cfg)
	default:
		return nil, config.ErrBadStorageType
	}
}
