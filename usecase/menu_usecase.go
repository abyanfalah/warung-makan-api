package usecase

import (
	"warung-makan/model"
	"warung-makan/repository"
)

type menuUsecase struct {
	menuRepository repository.MenuRepository
}

type MenuUsecase interface {
	GetAll() ([]model.Menu, error)
	// GetAllPaginated(page int, rows int) ([]model.Menu, error)
	GetById(id string) (model.Menu, error)
	GetByName(name string) ([]model.Menu, error)
	Insert(menu *model.Menu) (model.Menu, error)
	Update(menu *model.Menu) (model.Menu, error)
	Delete(id string) error
}

func (p *menuUsecase) GetAll() ([]model.Menu, error) {
	return p.menuRepository.GetAll()

}

// func (p *menuUsecase) GetAllPaginated(page int, rows int) ([]model.Menu, error) {
// 	return p.menuRepository.GetAllPaginated(page, rows)
// }

func (p *menuUsecase) GetById(id string) (model.Menu, error) {
	return p.menuRepository.GetById(id)
}

func (p *menuUsecase) GetByName(name string) ([]model.Menu, error) {
	return p.menuRepository.GetByName(name)
}

func (p *menuUsecase) Insert(newMenu *model.Menu) (model.Menu, error) {
	return p.menuRepository.Insert(newMenu)
}

func (p *menuUsecase) Update(newMenu *model.Menu) (model.Menu, error) {
	return p.menuRepository.Update(newMenu)
}

func (p *menuUsecase) Delete(id string) error {
	return p.menuRepository.Delete(id)
}

func NewMenuUsecase(menuRepository repository.MenuRepository) MenuUsecase {
	usecase := new(menuUsecase)
	usecase.menuRepository = menuRepository
	return usecase
}
