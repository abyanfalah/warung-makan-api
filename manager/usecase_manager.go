package manager

import "warung-makan/usecase"

type usecaseManager struct {
	repo RepoManager
}

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
	MenuUsecase() usecase.MenuUsecase
	TransactionUsecase() usecase.TransactionUsecase
	TransactionDetailUsecase() usecase.TransactionDetailUsecase
}

func (um *usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(um.repo.UserRepo())
}

func (um *usecaseManager) MenuUsecase() usecase.MenuUsecase {
	return usecase.NewMenuUsecase(um.repo.MenuRepo())
}

func (um *usecaseManager) TransactionUsecase() usecase.TransactionUsecase {
	return usecase.NewTransactionUsecase(um.repo.TransactionRepo())
}

func (um *usecaseManager) TransactionDetailUsecase() usecase.TransactionDetailUsecase {
	return usecase.NewTransactionDetailUsecase(um.repo.TransactionDetailRepo())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
