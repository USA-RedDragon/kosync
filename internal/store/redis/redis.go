package redis

import (
	"fmt"

	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/store/models"
)

type Redis struct {
}

func NewRedisStore(config *config.Config) (*Redis, error) {
	return &Redis{}, fmt.Errorf("not implemented")
}

func (s *Redis) CreateUser(username, password string) error {
	return fmt.Errorf("not implemented")
}

func (s *Redis) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, fmt.Errorf("not implemented")
}

func (s *Redis) GetProgress(username, document string) (models.Progress, error) {
	return models.Progress{}, fmt.Errorf("not implemented")
}

func (s *Redis) UpdateProgress(progress models.Progress) error {
	return fmt.Errorf("not implemented")
}
