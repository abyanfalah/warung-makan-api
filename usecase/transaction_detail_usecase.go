package usecase

// import (
// 	"warung-makan/model"
// 	"warung-makan/repository"
// )

// // import (
// // 	"warung-makan/model"
// // 	"warung-makan/repository"
// // )

// type transactionDetailUsecase struct {
// 	transactionDetailRepository repository.TransactionDetailRepository
// }

// type TransactionDetailUsecase interface {
// 	GetAll() ([]model.TransactionDetail, error)
// 	GetAllPaginated(page int, rows int) ([]model.TransactionDetail, error)
// 	// GetById(id string) (model.TransactionDetail, error)
// 	Insert(transactionDetail *model.TransactionDetail) (model.TransactionDetail, error)
// }

// // func (p *transactionDetailUsecase) GetAll() ([]model.TransactionDetail, error) {
// // 	return p.transactionDetailRepository.GetAll()
// // }

// // func (p *transactionDetailUsecase) GetById(id string) (model.TransactionDetail, error) {
// // 	return p.transactionDetailRepository.GetById(id)
// // }

// // func (p *transactionDetailUsecase) Insert(newTransactionDetail *model.TransactionDetail) (model.TransactionDetail, error) {
// // 	return p.transactionDetailRepository.Insert(newTransactionDetail)
// // }

// // func (p *transactionDetailUsecase) Update(newTransactionDetail *model.TransactionDetail) (model.TransactionDetail, error) {
// // 	return p.transactionDetailRepository.Update(newTransactionDetail)
// // }

// // func (p *transactionDetailUsecase) Delete(id string) error {
// // 	return p.transactionDetailRepository.Delete(id)
// // }

// // func NewTransactionDetailUsecase(transactionDetailRepository repository.TransactionDetailRepository) TransactionDetailUsecase {
// // 	usecase := new(transactionDetailUsecase)
// // 	usecase.transactionDetailRepository = transactionDetailRepository
// // 	return usecase
// // }
