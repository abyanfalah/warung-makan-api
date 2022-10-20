package repository

import (
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetById(id string) (model.User, error)
	GetByName(name string) ([]model.User, error)
	GetByCredentials(username, password string) (model.User, error)

	Insert(user *model.User) (model.User, error)
	Update(user *model.User) (model.User, error)
	Delete(id string) error
}

func (p *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := p.db.Select(&users, utils.USER_GET_ALL+" order by id")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *userRepository) GetById(id string) (model.User, error) {
	var user model.User
	err := p.db.Get(&user, utils.USER_GET_BY_ID, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (p *userRepository) GetByName(name string) ([]model.User, error) {
	var user []model.User
	err := p.db.Select(&user, utils.USER_GET_BY_NAME, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userRepository) GetByCredentials(username, password string) (model.User, error) {
	var user model.User
	err := p.db.Get(&user, utils.USER_GET_BY_CREDENTIALS, username, password)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (p *userRepository) Insert(newUser *model.User) (model.User, error) {
	_, err := p.db.NamedExec(utils.USER_INSERT, newUser)
	if err != nil {
		return model.User{}, err
	}
	user := newUser
	return *user, nil
}

func (p *userRepository) Update(newData *model.User) (model.User, error) {
	_, err := p.db.NamedExec(utils.USER_UPDATE, newData)
	if err != nil {
		return model.User{}, err
	}
	return *newData, nil
}

func (p *userRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.USER_DELETE, id)
	return err
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	repo := new(userRepository)
	repo.db = db
	return repo
}
