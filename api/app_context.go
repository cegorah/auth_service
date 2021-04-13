package api

import (
	"database/sql"
	"github.com/cegorah/auth_service/internal/config"
	"github.com/cegorah/auth_service/pkg/auth_code"
	"log"
)

type LocalContext struct {
	db            *sql.DB
	codeValidator *auth_code.CodeValidator
	cache         *auth_code.Cacher
}

func (ctx *LocalContext) initContext(config *config.Config) {
	if err := ctx.initDB(&config.DBCfg); err != nil {
		log.Fatal(err)
	}
	if err := ctx.initBroker(&config.RedisCfg); err != nil {
		log.Fatal(err)
	}
	if err := ctx.initCodeValidator(); err != nil {
		log.Fatal(err)
	}
}

func (ctx *LocalContext) initDB(cfg *config.DBConfig) error {
	return nil
}

func (ctx *LocalContext) initBroker(cfg *config.RedisConfig) error {
	return nil
}

func (ctx *LocalContext) initCodeValidator() error {
	return nil
}
