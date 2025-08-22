package modulechildservice

import (
	"module-service/bootstrap"
	"module-service/domain/usecase"
	"module-service/infrastructure/repo"

	proto "module-service/proto/gen/module/v1"

	"github.com/go-pg/pg/v10"
)

type moduleChildService struct {
	proto.UnsafeModuleChildServiceServer
	createChildUc  usecase.CreateModuleChildUsecase
	getChildUc     usecase.GetModuleChildUsecase
	listChildrenUc usecase.ListModuleChildrenUsecase
	updateChildUc  usecase.UpdateModuleChildUsecase
	deleteChildUc  usecase.DeleteModuleChildUsecase
}

func NewModuleChildService(db *pg.DB, env *bootstrap.Env) proto.ModuleChildServiceServer {

	moduleChildRepo := repo.NewModuleChildRepository(db)

	return &moduleChildService{
		createChildUc:  usecase.NewCreateModuleChildImpl(moduleChildRepo),
		getChildUc:     usecase.NewGetModuleChildImpl(moduleChildRepo),
		listChildrenUc: usecase.NewListModuleChildrenImpl(moduleChildRepo),
		updateChildUc:  usecase.NewUpdateModuleChildImpl(moduleChildRepo),
		deleteChildUc:  usecase.NewDeleteModuleChildImpl(moduleChildRepo),
	}
}
