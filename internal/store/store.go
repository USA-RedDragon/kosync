package store

import (
	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/store/gorm"
	"github.com/USA-RedDragon/kosync/internal/store/models"
	"github.com/USA-RedDragon/kosync/internal/store/redis"
)

type Store interface {
	CreateUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	GetProgress(book string) (*models.Progress, error)
	UpdateProgress(book string, progress *models.Progress) error
}

func NewStore(cfg *config.Config) (Store, error) {
	switch cfg.Storage.Type {
	case config.StorageTypeSQLite:
		return gorm.NewGormStore(cfg)
	case config.StorageTypePostgres:
		return gorm.NewGormStore(cfg)
	case config.StorageTypeMySQL:
		return gorm.NewGormStore(cfg)
	case config.StorageTypeRedis:
		return redis.NewRedisStore(cfg)
	default:
		return nil, config.ErrBadStorageType
	}
}
