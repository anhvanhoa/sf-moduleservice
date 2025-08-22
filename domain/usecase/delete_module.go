package usecase

import (
	"context"
	"module-service/domain/repository"
)

type DeleteModuleUsecase interface {
	DeleteByID(ctx context.Context, id string) error
}

type DeleteModuleImpl struct {
	moduleRepo repository.ModuleRepository
}

func NewDeleteModuleImpl(moduleRepo repository.ModuleRepository) DeleteModuleUsecase {
	return &DeleteModuleImpl{
		moduleRepo: moduleRepo,
	}
}

func (uc *DeleteModuleImpl) DeleteByID(ctx context.Context, id string) error {
	return uc.moduleRepo.DeleteByID(ctx, id)
}
