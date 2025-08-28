package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/boostrap/config"
)

type dbCache struct {
	Addr        string
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Network     string
}

type Env struct {
	NODE_ENV string
	URL_DB   string

	NAME_SERVICE   string
	PORT_GRPC      int
	HOST_GRPC      string
	INTERVAL_CHECK string
	TIMEOUT_CHECK  string

	DB_CACHE *dbCache
}

func NewEnv(env any) {
	sc := config.DefaultSettingsConfig()
	if sc.IsProduction() {
		sc.SetFile("prod.config")
	} else {
		sc.SetFile("dev.config")
	}
	config.NewConfig(sc, env)
}

func (env *Env) IsProduction() bool {
	return strings.ToLower(env.NODE_ENV) == "production"
}
