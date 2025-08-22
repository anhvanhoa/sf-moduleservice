package repo

import (
	"module-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type RepositoryFactory struct {
	ModuleRepository      repository.ModuleRepository
	ModuleChildRepository repository.ModuleChildRepository
	ManagerTransaction    repository.ManagerTransaction
}

func NewRepositoryFactory(db *pg.DB) *RepositoryFactory {
	return &RepositoryFactory{
		ModuleRepository:      NewModuleRepository(db),
		ModuleChildRepository: NewModuleChildRepository(db),
		ManagerTransaction:    NewManagerTransaction(db),
	}
}

func (rf *RepositoryFactory) GetModuleRepository() repository.ModuleRepository {
	return rf.ModuleRepository
}

func (rf *RepositoryFactory) GetModuleChildRepository() repository.ModuleChildRepository {
	return rf.ModuleChildRepository
}

func (rf *RepositoryFactory) GetManagerTransaction() repository.ManagerTransaction {
	return rf.ManagerTransaction
}
