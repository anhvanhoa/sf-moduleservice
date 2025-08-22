package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type ModuleChildUsecase interface {
	GetAll(ctx context.Context) ([]*entity.ModuleChild, error)
}

type ModuleChildUsecaseImpl struct {
	moduleChildRepo repository.ModuleChildRepository
}

func NewModuleChildUsecaseImpl(moduleChildRepo repository.ModuleChildRepository) ModuleChildUsecase {
	return &ModuleChildUsecaseImpl{
		moduleChildRepo: moduleChildRepo,
	}
}

func (uc *ModuleChildUsecaseImpl) GetAll(ctx context.Context) ([]*entity.ModuleChild, error) {
	return uc.moduleChildRepo.GetAll(ctx)
}
