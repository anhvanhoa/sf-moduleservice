package usecase

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type GetModuleChildUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.ModuleChild, error)
}

type GetModuleChildImpl struct {
	moduleChildRepo repository.ModuleChildRepository
}

func NewGetModuleChildImpl(moduleChildRepo repository.ModuleChildRepository) GetModuleChildUsecase {
	return &GetModuleChildImpl{
		moduleChildRepo: moduleChildRepo,
	}
}

func (uc *GetModuleChildImpl) GetByID(ctx context.Context, id string) (*entity.ModuleChild, error) {
	return uc.moduleChildRepo.GetByID(ctx, id)
}
