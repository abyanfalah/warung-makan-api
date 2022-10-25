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

var dummyUsers = []model.User{
	{
		Id:       "dummy id 1",
		Name:     "dummy name 1",
		Username: "dummyusername1",
		Password: "dummypassword1",
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

var auth = authenticator.NewAccessToken(config.NewConfig().TokenConfig)
var token, _ = auth.GenerateAccessToken(&model.User{
	Username: "admin",
	Password: "admin",
})

type ErrorResponse struct {
	Error   string
	Message string
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

	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
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

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
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

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
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

func (suite UserControllerTestSuite) TestInsertUserNoImageApi_Success() {
	user := dummyUsers[0]

	suite.useCaseMock.On("Insert", &user).Return(user, nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(user)
	request, err := http.NewRequest(http.MethodPost, "/user/no_image", bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualUser)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.Name, actualUser.Name)
	assert.Equal(suite.T(), user.Username, actualUser.Username)
	assert.Equal(suite.T(), user.Password, actualUser.Password)

}

func (suite UserControllerTestSuite) TestInsertUserNoImageApi_FailedBinding() {
	suite.useCaseMock.On("Insert").Return(model.Menu{}, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodPost, "/user/no_image", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), model.User{}, actualUser)
}

func (suite UserControllerTestSuite) TestInsertUserNoImageApi_Failed() {
	user := dummyUsers[0]
	suite.useCaseMock.On("Insert", &user).Return(model.Menu{}, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/user/no_image", bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
}

func (suite UserControllerTestSuite) TestUpdateUserApi_Success() {
	user := dummyUsers[0]

	suite.useCaseMock.On("Update", &user).Return(user, nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(user)
	request, err := http.NewRequest(http.MethodPut, "/user/"+user.Id, bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	jsonerr := json.Unmarshal([]byte(response), &actualUser)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), jsonerr)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), user.Name, actualUser.Name)
	assert.Equal(suite.T(), user.Username, actualUser.Username)
	assert.Equal(suite.T(), user.Password, actualUser.Password)

}

func (suite UserControllerTestSuite) TestUpdateUserApi_FailedBindingAndNoId() {
	suite.useCaseMock.On("Update").Return(model.Menu{}, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodPut, "/user/asdf", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.Equal(suite.T(), model.User{}, actualUser)

	r = httptest.NewRecorder()
	request, _ = http.NewRequest(http.MethodPut, "/user/", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	response = r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
	assert.Equal(suite.T(), model.User{}, actualUser)
}

func (suite UserControllerTestSuite) TestUpdateUserApi_Failed() {
	user := dummyUsers[0]

	suite.useCaseMock.On("Update", &user).Return(model.Menu{}, errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()

	reqBody, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPut, "/user/"+user.Id, bytes.NewBuffer(reqBody))
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	var actualUser = model.User{}
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualUser)

	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	assert.Equal(suite.T(), model.User{}, actualUser)
}

func (suite UserControllerTestSuite) TestDeleteUserApi_Success() {
	user := dummyUsers[0]

	suite.useCaseMock.On("GetById", user.Id).Return(user, nil)
	suite.useCaseMock.On("Delete", user.Id).Return(nil)

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/user/"+user.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusOK, r.Code)

}

func (suite UserControllerTestSuite) TestDeleteUserApi_FailedNotFound() {
	user := dummyUsers[0]

	suite.useCaseMock.On("GetById", user.Id).Return(model.User{}, errors.New("not found"))
	suite.useCaseMock.On("Delete", user.Id).Return(errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/user/"+user.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusNotFound, r.Code)

}

func (suite UserControllerTestSuite) TestDeleteUserApi_Failed() {
	user := dummyUsers[0]

	suite.useCaseMock.On("GetById", user.Id).Return(user, nil)
	suite.useCaseMock.On("Delete", user.Id).Return(errors.New("failed"))

	controller.NewUserController(suite.useCaseMock, suite.routerMock)

	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/user/"+user.Id, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	suite.routerMock.ServeHTTP(r, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)

}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
