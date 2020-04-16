package env

import (
	"fmt"
	parser "github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Configs struct {
	Mode string `env:"GIN_MODE" envDefault:"debug"`

	PostgresHost string `env:"POSTGRES_HOST"`
	PostgresUser string `env:"POSTGRES_USER"`
	PostgresPass string `env:"POSTGRES_PASSWORD"`
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
func (configs Configs) PostgresUrl(dbname string) string {
	return fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", configs.PostgresUser, configs.PostgresPass, configs.PostgresHost, dbname)
}
func (configs Configs) ConnectPostgres(dbname string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", configs.PostgresUrl(dbname))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
