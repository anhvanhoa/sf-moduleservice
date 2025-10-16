package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/bootstrap/config"
	"github.com/anhvanhoa/service-core/domain/grpc_client"
)

type dbCache struct {
	Addr        string `mapstructure:"addr"`
	Db          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxActive   int    `mapstructure:"max_active"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
	Network     string `mapstructure:"network"`
}

type Env struct {
	NodeEnv           string                    `mapstructure:"node_env"`
	UrlDb             string                    `mapstructure:"url_db"`
	NameService       string                    `mapstructure:"name_service"`
	PortGprc          int                       `mapstructure:"port_grpc"`
	HostGprc          string                    `mapstructure:"host_grpc"`
	IntervalCheck     string                    `mapstructure:"interval_check"`
	TimeoutCheck      string                    `mapstructure:"timeout_check"`
	DbCache           *dbCache                  `mapstructure:"db_cache"`
	GrpcClients       []*grpc_client.ConfigGrpc `mapstructure:"grpc_clients"`
	AddressPermission string                    `mapstructure:"address_permission"`
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
