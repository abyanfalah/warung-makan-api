package repository_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"warung-makan/model"
	"warung-makan/repository"
	"warung-makan/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyDetail = []model.TransactionDetail{
	{
		TransactionId: "dummy tr 1",
		MenuId:        "dummy menu 1",
		Qty:           5,
		Subtotal:      88888,
	},
}

type TransactionDetailRepositoryTestSuite struct {
	suite.Suite
	mockDb     *sql.DB
	mockSql    sqlmock.Sqlmock
	mockSqlxDb *sqlx.DB
}

func (suite *TransactionDetailRepositoryTestSuite) SetupTest() {
	db, sql, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	suite.mockDb = db
	suite.mockSql = sql
	suite.mockSqlxDb = sqlx.NewDb(suite.mockDb, "postgres")
}

func (suite *TransactionDetailRepositoryTestSuite) TearDownTest() {
	// var err error

	// err = suite.mockDb.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// err = suite.mockSqlxDb.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// Gak perlu teardown. Mungkin karena sql.DB gak dipake. Yang dipake hanya sqlx.DB
}

func (suite *TransactionDetailRepositoryTestSuite) TestGetByTransactionId_Success() {
	dummy := dummyDetail[0]
	rows := sqlmock.NewRows([]string{"transaction_id", "menu_id", "qty", "subtotal"})
	for _, dummy := range dummyDetail {
		rows.AddRow(dummy.TransactionId, dummy.MenuId, dummy.Qty, dummy.Subtotal)
	}

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.TRANSACTION_DETAIL_GET_BY_ID_TRANSACTION)).WillReturnRows(rows)

	repo := repository.NewTransactionDetailRepository(suite.mockSqlxDb)
	actual, err := repo.GetByTrasactionId(dummy.TransactionId)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(actual))
	assert.Equal(suite.T(), dummy.TransactionId, actual[0].TransactionId)
}
func (suite *TransactionDetailRepositoryTestSuite) TestGetByTransactionId_Failed() {
	dummy := dummyDetail[0]

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.TRANSACTION_DETAIL_GET_BY_ID_TRANSACTION)).WillReturnError(errors.New("failed"))

	repo := repository.NewTransactionDetailRepository(suite.mockSqlxDb)
	actual, err := repo.GetByTrasactionId(dummy.TransactionId)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(actual))
}

func (suite *TransactionDetailRepositoryTestSuite) TestInsert_Success() {
	var dummy = dummyDetail[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.TRANSACTION_DETAIL_INSERT_TEST)).WithArgs(dummy.TransactionId, dummy.MenuId, dummy.Qty, dummy.Subtotal).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewTransactionDetailRepository(suite.mockSqlxDb)
	actual, err := repo.Insert(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, actual)
}
func (suite *TransactionDetailRepositoryTestSuite) TestInsert_Failed() {
	var dummy = dummyDetail[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.TRANSACTION_DETAIL_INSERT_TEST)).WithArgs(dummy.TransactionId, dummy.MenuId, dummy.Qty, dummy.Subtotal).WillReturnError(errors.New("failed"))

	repo := repository.NewTransactionDetailRepository(suite.mockSqlxDb)
	actual, err := repo.Insert(&dummy)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.TransactionDetail{}, actual)
}

func TestTransactionDetailRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDetailRepositoryTestSuite))
}
