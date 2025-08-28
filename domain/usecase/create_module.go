package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/domain/goid"
)

type CreateModuleUsecase interface {
	CreateModule(ctx context.Context, module *entity.Module) error
}

type CreateModuleImpl struct {
	moduleRepo repository.ModuleRepository
	goid       goid.GoUUID
}

func NewCreateModule(moduleRepo repository.ModuleRepository, goid goid.GoUUID) CreateModuleUsecase {
	return &CreateModuleImpl{
		moduleRepo: moduleRepo,
		goid:       goid,
	}
}

func (uc *CreateModuleImpl) CreateModule(ctx context.Context, module *entity.Module) error {
	module.ID = uc.goid.Gen()
	return uc.moduleRepo.Create(ctx, module)
}
