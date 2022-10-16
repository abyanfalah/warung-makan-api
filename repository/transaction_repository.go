package repository

import (
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type transactionRepository struct {
	db *sqlx.DB
}

type TransactionRepository interface {
	GetAllPaginated(page int, rows int) ([]model.Transaction, error)
	GetAll() ([]model.Transaction, error)
	GetById(id string) (model.Transaction, error)
	// GetByName(name string) ([]model.Transaction, error)

	Insert(transaction *model.Transaction) (model.Transaction, error)
	Update(transaction *model.Transaction) (model.Transaction, error)
	Delete(id string) error
}

func (p *transactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction

	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL+" order by created_at desc")
	if err != nil {
		return nil, err
	}

	tdRepo := NewTransactionDetailRepository(p.db)

	for i, transaction := range transactions {
		items, err := tdRepo.GetByTrasactionId(transaction.Id)
		if err != nil {
			panic(err)
		}
		transactions[i].Items = items
	}

	return transactions, nil
}

func (p *transactionRepository) GetAllPaginated(page int, rows int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	limit := rows
	offset := limit * (page - 1)

	err := p.db.Select(&transactions, utils.TRANSACTION_GET_ALL_PAGINATED, limit, offset)
	if err != nil {
		return nil, err
	}

	tdRepo := NewTransactionDetailRepository(p.db)

	for i, transaction := range transactions {
		items, err := tdRepo.GetByTrasactionId(transaction.Id)
		if err != nil {
			panic(err)
		}
		transactions[i].Items = items
	}

	return transactions, nil
}

func (p *transactionRepository) GetById(id string) (model.Transaction, error) {
	var transaction model.Transaction
	err := p.db.Get(&transaction, utils.TRANSACTION_GET_BY_ID, id)
	if err != nil {
		return model.Transaction{}, err
	}

	tdRepo := NewTransactionDetailRepository(p.db)

	items, err := tdRepo.GetByTrasactionId(transaction.Id)
	if err != nil {
		panic(err)
	}
	transaction.Items = items

	return transaction, nil
}

func (p *transactionRepository) Insert(newTransaction *model.Transaction) (model.Transaction, error) {
	// ===================================
	tx, err := p.db.Beginx()
	if err != nil {
		panic(err)
	}
	_, err = tx.NamedExec(utils.TRANSACTION_INSERT, newTransaction)
	if err != nil {
		panic("gagal insert transaksi")
	}

	for _, each := range newTransaction.Items {
		_, err = tx.NamedExec(utils.TRANSACTION_DETAIL_INSERT, each)
		if err != nil {
			panic(err)
		}

		_, err = tx.NamedExec(utils.MENU_UPDATE_STOCK, each)
		if err != nil {
			panic("gagal update stok")
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	// ===================================

	if err != nil {
		return model.Transaction{}, err
	}

	return *newTransaction, nil
}

func (p *transactionRepository) Update(newData *model.Transaction) (model.Transaction, error) {
	_, err := p.db.NamedExec(utils.TRANSACTION_UPDATE, newData)
	if err != nil {
		return model.Transaction{}, err
	}
	return *newData, nil
}

func (p *transactionRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.TRANSACTION_DELETE, id)
	return err
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
