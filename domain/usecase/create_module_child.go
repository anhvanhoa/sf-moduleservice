package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreateModuleChildUsecase interface {
	CreateModuleChild(ctx context.Context, moduleChild *entity.ModuleChild) error
}

type CreateModuleChildImpl struct {
	moduleChildRepo repository.ModuleChildRepository
}

func NewCreateModuleChildImpl(moduleChildRepo repository.ModuleChildRepository) CreateModuleChildUsecase {
	return &CreateModuleChildImpl{
		moduleChildRepo: moduleChildRepo,
	}
}

func (uc *CreateModuleChildImpl) CreateModuleChild(ctx context.Context, moduleChild *entity.ModuleChild) error {
	return uc.moduleChildRepo.Create(ctx, moduleChild)
}
