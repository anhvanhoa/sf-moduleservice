package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/boostrap/config"
)

type dbCache struct {
	Addr        string
	Db          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Network     string
}

type Env struct {
	NodeEnv string
	UrlDb   string

	NameService   string
	PortGrpc      int
	HostGrpc      string
	IntervalCheck string
	TimeoutCheck  string

	DbCache *dbCache
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
	return strings.ToLower(env.NodeEnv) == "production"
}
