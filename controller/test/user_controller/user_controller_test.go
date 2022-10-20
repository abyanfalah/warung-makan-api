package controller_test

import (
	"warung-makan/controller"
	"warung-makan/model"
	"warung-makan/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyUsers = []model.User{
	{
		Id:       "dummy id 1",
		Name:     "dummy name 1",
		Username: "dummy username 1",
		Image:    "dummy image path 1",
	},
	{
		Id:       "dummy id 2",
		Name:     "dummy name 2",
		Username: "dummy username 2",
		Image:    "dummy image path 2",
	},
}

type UserUsecaseMock struct {
	mock.Mock
}

// Buat TestSuite
type UserControllerTestSuite struct {
	suite.Suite
	routerMock  *gin.Engine
	ucManMock   *UsecaseManagerMock
	useCaseMock *UserUsecaseMock
}

func (um *UserControllerTestSuite) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(um.repo.UserRepo())
}

func (um *UserControllerTestSuite) MenuUsecase() usecase.MenuUsecase {
	return usecase.NewMenuUsecase(um.repo.MenuRepo())
}

func (um *UserControllerTestSuite) TransactionUsecase() usecase.TransactionUsecase {
	return usecase.NewTransactionUsecase(um.repo.TransactionRepo())
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
	suite.useCaseMock.On("GetAll").Return(dummyUsers, nil)

	controller.NewUserController(suite.ucManMock, suite.routerMock)
}
