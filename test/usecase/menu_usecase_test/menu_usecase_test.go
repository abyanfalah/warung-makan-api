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

var dummyMenus = []model.Menu{
	{
		Id:    "dummy id 1",
		Name:  "dummy name 1",
		Price: 113,
		Stock: 999,
		Image: "dummy image path 1",
	},
	{
		Id:    "dummy id 2",
		Name:  "dummy name 2",
		Price: 123,
		Stock: 999,
		Image: "dummy image path 2",
	},
}

type repoMock struct {
	mock.Mock
}

type MenuUsecaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (r *repoMock) GetAll() ([]model.Menu, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}

func (r *repoMock) GetById(id string) (model.Menu, error) {
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *repoMock) GetByName(name string) ([]model.Menu, error) {
	args := r.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}

func (r *repoMock) Insert(menu *model.Menu) (model.Menu, error) {
	args := r.Called(menu)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *repoMock) Update(newUser *model.Menu) (model.Menu, error) {
	args := r.Called(newUser)
	if args.Get(1) != nil {
		return model.Menu{}, args.Error(1)
	}
	return args.Get(0).(model.Menu), nil
}

func (r *repoMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *MenuUsecaseTestSuite) TestMenuGetAll_Success() {
	suite.repoMock.On("GetAll").Return(dummyMenus, nil)

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menus, err := MenuUsecaseTest.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenus, menus)
}

func (suite *MenuUsecaseTestSuite) TestMenuGetAll_Failed() {
	suite.repoMock.On("GetAll").Return(nil, errors.New("failed"))

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menus, err := MenuUsecaseTest.GetAll()

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(menus))
}

func (suite *MenuUsecaseTestSuite) TestMenuGetById_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("GetById", dummy.Id).Return(dummy, nil)

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menu, err := MenuUsecaseTest.GetById(dummy.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, menu)
}

func (suite *MenuUsecaseTestSuite) TestMenuGetById_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("GetById", dummy.Id).Return(model.Menu{}, errors.New("failed"))

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menu, err := MenuUsecaseTest.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Menu{}, menu)
}

func (suite *MenuUsecaseTestSuite) TestMenuGetByName_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("GetByName", dummy.Name).Return(dummyMenus, nil)

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menus, err := MenuUsecaseTest.GetByName(dummy.Name)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenus, menus)
}

func (suite *MenuUsecaseTestSuite) TestMenuGetByName_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("GetByName", dummy.Name).Return(nil, errors.New("failed"))

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menus, err := MenuUsecaseTest.GetByName(dummy.Name)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), []model.Menu(nil), menus)
}

func (suite *MenuUsecaseTestSuite) TestMenuInsert_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Insert", &dummy).Return(dummy, nil)

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menu, err := MenuUsecaseTest.Insert(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, menu)

}

func (suite *MenuUsecaseTestSuite) TestMenuInsert_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Insert", &dummy).Return(model.Menu{}, errors.New("failed"))

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menu, err := MenuUsecaseTest.Insert(&dummy)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Menu{}, menu)

}

func (suite *MenuUsecaseTestSuite) TestMenuUpdate_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Update", &dummy).Return(dummy, nil)

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menu, err := MenuUsecaseTest.Update(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, menu)

}

func (suite *MenuUsecaseTestSuite) TestMenuUpdate_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Update", &dummy).Return(model.Menu{}, errors.New("failed"))

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	menu, err := MenuUsecaseTest.Update(&dummy)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Menu{}, menu)

}

func (suite *MenuUsecaseTestSuite) TestMenuDelete_Success() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Delete", dummy.Id).Return(nil)

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	err := MenuUsecaseTest.Delete(dummy.Id)

	assert.Nil(suite.T(), err)
}

func (suite *MenuUsecaseTestSuite) TestMenuDelete_Failed() {
	dummy := dummyMenus[0]
	suite.repoMock.On("Delete", dummy.Id).Return(errors.New("failed"))

	MenuUsecaseTest := usecase.NewMenuUsecase(suite.repoMock)
	err := MenuUsecaseTest.Delete(dummy.Id)

	assert.Error(suite.T(), err)
}

func (suite *MenuUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestMenuUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(MenuUsecaseTestSuite))
}
