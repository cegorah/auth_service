package config

import (
	"encoding/json"
	"github.com/cegorah/auth_service/internal"
	"os"
	"testing"
)

var (
	filePath = os.Getenv("CONFIG_PATH") + "/config.json"
	prefix = os.Getenv("CONFIG_PREFIX")
	aps    = prefix + "_API_STRING"
	dbs    = prefix + "_DB_STRING"
	dbpr   = prefix + "_DB_PROVIDER"
	rds    = prefix + "_REDIS_STRING"
	rdun   = prefix + "_REDIS_USERNAME"
	rdpwd  = prefix + "_REDIS_PASSWORD"
	jwts   = prefix + "_JWT_SECRET"
)

type TestCase struct {
	envValues map[string]string
	fail      bool
}

var Tests = []TestCase{
	{
		envValues: internal.DefaultEnvs, fail: false,
	},
	{
		envValues: map[string]string{
			rds:   "localhost",
			rdun:  "admin",
			rdpwd: "password",
			aps:   "localhost:8080",
			dbs:   "postgres://admin:admin@localhost/admin",
			dbpr:  "psql",
			jwts:  "",
		}, fail: true,
	},
}

func TestConfig_FromEnv(t *testing.T) {
	for _, testCase := range Tests {
		setEnv(testCase.envValues)
		cfg, err := FromEnv(prefix)
		if err != nil {
			if e, ok := err.(*JwtEmptyError); ok && testCase.fail {
				t.Log(e)
				continue
			}
			t.Fatal(err)
		}
		checkEnv(t, testCase.envValues, cfg)
	}
}

func TestConfig_FromFile(t *testing.T) {
	plainCfg := Config{}
	cfg, err := FromFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	fl, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}
	dc := json.NewDecoder(fl)
	err = dc.Decode(&plainCfg)
	if err != nil {
		t.Fatal(err)
	}
	internal.TestEqual(t, plainCfg.ApiCfg.ConnectionString, cfg.ApiCfg.ConnectionString)
}

func TestConfig_Extend(t *testing.T) {
	fileCfg, err := FromFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	setEnv(Tests[0].envValues)
	envCfg, err := FromEnv(prefix)
	if err != nil {
		t.Fatal(err)
	}
	envCfg.ExtendFromFile(&fileCfg)
	internal.TestEqual(t, fileCfg.ApiCfg.ReadTimeoutSecond, envCfg.ApiCfg.ReadTimeoutSecond)
	internal.TestEqual(t, fileCfg.ApiCfg.WriteTimeoutSecond, envCfg.ApiCfg.WriteTimeoutSecond)
}

func setEnv(envValues map[string]string) {
	for k, el := range envValues {
		_ = os.Setenv(k, el)
	}
}
func checkEnv(t *testing.T, envValues map[string]string, config Config) {
	tb := testing.TB(t)
	rdCfg := config.RedisCfg
	dbCfg := config.DBCfg
	apiCfg := config.ApiCfg
	internal.TestEqual(tb, config.JWTSecret, envValues[jwts])
	internal.TestEqual(tb, rdCfg.ConnectionString, envValues[rds])
	internal.TestEqual(tb, rdCfg.Username, envValues[rdun])
	internal.TestEqual(tb, rdCfg.Password, envValues[rdpwd])
	internal.TestEqual(tb, dbCfg.ConnectionString, envValues[dbs])
	internal.TestEqual(tb, dbCfg.Provider, envValues[dbpr])
	internal.TestEqual(tb, apiCfg.ConnectionString, envValues[aps])
}
