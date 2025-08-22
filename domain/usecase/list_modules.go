package usecase

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type ListModulesUsecase interface {
	List(ctx context.Context, pagination *common.Pagination) ([]*entity.Module, int64, error)
}

type ListModulesImpl struct {
	moduleRepo repository.ModuleRepository
}

func NewListModulesImpl(moduleRepo repository.ModuleRepository) ListModulesUsecase {
	return &ListModulesImpl{
		moduleRepo: moduleRepo,
	}
}

func (uc *ListModulesImpl) List(ctx context.Context, pagination *common.Pagination) ([]*entity.Module, int64, error) {
	return uc.moduleRepo.GetWithPagination(ctx, pagination)
}
