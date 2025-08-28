package moduleservice

import (
	"module-service/domain/usecase"
	"module-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/domain/goid"
	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"github.com/go-pg/pg/v10"
)

type moduleService struct {
	proto_module.UnsafeModuleServiceServer
	createUc usecase.CreateModuleUsecase
	getUc    usecase.GetModuleUsecase
	listUc   usecase.ListModulesUsecase
	updateUc usecase.UpdateModuleUsecase
	deleteUc usecase.DeleteModuleUsecase
}

func NewModuleService(db *pg.DB) proto_module.ModuleServiceServer {
	moduleRepo := repo.NewModuleRepository(db)
	uuid := goid.NewGoId().UUID()
	return &moduleService{
		createUc: usecase.NewCreateModule(moduleRepo, uuid),
		getUc:    usecase.NewGetModuleImpl(moduleRepo),
		listUc:   usecase.NewListModulesImpl(moduleRepo),
		updateUc: usecase.NewUpdateModuleImpl(moduleRepo),
		deleteUc: usecase.NewDeleteModuleImpl(moduleRepo),
	}
}
