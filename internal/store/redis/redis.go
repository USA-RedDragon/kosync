package redis

import (
	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/store/models"
)

type Redis struct {
}

func NewRedisStore(config *config.Config) (*Redis, error) {
	return &Redis{}, nil
}

func (s *Redis) CreateUser(user *models.User) error {
	return nil
}

func (s *Redis) GetUser(id string) (*models.User, error) {
	return nil, nil
}

func (s *Redis) GetProgress(book string) (*models.Progress, error) {
	return nil, nil
}

func (s *Redis) UpdateProgress(book string, progress *models.Progress) error {
	return nil
}
