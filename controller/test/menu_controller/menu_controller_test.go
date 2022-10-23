package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
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

var auth = authenticator.NewAccessToken(config.NewConfig().TokenConfig)
var token, _ = auth.GenerateAccessToken(&model.User{
	Username: "admin",
	Password: "admin",
})

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

func (r *MenuUsecaseMock) Update(newMenu *model.Menu) (model.Menu, error) {
	args := r.Called(newMenu)
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

func (suite MenuControllerTestSuite) TestGetAllMenuApi_Success() {
	menus := dummyMenus
	suite.useCaseMock.On("GetAll").Return(menus, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/menu", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenus []model.Menu
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualMenus)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 2, len(actualMenus))
	assert.Equal(suite.T(), menus[0].Id, actualMenus[0].Id)
}

func (suite MenuControllerTestSuite) TestGetAllMenuApi_Failed() {
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

func (suite MenuControllerTestSuite) TestGetByIdMenuApi_Success() {
	menu := dummyMenus[0]
	suite.useCaseMock.On("GetById", menu.Id).Return(menu, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/menu/"+menu.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu model.Menu
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), menu.Id, actualMenu.Id)
}

func (suite MenuControllerTestSuite) TestGetByIdMenuApi_Failed() {
	menu := dummyMenus[0]
	suite.useCaseMock.On("GetById", menu.Id).Return(model.Menu{}, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/menu/"+menu.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu model.Menu
	response := r.Body.String()

	json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.NotEqual(suite.T(), menu, actualMenu)
}

func (suite MenuControllerTestSuite) TestGetByNameMenuApi_Success() {
	menus := dummyMenus
	suite.useCaseMock.On("GetByName", "dummy").Return(menus, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/menu?name=dummy", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenus []model.Menu
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualMenus)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 2, len(actualMenus))
	assert.Equal(suite.T(), menus[0].Id, actualMenus[0].Id)
}

func (suite MenuControllerTestSuite) TestGetByNameMenuApi_Failed() {
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

func (suite MenuControllerTestSuite) TestInsertMenuNoImageApi_Success() {
	menu := dummyMenus[0]

	suite.useCaseMock.On("Insert", &menu).Return(menu, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(menu)
	request, err := http.NewRequest(http.MethodPost, "/menu/no_image", bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu = model.Menu{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualMenu)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), menu.Name, actualMenu.Name)

}

func (suite MenuControllerTestSuite) TestInsertMenuNoImageApi_FailedBinding() {
	suite.useCaseMock.On("Insert").Return(model.Menu{}, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodPost, "/menu/no_image", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu = model.Menu{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), model.Menu{}, actualMenu)
}

func (suite MenuControllerTestSuite) TestInsertMenuNoImageApi_Failed() {
	menu := dummyMenus[0]
	suite.useCaseMock.On("Insert", &menu).Return(model.Menu{}, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(menu)
	request, _ := http.NewRequest(http.MethodPost, "/menu/no_image", bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu = model.Menu{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)
}

func (suite MenuControllerTestSuite) TestUpdateMenuApi_Success() {
	menu := dummyMenus[0]

	suite.useCaseMock.On("Update", &menu).Return(menu, nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(menu)
	request, err := http.NewRequest(http.MethodPut, "/menu/"+menu.Id, bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu = model.Menu{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualMenu)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), menu.Name, actualMenu.Name)
}

func (suite MenuControllerTestSuite) TestUpdateMenuApi_FailedBindingAndNoId() {
	suite.useCaseMock.On("Update").Return(model.Menu{}, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodPut, "/menu/asdf", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu = model.Menu{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), model.Menu{}, actualMenu)

	r = httptest.NewRecorder()
	request, _ = http.NewRequest(http.MethodPut, "/menu/", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	response = r.Body.String()
	json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
	assert.Equal(suite.T(), model.Menu{}, actualMenu)
}

func (suite MenuControllerTestSuite) TestUpdateMenuApi_Failed() {
	menu := dummyMenus[0]

	suite.useCaseMock.On("Update", &menu).Return(model.Menu{}, errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(menu)
	request, _ := http.NewRequest(http.MethodPut, "/menu/"+menu.Id, bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualMenu = model.Menu{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualMenu)

	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)
	assert.Equal(suite.T(), model.Menu{}, actualMenu)
}

func (suite MenuControllerTestSuite) TestDeleteMenuApi_Success() {
	menu := dummyMenus[0]

	suite.useCaseMock.On("GetById", menu.Id).Return(menu, nil)
	suite.useCaseMock.On("Delete", menu.Id).Return(nil)

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/menu/"+menu.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusOK, r.Code)

}

func (suite MenuControllerTestSuite) TestDeleteMenuApi_FailedNotFound() {
	menu := dummyMenus[0]

	suite.useCaseMock.On("GetById", menu.Id).Return(model.Menu{}, errors.New("not found"))
	suite.useCaseMock.On("Delete", menu.Id).Return(errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/menu/"+menu.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)

}

func (suite MenuControllerTestSuite) TestDeleteMenuApi_Failed() {
	menu := dummyMenus[0]

	suite.useCaseMock.On("GetById", menu.Id).Return(menu, nil)
	suite.useCaseMock.On("Delete", menu.Id).Return(errors.New("failed"))

	controller.NewMenuController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/menu/"+menu.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)

}

func TestMenuControllerTestSuite(t *testing.T) {
	suite.Run(t, new(MenuControllerTestSuite))
}
