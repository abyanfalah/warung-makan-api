package repository_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"
	"warung-makan/model"
	"warung-makan/repository"
	"warung-makan/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyTransaction = []model.TransactionTest{
	{
		Id:         "dummy id 1",
		TotalPrice: 8888,
		Created_at: "creation date",
		Updated_at: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	},
}

type TransactionRepositoryTestSuite struct {
	suite.Suite
	mockDb     *sql.DB
	mockSql    sqlmock.Sqlmock
	mockSqlxDb *sqlx.DB
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	db, sql, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	suite.mockDb = db
	suite.mockSql = sql
	suite.mockSqlxDb = sqlx.NewDb(suite.mockDb, "postgres")
}

func (suite *TransactionRepositoryTestSuite) TearDownTest() {
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

func (suite *TransactionRepositoryTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "total_price", "created_at", "updated_at"})
	for _, dummy := range dummyTransaction {
		rows.AddRow(dummy.Id, dummy.TotalPrice, dummy.Created_at, dummy.Updated_at.Time)
	}

	suite.mockSql.ExpectQuery(utils.TRANSACTION_GET_ALL).WillReturnRows(rows)

	repo := repository.NewTransactionRepository(suite.mockSqlxDb)
	actual, err := repo.GetAllTest()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(actual))
	assert.Equal(suite.T(), "dummy id 1", actual[0].Id)
}

func (suite *TransactionRepositoryTestSuite) TestGetAll_Failed() {
	rows := sqlmock.NewRows([]string{"id", "total_price", "created_at", "updated_at"})
	for _, dummy := range dummyTransaction {
		rows.AddRow(dummy.Id, dummy.TotalPrice, dummy.Created_at, dummy.Updated_at.Time)
	}

	suite.mockSql.ExpectQuery(utils.TRANSACTION_GET_ALL).WillReturnError(errors.New("failed"))

	repo := repository.NewTransactionRepository(suite.mockSqlxDb)
	actual, err := repo.GetAll()

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite *TransactionRepositoryTestSuite) TestGetById_Success() {
	dummy := dummyTransaction[0]
	row := sqlmock.NewRows([]string{"id", "total_price", "created_at", "updated_at"})
	row.AddRow(dummy.Id, dummy.TotalPrice, dummy.Created_at, dummy.Updated_at.Time)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.TRANSACTION_GET_BY_ID)).WithArgs(dummy.Id).WillReturnRows(row)

	repo := repository.NewTransactionRepository(suite.mockSqlxDb)
	actual, err := repo.GetByIdTest(dummy.Id)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), dummyTransaction[0], actual)
}

func (suite *TransactionRepositoryTestSuite) TestGetById_Failed() {
	dummy := dummyTransaction[0]
	row := sqlmock.NewRows([]string{"id", "total_price", "created_at", "updated_at"})
	row.AddRow(dummy.Id, dummy.TotalPrice, dummy.Created_at, dummy.Updated_at.Time)

	suite.mockSql.ExpectQuery(utils.TRANSACTION_GET_BY_ID).WillReturnError(errors.New("failed"))

	repo := repository.NewTransactionRepository(suite.mockSqlxDb)
	actual, err := repo.GetByIdTest(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.TransactionTest{}, actual)
}

func (suite *TransactionRepositoryTestSuite) TestInsert_Success() {
	var dummy = dummyTransaction[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.TRANSACTION_INSERT_TEST)).WithArgs(dummy.Id, dummy.TotalPrice).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewTransactionRepository(suite.mockSqlxDb)
	actual, err := repo.InsertTest(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, actual)
}

func (suite *TransactionRepositoryTestSuite) TestInsert_Failed() {
	var dummy = dummyTransaction[0]
	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.TRANSACTION_INSERT_TEST)).WillReturnError(errors.New("failed"))

	repo := repository.NewTransactionRepository(suite.mockSqlxDb)
	actual, err := repo.InsertTest(&dummy)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.TransactionTest{}, actual)
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
