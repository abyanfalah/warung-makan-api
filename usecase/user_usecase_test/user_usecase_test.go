package usecase_test

import (
	"errors"
	"testing"
	"warung-makan/model"
	"warung-makan/usecase"

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

type repoMock struct {
	mock.Mock
}

type UserUsecaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (r *repoMock) GetAll() ([]model.User, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), nil
}

func (r *repoMock) GetById(id string) (model.User, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *repoMock) GetByName(name string) ([]model.User, error) {
	args := r.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), nil
}

func (r *repoMock) GetByCredentials(username, password string) (model.User, error) {
	args := r.Called(username, password)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *repoMock) Insert(user *model.User) (model.User, error) {
	args := r.Called(user)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *repoMock) Update(newUser *model.User) (model.User, error) {
	args := r.Called(newUser)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (r *repoMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *UserUsecaseTestSuite) TestUserGetAll_Success() {
	suite.repoMock.On("GetAll").Return(dummyUsers, nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	users, err := UserUsecaseTest.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyUsers, users)
}

func (suite *UserUsecaseTestSuite) TestUserGetAll_Failed() {
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	users, err := UserUsecaseTest.GetAll()

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(users))
}

func (suite *UserUsecaseTestSuite) TestUserGetById_Success() {
	dummy := dummyUsers[0]
	suite.repoMock.On("GetById", dummy.Id).Return(dummy, nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.GetById(dummy.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, user)
}

func (suite *UserUsecaseTestSuite) TestUserGetById_Failed() {
	dummy := dummyUsers[0]
	suite.repoMock.On("GetById", dummy.Id).Return(model.User{}, errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, user)
}

func (suite *UserUsecaseTestSuite) TestUserGetByName_Success() {
	dummy := dummyUsers[0]
	suite.repoMock.On("GetByName", dummy.Name).Return(dummyUsers, nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	users, err := UserUsecaseTest.GetByName(dummy.Name)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyUsers, users)
}

func (suite *UserUsecaseTestSuite) TestUserGetByName_Failed() {
	dummy := dummyUsers[0]
	suite.repoMock.On("GetByName", dummy.Name).Return(nil, errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	users, err := UserUsecaseTest.GetByName(dummy.Name)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), []model.User(nil), users)
}

func (suite *UserUsecaseTestSuite) TestUserGetByCredentials_Success() {
	dummy := dummyUsers[0]
	suite.repoMock.On("GetByCredentials", dummy.Username, dummy.Password).Return(dummy, nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.GetByCredentials(dummy.Username, dummy.Password)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, user)
}

func (suite *UserUsecaseTestSuite) TestUserGetByCredentials_Failed() {
	dummy := dummyUsers[0]
	suite.repoMock.On("GetByCredentials", dummy.Username, dummy.Password).Return(model.User{}, errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.GetByCredentials(dummy.Username, dummy.Password)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, user)
}

func (suite *UserUsecaseTestSuite) TestUserInsert_Success() {
	dummy := dummyUsers[0]
	suite.repoMock.On("Insert", &dummy).Return(dummy, nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.Insert(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, user)

}

func (suite *UserUsecaseTestSuite) TestUserInsert_Failed() {
	dummy := dummyUsers[0]
	suite.repoMock.On("Insert", &dummy).Return(model.User{}, errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.Insert(&dummy)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, user)

}

func (suite *UserUsecaseTestSuite) TestUserUpdate_Success() {
	dummy := dummyUsers[0]
	suite.repoMock.On("Update", &dummy).Return(dummy, nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.Update(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, user)

}

func (suite *UserUsecaseTestSuite) TestUserUpdate_Failed() {
	dummy := dummyUsers[0]
	suite.repoMock.On("Update", &dummy).Return(model.User{}, errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	user, err := UserUsecaseTest.Update(&dummy)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, user)

}

func (suite *UserUsecaseTestSuite) TestUserDelete_Success() {
	dummy := dummyUsers[0]
	suite.repoMock.On("Delete", dummy.Id).Return(nil)

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	err := UserUsecaseTest.Delete(dummy.Id)

	assert.Nil(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestUserDelete_Failed() {
	dummy := dummyUsers[0]
	suite.repoMock.On("Delete", dummy.Id).Return(errors.New("failed"))

	UserUsecaseTest := usecase.NewUserUsecase(suite.repoMock)
	err := UserUsecaseTest.Delete(dummy.Id)

	assert.Error(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
