package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type UpdateModuleChildUsecase interface {
	Update(ctx context.Context, moduleChild *entity.ModuleChild) error
}

type UpdateModuleChildImpl struct {
	moduleChildRepo repository.ModuleChildRepository
}

func NewUpdateModuleChildImpl(moduleChildRepo repository.ModuleChildRepository) UpdateModuleChildUsecase {
	return &UpdateModuleChildImpl{
		moduleChildRepo: moduleChildRepo,
	}
}

func (uc *UpdateModuleChildImpl) Update(ctx context.Context, moduleChild *entity.ModuleChild) error {
	return uc.moduleChildRepo.Update(ctx, moduleChild)
}
