package bootstrap

import (
	"github.com/anhvanhoa/service-core/boostrap/db"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env *Env
	DB  *pg.DB
	Log *log.LogGRPCImpl
}

func App() *Application {
	env := Env{}
	NewEnv(&env)

	logConfig := log.NewConfig()
	log := log.InitLogGRPC(logConfig, zapcore.DebugLevel, env.IsProduction())

	db := db.NewPostgresDB(db.ConfigDB{
		URL:  env.URL_DB,
		Mode: env.NODE_ENV,
	})

	return &Application{
		Env: &env,
		DB:  db,
		Log: log,
	}
}
