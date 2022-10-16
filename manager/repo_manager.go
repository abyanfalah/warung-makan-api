package manager

import "warung-makan/repository"

type repoManager struct {
	infra InfraManager
}

type RepoManager interface {
	UserRepo() repository.UserRepository
	MenuRepo() repository.MenuRepository
	TransactionRepo() repository.TransactionRepository
	TransactionDetailRepo() repository.TransactionDetailRepository
}

func (rm *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(rm.infra.GetSqlDb())
}

func (rm *repoManager) MenuRepo() repository.MenuRepository {
	return repository.NewMenuRepository(rm.infra.GetSqlDb())
}

func (rm *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(rm.infra.GetSqlDb())
}

func (rm *repoManager) TransactionDetailRepo() repository.TransactionDetailRepository {
	return repository.NewTransactionDetailRepository(rm.infra.GetSqlDb())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
