/*
Config package handle many different configs for other packages.
*/
package config

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	CodeLength         int `json:"code_length"`
	CodeExpirationTime int `json:"code_expiration_time"`
	JWTSecret          string
	ApiCfg             ApiConfig   `json:"api_config"`
	DBCfg              DBConfig    `json:"db_config"`
	RedisCfg           RedisConfig `json:"redis_config"`
}

type DBConfig struct {
	MaxOpenCons      int    `json:"max_open_cons"`
	MaxIdleCons      int    `json:"max_idle_cons"`
	ConnectionString string `json:"connection_string"`
	Provider         string `json:"provider"`
}

type RedisConfig struct {
	Database         int    `json:"database"`
	ConnectionString string `json:"connection_string"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	TlsConfig        *tls.Config
}

type ApiConfig struct {
	WriteTimeoutSecond int    `json:"write_timeout_second"`
	ReadTimeoutSecond  int    `json:"read_timeout_second"`
	ConnectionString   string `json:"connection_string"`
}

type Configer interface {
	loadConfig(string) error
}

func FromEnv(prefix string) (Config, error) {
	cfgTypes := map[string]Configer{
		"api":   new(ApiConfig),
		"db":    new(DBConfig),
		"redis": new(RedisConfig),
	}
	mainConfig := Config{}
	for name, cfg := range cfgTypes {
		if err := cfg.loadConfig(prefix); err != nil {
			return Config{}, err
		}
		switch name {
		case "api":
			{
				apiCfg, ok := cfg.(*ApiConfig)
				if !ok {
					return Config{}, fmt.Errorf("can't cast %s", name)
				}
				mainConfig.ApiCfg = *apiCfg
			}
		case "db":
			{
				dbCfg, ok := cfg.(*DBConfig)
				if !ok {
					return Config{}, fmt.Errorf("can't cast %s", name)
				}
				mainConfig.DBCfg = *dbCfg
			}
		case "redis":
			{
				redisCfg, ok := cfg.(*RedisConfig)
				if !ok {
					return Config{}, fmt.Errorf("can't cast %s", name)
				}
				mainConfig.RedisCfg = *redisCfg

			}
		}
	}
	jwtSecret := os.Getenv(prefix + "_JWT_SECRET")
	if jwtSecret == "" {
		return Config{}, &JwtEmptyError{"JwtSecret should be provided by envs"}
	}
	mainConfig.JWTSecret = jwtSecret
	return mainConfig, nil
}

func FromFile(configPath string) (Config, error) {
	cfg := Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (c *Config) ExtendFromFile(fVars *Config) {
	c.ApiCfg.WriteTimeoutSecond = fVars.ApiCfg.WriteTimeoutSecond
	c.ApiCfg.ReadTimeoutSecond = fVars.ApiCfg.ReadTimeoutSecond
	c.DBCfg.MaxIdleCons = fVars.DBCfg.MaxIdleCons
	c.DBCfg.MaxOpenCons = fVars.DBCfg.MaxOpenCons
}
