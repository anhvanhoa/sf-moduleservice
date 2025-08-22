package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type GetModuleUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.Module, error)
}

type GetModuleImpl struct {
	moduleRepo repository.ModuleRepository
}

func NewGetModuleImpl(moduleRepo repository.ModuleRepository) GetModuleUsecase {
	return &GetModuleImpl{
		moduleRepo: moduleRepo,
	}
}

func (uc *GetModuleImpl) GetByID(ctx context.Context, id string) (*entity.Module, error) {
	return uc.moduleRepo.GetByID(ctx, id)
}
