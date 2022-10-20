package usecase

import (
	"warung-makan/model"
	"warung-makan/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	GetAll() ([]model.User, error)
	GetById(id string) (model.User, error)
	GetByName(name string) ([]model.User, error)
	GetByCredentials(username, password string) (model.User, error)

	Insert(user *model.User) (model.User, error)
	Update(user *model.User) (model.User, error)
	Delete(id string) error
}

func (p *userUsecase) GetAll() ([]model.User, error) {
	return p.userRepository.GetAll()

}

func (p *userUsecase) GetById(id string) (model.User, error) {
	return p.userRepository.GetById(id)
}

func (p *userUsecase) GetByName(name string) ([]model.User, error) {
	return p.userRepository.GetByName(name)
}

func (p *userUsecase) GetByCredentials(username, password string) (model.User, error) {
	return p.userRepository.GetByCredentials(username, password)
}

func (p *userUsecase) Insert(newUser *model.User) (model.User, error) {
	return p.userRepository.Insert(newUser)
}

func (p *userUsecase) Update(newUser *model.User) (model.User, error) {
	return p.userRepository.Update(newUser)
}

func (p *userUsecase) Delete(id string) error {
	return p.userRepository.Delete(id)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	usecase := new(userUsecase)
	usecase.userRepository = userRepository
	return usecase
}
