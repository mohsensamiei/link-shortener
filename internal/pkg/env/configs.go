package env

import (
	"fmt"
	parser "github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"time"
)

type Configs struct {
	Mode string `env:"GIN_MODE" envDefault:"debug"`

	PostgresHost string `env:"POSTGRES_HOST"`
	PostgresUser string `env:"POSTGRES_USER"`
	PostgresPass string `env:"POSTGRES_PASSWORD"`

	RedisHost string `env:"REDIS_HOST"`
	RedisPass string `env:"REDIS_PASSWORD"`

	AuthenticatePostgresDB    string        `env:"AUTHENTICATE_POSTGRES_DB" envDefault:"authenticate"`
	AuthenticateTokenKey      string        `env:"AUTHENTICATE_TOKEN_KEY" envDefault:"secret"`
	AuthenticateTokenDuration time.Duration `env:"AUTHENTICATE_TOKEN_DURATION" envDefault:"24h"`
	AuthenticateHttpHost      string        `env:"AUTHENTICATE_HTTP_HOST" envDefault:"localhost:2080"`
	AuthenticateGrpcHost      string        `env:"AUTHENTICATE_GRPC_HOST" envDefault:"localhost:2135"`
}

func (configs *Configs) LoadEnv(filenames ...string) error {
	_ = godotenv.Load(filenames...)
	if err := parser.Parse(configs); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (configs Configs) IsDebugMode() bool {
	return configs.Mode == gin.DebugMode
}
func (configs Configs) PostgresUrl(db string) string {
	return fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", configs.PostgresUser, configs.PostgresPass, configs.PostgresHost, db)
}
func (configs Configs) ConnectPostgres(db string) (*gorm.DB, error) {
	client, err := gorm.Open("postgres", configs.PostgresUrl(db))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return client, nil
}
func (configs Configs) ConnectRedis(db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost,
		Password: configs.RedisPass,
		DB:       db,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.WithStack(err)
	}
	return client, nil
}
