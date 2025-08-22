package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type UpdateModuleUsecase interface {
	Update(ctx context.Context, module *entity.Module) error
}

type UpdateModuleImpl struct {
	moduleRepo repository.ModuleRepository
}

func NewUpdateModuleImpl(moduleRepo repository.ModuleRepository) UpdateModuleUsecase {
	return &UpdateModuleImpl{
		moduleRepo: moduleRepo,
	}
}

func (uc *UpdateModuleImpl) Update(ctx context.Context, module *entity.Module) error {
	return uc.moduleRepo.Update(ctx, module)
}
