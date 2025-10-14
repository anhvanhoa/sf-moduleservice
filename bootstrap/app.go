package bootstrap

import (
	"module-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/bootstrap/db"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env   *Env
	DB    *pg.DB
	Log   *log.LogGRPCImpl
	Repos repo.Repositories
}

func App() *Application {
	env := Env{}
	NewEnv(&env)

	logConfig := log.NewConfig()
	log := log.InitLogGRPC(logConfig, zapcore.DebugLevel, env.IsProduction())

	db := db.NewPostgresDB(db.ConfigDB{
		URL:  env.UrlDb,
		Mode: env.NodeEnv,
	})
	helper := utils.NewHelper()
	repos := repo.NewRepositories(db, helper)

	return &Application{
		Env:   &env,
		DB:    db,
		Log:   log,
		Repos: repos,
	}
}
