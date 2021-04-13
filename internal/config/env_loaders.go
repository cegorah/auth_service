package config

import (
	"os"
	"strconv"
)

func (apiCfg *ApiConfig) loadConfig(prefix string) error {
	apiString := os.Getenv(prefix + "_API_STRING")
	if apiString == "" {
		apiString = "localhost:8080"
	}
	apiCfg.ConnectionString = apiString
	return nil
}


func (rdsCfg *RedisConfig) loadConfig(prefix string) error {
	addr := os.Getenv(prefix + "_REDIS_STRING")
	if addr == "" {
		addr = "localhost:6379"
	}
	username := os.Getenv(prefix + "_REDIS_USERNAME")
	password := os.Getenv(prefix + "_REDIS_PASSWORD")
	dbS := os.Getenv(prefix + "_REDIS_DATABASE")
	if dbS == "" {
		dbS = "0"
	}
	dbI, err := strconv.Atoi(dbS)
	if err != nil {
		return err
	}
	rdsCfg.ConnectionString = addr
	rdsCfg.Username = username
	rdsCfg.Password = password
	rdsCfg.Database = dbI
	return nil
}

func (dbCfg *DBConfig) loadConfig(prefix string) error {
	dbprovider := os.Getenv(prefix + "_DB_PROVIDER")
	if dbprovider == "" {
		dbprovider = "postgres"
	}
	addr := os.Getenv(prefix + "_DB_STRING")
	if addr == "" {
		addr = "postgres://user:pass@localhost/db"
	}
	dbCfg.ConnectionString = addr
	dbCfg.Provider = dbprovider
	return nil
}
