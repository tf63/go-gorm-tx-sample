package infrastracture

import "github.com/tf63/go-gorm-tx-sample/internal/uow-pattern/domain"

type repositoryManager struct {
	ar domain.AccountRepository
}

func (r *repositoryManager) AccountRepository() domain.AccountRepository {
	return r.ar
}

func NewRepositoryManager(ar domain.AccountRepository) domain.RepositoryManager {
	return &repositoryManager{ar: ar}
}
