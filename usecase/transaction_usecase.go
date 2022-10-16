package usecase

import (
	"warung-makan/model"
	"warung-makan/repository"
)

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUsecase interface {
	GetAll() ([]model.Transaction, error)
	GetAllPaginated(page int, rows int) ([]model.Transaction, error)
	GetById(id string) (model.Transaction, error)
	// GetByName(name string) ([]model.Transaction, error)
	Insert(transaction *model.Transaction) (model.Transaction, error)
	Update(transaction *model.Transaction) (model.Transaction, error)
	Delete(id string) error
}

func (p *transactionUsecase) GetAll() ([]model.Transaction, error) {
	return p.transactionRepository.GetAll()
}

func (p *transactionUsecase) GetAllPaginated(page int, rows int) ([]model.Transaction, error) {
	return p.transactionRepository.GetAllPaginated(page, rows)
}

func (p *transactionUsecase) GetById(id string) (model.Transaction, error) {
	return p.transactionRepository.GetById(id)
}

func (p *transactionUsecase) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	return p.transactionRepository.Insert(newTransaction)
}

func (p *transactionUsecase) Update(newTransaction *model.Transaction) (model.Transaction, error) {
	return p.transactionRepository.Update(newTransaction)
}

func (p *transactionUsecase) Delete(id string) error {
	return p.transactionRepository.Delete(id)
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	usecase := new(transactionUsecase)
	usecase.transactionRepository = transactionRepository
	return usecase
}
