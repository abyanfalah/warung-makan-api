package controller_test

import (
	"bytes"
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

var dummyUsers = []model.User{
	{
		Id:       "dummy id 1",
		Name:     "dummy name 1",
		Username: "dummy username 1",
		Password: "dummy password 1",
		Image:    "dummy image path 1",
	},
	{
		Id:       "dummy id 2",
		Name:     "dummy name 2",
		Username: "dummy username 2",
		Password: "dummy password 2",
		Image:    "dummy image path 2",
	},
}

type ErrorResponse struct {
	Error   string
	Message string
}

type LoginResponse struct {
	message string
	token   string
}

type UserUsecaseMock struct {
	mock.Mock
}

// Buat TestSuite
type UserControllerTestSuite struct {
	suite.Suite
	useCaseMock *UserUsecaseMock
	routerMock  *gin.Engine
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.useCaseMock = new(UserUsecaseMock)
}

func (r *UserUsecaseMock) GetAll() ([]model.User, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), nil
}

func (r *UserUsecaseMock) GetById(id string) (model.User, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *UserUsecaseMock) GetByName(name string) ([]model.User, error) {
	args := r.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), nil
}

func (r *UserUsecaseMock) GetByCredentials(username, password string) (model.User, error) {
	args := r.Called(username, password)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *UserUsecaseMock) Insert(user *model.User) (model.User, error) {
	args := r.Called(user)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *UserUsecaseMock) Update(newUser *model.User) (model.User, error) {
	args := r.Called(newUser)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *UserUsecaseMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite UserControllerTestSuite) TestGetAllUserApi_Success() {
	users := dummyUsers
	suite.useCaseMock.On("GetAll").Return(users, nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/user", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUsers []model.User
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualUsers)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 2, len(actualUsers))
	assert.Equal(suite.T(), users[0].Id, actualUsers[0].Id)
}

func (suite UserControllerTestSuite) TestGetAllUserApi_Failed() {
	suite.useCaseMock.On("GetAll").Return(nil, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user", nil)
	suite.routerMock.ServeHTTP(r, request)

	var errorResponse ErrorResponse
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &errorResponse)

	assert.Equal(suite.T(), http.StatusBadGateway, r.Code)
	assert.Equal(suite.T(), "failed", errorResponse.Error)
	assert.Nil(suite.T(), jsonerr)

}

func (suite UserControllerTestSuite) TestGetByIdUserApi_Success() {
	user := dummyUsers[0]
	suite.useCaseMock.On("GetById", user.Id).Return(user, nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/user/"+user.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser model.User
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), user.Id, actualUser.Id)
}

func (suite UserControllerTestSuite) TestGetByIdUserApi_Failed() {
	user := dummyUsers[0]
	suite.useCaseMock.On("GetById", user.Id).Return(model.User{}, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user/"+user.Id, nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser model.User
	response := r.Body.String()

	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.NotEqual(suite.T(), user, actualUser)
}

func (suite UserControllerTestSuite) TestGetByNameUserApi_Success() {
	users := dummyUsers
	suite.useCaseMock.On("GetByName", "dummy").Return(users, nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/user?name=dummy", nil)
	suite.routerMock.ServeHTTP(r, request)

	var actualUsers []model.User
	response := r.Body.String()

	jsonerr := json.Unmarshal([]byte(response), &actualUsers)

	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), 2, len(actualUsers))
	assert.Equal(suite.T(), users[0].Id, actualUsers[0].Id)
}

func (suite UserControllerTestSuite) TestGetByNameUserApi_Failed() {
	suite.useCaseMock.On("GetByName", "dummy").Return(nil, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user?name=dummy", nil)
	suite.routerMock.ServeHTTP(r, request)

	var errorResponse ErrorResponse
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &errorResponse)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), "failed", errorResponse.Error)
	assert.Nil(suite.T(), jsonerr)
}

func (suite UserControllerTestSuite) TestLoginUserApi_Success() {
	user := dummyUsers[0]

	suite.useCaseMock.On("GetByCredentials", user.Username, user.Password).Return(user, nil)

	controller.NewLoginController(suite.useCaseMock, suite.routerMock)
	var credentials = model.Credential{
		Username: user.Username,
		Password: user.Password,
	}
	body, _ := json.Marshal(credentials)

	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/test/login", bytes.NewReader(body))

	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualUser)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.Id, actualUser.Id)
}

func (suite UserControllerTestSuite) TestLoginUserApi_Failed() {
	user := dummyUsers[0]

	suite.useCaseMock.On("GetByCredentials", user.Username, user.Password).Return(model.User{}, errors.New("failed"))

	controller.NewLoginController(suite.useCaseMock, suite.routerMock)
	var credentials = model.Credential{
		Username: user.Username,
		Password: user.Password,
	}
	body, _ := json.Marshal(credentials)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/test/login", bytes.NewReader(body))

	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()

	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), model.User{}, actualUser)
}

// func (suite UserControllerTestSuite) TestGetByIdUserApi_Failed() {
// 	user := dummyUsers[0]
// 	suite.useCaseMock.On("GetById", user.Id).Return(model.User{}, errors.New("failed"))

// 	controller.NewUserController(suite.useCaseMock, suite.routerMock)

// 	r := httptest.NewRecorder()
// 	request, _ := http.NewRequest(http.MethodGet, "/user/"+user.Id, nil)
// 	suite.routerMock.ServeHTTP(r, request)

// 	var actualUser model.User
// 	response := r.Body.String()

// 	json.Unmarshal([]byte(response), &actualUser)

// 	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
// 	assert.NotEqual(suite.T(), user, actualUser)
// }

func (suite UserControllerTestSuite) TestInsertUserApi_Success() {
	user := dummyUsers[0]

	suite.useCaseMock.On("Insert", user).Return(user, nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	// formData := struct{

	// }
	body, _ := json.Marshal(user)
	request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualUser)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	return
	assert.Equal(suite.T(), user.Id, actualUser.Id)

}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
