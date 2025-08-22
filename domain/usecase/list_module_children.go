package usecase

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type ListModuleChildrenUsecase interface {
	List(ctx context.Context, pagination *common.Pagination, moduleID string) ([]*entity.ModuleChild, int64, error)
}

type ListModuleChildrenImpl struct {
	moduleChildRepo repository.ModuleChildRepository
}

func NewListModuleChildrenImpl(moduleChildRepo repository.ModuleChildRepository) ListModuleChildrenUsecase {
	return &ListModuleChildrenImpl{
		moduleChildRepo: moduleChildRepo,
	}
}

func (uc *ListModuleChildrenImpl) List(ctx context.Context, pagination *common.Pagination, moduleID string) ([]*entity.ModuleChild, int64, error) {
	// For now, we'll get all module children and filter by moduleID
	// In a real implementation, you might want to modify the repository to support filtering by moduleID
	moduleChildren, total, err := uc.moduleChildRepo.GetWithPagination(ctx, pagination)
	if err != nil {
		return nil, 0, err
	}
	
	// Filter by moduleID if provided
	if moduleID != "" {
		filtered := make([]*entity.ModuleChild, 0)
		for _, child := range moduleChildren {
			if child.ModuleID == moduleID {
				filtered = append(filtered, child)
			}
		}
		return filtered, int64(len(filtered)), nil
	}
	
	return moduleChildren, total, nil
}
