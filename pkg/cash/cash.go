package cash

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"time"
)

type ManagerInstance struct{}
type Manager interface {
	Set(key string, value interface{}, duration time.Duration) error
	Get(key string, value interface{}) error
}

func NewManager(db *redis.Client) Manager {
	return &manager{
		DB: db,
	}
}

type manager struct {
	DB *redis.Client
}

func (m manager) Set(key string, value interface{}, duration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := m.DB.Set(key, bytes, duration).Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (m manager) Get(key string, value interface{}) error {
	bytes, err := m.DB.Get(key).Bytes()
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(bytes, value); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
