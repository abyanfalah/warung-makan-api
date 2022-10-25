package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"warung-makan/config"
	"warung-makan/controller"
	"warung-makan/model"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var auth = authenticator.NewAccessToken(config.NewConfig().TokenConfig)
var token, _ = auth.GenerateAccessToken(&model.User{
	Username: "admin",
	Password: "admin",
})

var dummyTransactions = []model.Transaction{
	{
		Id:         "dummy transaction 1",
		TotalPrice: 15000,
		// Created_at: time.Now().Format("YYYY-MM-DD"),
		Items: []model.TransactionDetail{
			{
				TransactionId: "dummy transaction 1",
				MenuId:        "menu 1",
				Qty:           10,
				Subtotal:      10000,
			},
			{
				TransactionId: "dummy transaction 1",
				MenuId:        "menu 2",
				Qty:           5,
				Subtotal:      5000,
			},
		},
	},
}

type itemList struct {
	Items []model.TransactionDetail
}

// var dummyItemList = []model.ItemList{
// 	{
// 		MenuId: "menu 1",
// 		Qty:    10,
// 	},
// 	{
// 		MenuId: "menu 2",
// 		Qty:    5,
// 	},
// }

var itemlist = itemList{
	Items: dummyTransactions[0].Items,
}

type ErrorResponse struct {
	Error   string
	Message string
}

type TransactionUsecaseMock struct {
	mock.Mock
}

// Buat TestSuite
type TransactionControllerTestSuite struct {
	suite.Suite
	useCaseMock            *TransactionUsecaseMock
	transactionUsecaseMock *MenuUsecaseMock
	routerMock             *gin.Engine
}

func (suite *TransactionControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.useCaseMock = new(TransactionUsecaseMock)
}

func (r *TransactionUsecaseMock) GetAll() ([]model.Transaction, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Transaction), nil
}

func (r *TransactionUsecaseMock) GetById(id string) (model.Transaction, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}
	return args.Get(0).(model.Transaction), nil
}

func (r *TransactionUsecaseMock) Insert(user *model.Transaction) (model.Transaction, error) {
	args := r.Called(user)
	if args.Get(1) != nil {
		return model.Transaction{}, args.Error(1)
	}
	return args.Get(0).(model.Transaction), nil
}

func (suite TransactionControllerTestSuite) TestGetAllTransactionApi_Success() {
	transaction := dummyTransactions[0]
	suite.useCaseMock.On("GetAll").Return(dummyTransactions, nil)

	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualTransaction []model.Transaction
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualTransaction)

	fmt.Println("err: ", err)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 1, len(actualTransaction))
	assert.Equal(suite.T(), transaction.Id, actualTransaction[0].Id)
}

func (suite TransactionControllerTestSuite) TestGetAllTransactionApi_Failed() {
	suite.useCaseMock.On("GetAll").Return(nil, errors.New("failed"))

	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/transaction", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var errorResponse ErrorResponse
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &errorResponse)

	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)
	assert.Equal(suite.T(), "failed", errorResponse.Error)
	assert.Nil(suite.T(), jsonerr)

}

func (suite TransactionControllerTestSuite) TestGetByIdTransactionApi_Success() {
	transaction := dummyTransactions[0]
	suite.useCaseMock.On("GetById", transaction.Id).Return(transaction, nil)

	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/transaction/"+transaction.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualTransaction model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualTransaction)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), transaction.Id, actualTransaction.Id)
}

func (suite TransactionControllerTestSuite) TestGetByIdTransactionApi_Failed() {
	transaction := dummyTransactions[0]
	suite.useCaseMock.On("GetById", transaction.Id).Return(model.Transaction{}, errors.New("failed"))

	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/transaction/"+transaction.Id, nil)

	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualTransaction model.Transaction
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualTransaction)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), model.Transaction{}, actualTransaction)
}

func (suite TransactionControllerTestSuite) TestInsertTransactionApi_Success() {
	transaction := dummyTransactions[0]

	suite.useCaseMock.On("Insert", &transaction).Return(transaction, nil)

	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(transaction.Items)
	request, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualTransaction = model.Transaction{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualTransaction.Items)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), transaction, actualTransaction)
	return

}

// func (suite TransactionControllerTestSuite) TestInsertTransactionApi_FailedBinding() {
// 	suite.useCaseMock.On("Insert").Return(model.Menu{}, errors.New("failed"))

// 	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

// 	r := httptest.NewRecorder()

// 	request, _ := http.NewRequest(http.MethodPost, "/transaction/no_image", nil)
// 	request.Header.Add("Authorization", "Bearer "+token)
// 	suite.routerMock.ServeHTTP(r, request)

// 	var actualMenu = model.Menu{}
// 	response := r.Body.String()
// 	json.Unmarshal([]byte(response), &actualMenu)

// 	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
// 	assert.Equal(suite.T(), model.Menu{}, actualMenu)
// }

// func (suite TransactionControllerTestSuite) TestInsertTransactionApi_Failed() {
// 	transaction := dummyTransactions[0]
// 	suite.useCaseMock.On("Insert", &transaction).Return(model.Menu{}, errors.New("failed"))

// 	controller.NewTransactionController(suite.useCaseMock, suite.transactionUsecaseMock, suite.routerMock)

// 	r := httptest.NewRecorder()

// 	reqBody, _ := json.Marshal(transaction)
// 	request, _ := http.NewRequest(http.MethodPost, "/transaction/no_image", bytes.NewBuffer(reqBody))
// 	request.Header.Add("Authorization", "Bearer "+token)
// 	suite.routerMock.ServeHTTP(r, request)

// 	var actualMenu = model.Menu{}
// 	response := r.Body.String()
// 	json.Unmarshal([]byte(response), &actualMenu)

// 	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)
// }

func TestMenuControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionControllerTestSuite))
}
