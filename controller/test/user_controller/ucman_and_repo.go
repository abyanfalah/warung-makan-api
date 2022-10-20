package controller_test

import (
	"warung-makan/model"
	"warung-makan/usecase"

	"github.com/stretchr/testify/mock"
)

type repoMock struct {
	mock.Mock
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

type UsecaseManagerMock struct {
	repoMock *repoMock
	mock.Mock
}

func (um *UsecaseManagerMock) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(um.repoMock.UserRepo())
}

func (um *UsecaseManagerMock) MenuUsecase() usecase.MenuUsecase {
	return usecase.NewMenuUsecase(um.repoMock.MenuRepo())
}

func (um *UsecaseManagerMock) TransactionUsecase() usecase.TransactionUsecase {
	return usecase.NewTransactionUsecase(um.repoMock.TransactionRepo())
}
