package moduleservice

import (
	"module-service/bootstrap"
	"module-service/domain/usecase"
	"module-service/infrastructure/repo"
	goid "module-service/infrastructure/service/goid"

	proto "module-service/proto/gen/module/v1"

	"github.com/go-pg/pg/v10"
)

type moduleService struct {
	proto.UnsafeModuleServiceServer
	createUc usecase.CreateModuleUsecase
	getUc    usecase.GetModuleUsecase
	listUc   usecase.ListModulesUsecase
	updateUc usecase.UpdateModuleUsecase
	deleteUc usecase.DeleteModuleUsecase
}

func NewModuleService(db *pg.DB, env *bootstrap.Env) proto.ModuleServiceServer {

	moduleRepo := repo.NewModuleRepository(db)

	return &moduleService{
		createUc: usecase.NewCreateModule(moduleRepo, goid.NewGoId()),
		getUc:    usecase.NewGetModuleImpl(moduleRepo),
		listUc:   usecase.NewListModulesImpl(moduleRepo),
		updateUc: usecase.NewUpdateModuleImpl(moduleRepo),
		deleteUc: usecase.NewDeleteModuleImpl(moduleRepo),
	}
}
