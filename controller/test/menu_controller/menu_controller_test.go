package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"warung-makan/controller"
	"warung-makan/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMenus = []model.Menu{
	{
		Id:    "dummy id 1",
		Name:  "dummy name 1",
		Price: 8888,
		Stock: 99,
		Image: "dummy image path 1",
	},
	{
		Id:    "dummy id 2",
		Name:  "dummy name 2",
		Price: 8888,
		Stock: 99,
		Image: "dummy image path 2",
	},
}

type ErrorResponse struct {
	Error   string
	Message string
}

type MenuUsecaseMock struct {
	mock.Mock
}

// Buat TestSuite
type MenuControllerTestSuite struct {
	suite.Suite
	useCaseMock *MenuUsecaseMock
	routerMock  *gin.Engine
}

func (suite *MenuControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.useCaseMock = new(MenuUsecaseMock)
}

func (r *MenuUsecaseMock) GetAll() ([]model.Menu, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}

func (r *MenuUsecaseMock) GetById(id string) (model.Menu, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *MenuUsecaseMock) GetByName(name string) ([]model.Menu, error) {
	args := r.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}

func (r *MenuUsecaseMock) GetByCredentials(menuname, password string) (model.Menu, error) {
	args := r.Called(menuname, password)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *MenuUsecaseMock) Insert(menu *model.Menu) (model.Menu, error) {
	args := r.Called(menu)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *MenuUsecaseMock) Update(newUser *model.Menu) (model.Menu, error) {
	args := r.Called(newUser)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *MenuUsecaseMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite MenuControllerTestSuite) TestGetAllUserApi_Success() {
	menus := dummyMenus
	suite.useCaseMock.On("GetAll").Return(menus, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/menu", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUsers []model.Menu
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualUsers)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 2, len(actualUsers))
	assert.Equal(suite.T(), menus[0].Id, actualUsers[0].Id)
}

func (suite MenuControllerTestSuite) TestGetAllUserApi_Failed() {
	suite.useCaseMock.On("GetAll").Return(nil, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/menu", nil)
	suite.routerMock.ServeHTTP(r, request)

	var errorResponse ErrorResponse
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &errorResponse)

	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)
	assert.Equal(suite.T(), "failed", errorResponse.Error)
	assert.Nil(suite.T(), jsonerr)

}

func (suite MenuControllerTestSuite) TestGetByIdUserApi_Success() {
	menu := dummyMenus[0]
	suite.useCaseMock.On("GetById", menu.Id).Return(menu, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/menu/"+menu.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser model.Menu
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), menu.Id, actualUser.Id)
}

func (suite MenuControllerTestSuite) TestGetByIdUserApi_Failed() {
	menu := dummyMenus[0]
	suite.useCaseMock.On("GetById", menu.Id).Return(model.Menu{}, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/menu/"+menu.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser model.Menu
	response := r.Body.String()

	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.NotEqual(suite.T(), menu, actualUser)
}

func (suite MenuControllerTestSuite) TestGetByNameUserApi_Success() {
	menus := dummyMenus
	suite.useCaseMock.On("GetByName", "dummy").Return(menus, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/menu?name=dummy", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUsers []model.Menu
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualUsers)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 2, len(actualUsers))
	assert.Equal(suite.T(), menus[0].Id, actualUsers[0].Id)
}

func (suite MenuControllerTestSuite) TestGetByNameUserApi_Failed() {
	suite.useCaseMock.On("GetByName", "dummy").Return(nil, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/menu?name=dummy", nil)
	suite.routerMock.ServeHTTP(r, request)

	var errorResponse ErrorResponse
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &errorResponse)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), "failed", errorResponse.Error)
	assert.Nil(suite.T(), jsonerr)
}

func TestMenuControllerTestSuite(t *testing.T) {
	suite.Run(t, new(MenuControllerTestSuite))
}
