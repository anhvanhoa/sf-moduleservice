package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type ModuleUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Module, error)
}

type ModuleUsecaseImpl struct {
	moduleRepo repository.ModuleRepository
}

func NewModuleUsecaseImpl(moduleRepo repository.ModuleRepository) ModuleUsecase {
	return &ModuleUsecaseImpl{
		moduleRepo: moduleRepo,
	}
}

func (uc *ModuleUsecaseImpl) GetAll(ctx context.Context) ([]*entity.Module, error) {
	return uc.moduleRepo.GetAll(ctx)
}
