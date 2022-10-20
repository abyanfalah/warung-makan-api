package repository

import (
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type transactionDetailRepository struct {
	db *sqlx.DB
}

type TransactionDetailRepository interface {
	// GetAll() ([]model.TransactionDetail, error)
	GetByTrasactionId(id string) ([]model.TransactionDetail, error)
	// GetByName(name string) ([]model.TransactionDetail, error)

	Insert(transactionDetail *model.TransactionDetail) (model.TransactionDetail, error)
}

// func (p *transactionDetailRepository) GetAll() ([]model.TransactionDetail, error) {
// 	var transactionDetails []model.TransactionDetail
// 	err := p.db.Select(&transactionDetails, utils.TRANSACTION_GET_ALL+" order by id")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transactionDetails, nil
// }

// func (p *transactionDetailRepository) GetAllPaginated(page int, rows int) ([]model.TransactionDetail, error) {
// 	var transactionDetails []model.TransactionDetail
// 	limit := rows
// 	offset := limit * (page - 1)

// 	err := p.db.Select(&transactionDetails, utils.TRANSACTION_GET_ALL_PAGINATED, limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transactionDetails, nil
// }

// func (p *transactionDetailRepository) GetById(id string) (model.TransactionDetail, error) {
// 	var transactionDetail model.TransactionDetail
// 	err := p.db.Get(&transactionDetail, utils.TRANSACTION_GET_BY_ID, id)
// 	if err != nil {
// 		return model.TransactionDetail{}, err
// 	}
// 	return transactionDetail, nil
// }

func (p *transactionDetailRepository) GetByTrasactionId(id string) ([]model.TransactionDetail, error) {
	var transactionDetails []model.TransactionDetail
	err := p.db.Select(&transactionDetails, utils.TRANSACTION_DETAIL_GET_BY_ID_TRANSACTION, id)
	if err != nil {
		return nil, err
	}
	return transactionDetails, nil
}

func (p *transactionDetailRepository) Insert(newTransactionDetail *model.TransactionDetail) (model.TransactionDetail, error) {

	_, err := p.db.NamedExec(utils.TRANSACTION_DETAIL_INSERT, newTransactionDetail)
	if err != nil {
		return model.TransactionDetail{}, err
	}
	transactionDetail := newTransactionDetail
	return *transactionDetail, nil
}

// func (p *transactionDetailRepository) Update(newData *model.TransactionDetail) (model.TransactionDetail, error) {
// 	_, err := p.db.NamedExec(utils.TRANSACTION_UPDATE, newData)
// 	if err != nil {
// 		return model.TransactionDetail{}, err
// 	}
// 	return *newData, nil
// }

// func (p *transactionDetailRepository) Delete(id string) error {
// 	_, err := p.db.Exec(utils.TRANSACTION_DELETE, id)
// 	return err
// }

func NewTransactionDetailRepository(db *sqlx.DB) TransactionDetailRepository {
	repo := new(transactionDetailRepository)
	repo.db = db
	return repo
}
