package usecase_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"
	"warung-makan/model"
	"warung-makan/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMenus = []model.Transaction{
	{
		Id:         "dummy id 1",
		TotalPrice: 113,
		Created_at: "today",
		Updated_at: sql.NullTime{Time: time.Now(), Valid: true},
	},
	{
		Id:         "dummy id 2",
		TotalPrice: 113,
		Created_at: "today",
		Updated_at: sql.NullTime{Time: time.Now(), Valid: true},
	},
}

type repoMock struct {
	mock.Mock
}

type TransactionUsecaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (r *repoMock) GetAll() ([]model.Transaction, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Transaction), nil
}

func (r *repoMock) GetById(id string) (model.Transaction, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}
	return args.Get(0).(model.Transaction), nil
}

func (r *repoMock) Insert(menu *model.Transaction) (model.Transaction, error) {
	args := r.Called(menu)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}
	return args.Get(0).(model.Transaction), nil
}

func (r *repoMock) GetByIdTest(id string) (model.TransactionTest, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.TransactionTest{}, args.Error(1)
	}
	return args.Get(0).(model.TransactionTest), nil
}

func (r *repoMock) InsertTest(menu *model.TransactionTest) (model.TransactionTest, error) {
	args := r.Called(menu)
	if args.Get(1) != nil {
		return model.TransactionTest{}, args.Error(1)
	}
	return args.Get(0).(model.TransactionTest), nil
}

func (suite *TransactionUsecaseTestSuite) TestTransactionGetAll_Success() {
	suite.repoMock.On("GetAll").Return(dummyMenus, nil)

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.repoMock)
	menus, err := TransactionUsecaseTest.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenus, menus)
}

func (suite *TransactionUsecaseTestSuite) TestTransactionGetAll_Failed() {
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed"))

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.repoMock)
	menus, err := TransactionUsecaseTest.GetAll()

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(menus))
}

func (suite *TransactionUsecaseTestSuite) TestTransactionGetById_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("GetById", dummy.Id).Return(dummy, nil)

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.repoMock)
	menu, err := TransactionUsecaseTest.GetById(dummy.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, menu)
}

func (suite *TransactionUsecaseTestSuite) TestTransactionGetById_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("GetById", dummy.Id).Return(model.Transaction{}, errors.New("failed"))

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.repoMock)
	menu, err := TransactionUsecaseTest.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Transaction{}, menu)
}

func (suite *TransactionUsecaseTestSuite) TestTransactionInsert_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Insert", &dummy).Return(dummy, nil)

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.repoMock)
	menu, err := TransactionUsecaseTest.Insert(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, menu)

}

func (suite *TransactionUsecaseTestSuite) TestTransactionInsert_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Insert", &dummy).Return(model.Transaction{}, errors.New("failed"))

	TransactionUsecaseTest := usecase.NewTransactionUsecase(suite.repoMock)
	menu, err := TransactionUsecaseTest.Insert(&dummy)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Transaction{}, menu)

}

func (suite *TransactionUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUsecaseTestSuite))
}
