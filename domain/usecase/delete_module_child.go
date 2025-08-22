package usecase

import (
	"context"
	"module-service/domain/repository"
)

type DeleteModuleChildUsecase interface {
	DeleteByID(ctx context.Context, id string) error
}

type DeleteModuleChildImpl struct {
	moduleChildRepo repository.ModuleChildRepository
}

func NewDeleteModuleChildImpl(moduleChildRepo repository.ModuleChildRepository) DeleteModuleChildUsecase {
	return &DeleteModuleChildImpl{
		moduleChildRepo: moduleChildRepo,
	}
}

func (uc *DeleteModuleChildImpl) DeleteByID(ctx context.Context, id string) error {
	return uc.moduleChildRepo.DeleteByID(ctx, id)
}
