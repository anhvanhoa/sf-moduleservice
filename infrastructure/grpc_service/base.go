package grpcservice

import (
	"github.com/anhvanhoa/module-service/bootstrap"
	// "github.com/anhvanhoa/module-service/infrastructure/repo"
	proto "github.com/anhvanhoa/module-service/proto/gen/exam/v1"

	"github.com/go-pg/pg/v10"
)

type examService struct {
	proto.UnsafeExamServiceServer
}

func NewExamService(db *pg.DB, env *bootstrap.Env) proto.ExamServiceServer {
	// tx := repo.NewManagerTransaction(db)
	// configRedis := bootstrap.NewRedisConfig(
	// 	env.DB_CACHE.Addr,
	// 	env.DB_CACHE.Password,
	// 	env.DB_CACHE.DB,
	// 	env.DB_CACHE.Network,
	// 	env.DB_CACHE.MaxIdle,
	// 	env.DB_CACHE.MaxActive,
	// 	env.DB_CACHE.IdleTimeout,
	// )
	// cacheService := bootstrap.NewRedis(configRedis)

	return &examService{}
}
