package repository

import (
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type menuRepository struct {
	db *sqlx.DB
}

type MenuRepository interface {
	GetAllPaginated(page int, rows int) ([]model.Menu, error)
	GetAll() ([]model.Menu, error)
	GetById(id string) (model.Menu, error)
	// GetByName(name string) ([]model.Menu, error)

	Insert(menu *model.Menu) (model.Menu, error)
	Update(menu *model.Menu) (model.Menu, error)
	Delete(id string) error
}

func (p *menuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := p.db.Select(&menus, utils.MENU_GET_ALL+" order by id")
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (p *menuRepository) GetAllPaginated(page int, rows int) ([]model.Menu, error) {
	var menus []model.Menu
	limit := rows
	offset := limit * (page - 1)

	err := p.db.Select(&menus, utils.MENU_GET_ALL_PAGINATED, limit, offset)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (p *menuRepository) GetById(id string) (model.Menu, error) {
	var menu model.Menu
	err := p.db.Get(&menu, utils.MENU_GET_BY_ID, id)
	if err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func (p *menuRepository) Insert(newMenu *model.Menu) (model.Menu, error) {
	newMenu.Id = utils.GenerateId()
	_, err := p.db.NamedExec(utils.MENU_INSERT, newMenu)
	if err != nil {
		return model.Menu{}, err
	}
	menu := newMenu
	return *menu, nil
}

func (p *menuRepository) Update(newData *model.Menu) (model.Menu, error) {
	_, err := p.db.NamedExec(utils.MENU_UPDATE, newData)
	if err != nil {
		return model.Menu{}, err
	}
	return *newData, nil
}

func (p *menuRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.MENU_DELETE, id)
	return err
}

func NewMenuRepository(db *sqlx.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo
}
