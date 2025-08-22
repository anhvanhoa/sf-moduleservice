package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
	"module-service/domain/service/goid"
)

type CreateModuleUsecase interface {
	CreateModule(ctx context.Context, module *entity.Module) error
}

type CreateModuleImpl struct {
	moduleRepo repository.ModuleRepository
	goid       goid.GoId
}

func NewCreateModule(moduleRepo repository.ModuleRepository, goid goid.GoId) CreateModuleUsecase {
	return &CreateModuleImpl{
		moduleRepo: moduleRepo,
		goid:       goid,
	}
}

func (uc *CreateModuleImpl) CreateModule(ctx context.Context, module *entity.Module) error {
	module.ID = uc.goid.NewUUID()
	return uc.moduleRepo.Create(ctx, module)
}
