package controller_test

import (
	"warung-makan/model"

	"github.com/stretchr/testify/mock"
)

type MenuUsecaseMock struct {
	mock.Mock
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
