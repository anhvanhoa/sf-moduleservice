package bootstrap

import (
	pkglog "module-service/infrastructure/service/logger"

	"github.com/go-pg/pg/v10"
	valid "github.com/go-playground/validator/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env   *Env
	DB    *pg.DB
	Log   pkglog.Logger
	Valid IValidator
}

func App() *Application {
	env := Env{}
	NewEnv(&env)

	logConfig := pkglog.NewConfig()
	log := pkglog.InitLogger(logConfig, zapcore.DebugLevel, env.IsProduction())

	db := NewPostgresDB(&env, log)
	valid := RegisterCustomValidations(valid.New())
	return &Application{
		Env:   &env,
		DB:    db,
		Log:   log,
		Valid: valid,
	}
}
