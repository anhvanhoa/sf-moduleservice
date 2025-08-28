package modulechildservice

import (
	"module-service/domain/usecase"
	"module-service/infrastructure/repo"

	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"

	"github.com/go-pg/pg/v10"
)

type moduleChildService struct {
	proto_module_child.UnsafeModuleChildServiceServer
	createChildUc  usecase.CreateModuleChildUsecase
	getChildUc     usecase.GetModuleChildUsecase
	listChildrenUc usecase.ListModuleChildrenUsecase
	updateChildUc  usecase.UpdateModuleChildUsecase
	deleteChildUc  usecase.DeleteModuleChildUsecase
}

func NewModuleChildService(db *pg.DB) proto_module_child.ModuleChildServiceServer {
	moduleChildRepo := repo.NewModuleChildRepository(db)
	return &moduleChildService{
		createChildUc:  usecase.NewCreateModuleChildImpl(moduleChildRepo),
		getChildUc:     usecase.NewGetModuleChildImpl(moduleChildRepo),
		listChildrenUc: usecase.NewListModuleChildrenImpl(moduleChildRepo),
		updateChildUc:  usecase.NewUpdateModuleChildImpl(moduleChildRepo),
		deleteChildUc:  usecase.NewDeleteModuleChildImpl(moduleChildRepo),
	}
}
