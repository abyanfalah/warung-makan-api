package repository_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"warung-makan/model"
	"warung-makan/repository"
	"warung-makan/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyMenus = []model.Menu{
	{
		Id:    "dummy id 1",
		Name:  "dummy name 1",
		Price: 8888,
		Stock: 99,
		Image: "dummy image path 1",
	},
	{
		Id:    "dummy id 2",
		Name:  "dummy name 2",
		Price: 8888,
		Stock: 99,
		Image: "dummy image path 2",
	},
}

type MenuRepositoryTestSuite struct {
	suite.Suite
	mockDb     *sql.DB
	mockSql    sqlmock.Sqlmock
	mockSqlxDb *sqlx.DB
}

func (suite *MenuRepositoryTestSuite) SetupTest() {
	db, sql, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	suite.mockDb = db
	suite.mockSql = sql
	suite.mockSqlxDb = sqlx.NewDb(suite.mockDb, "postgres")
}

func (suite *MenuRepositoryTestSuite) TearDownTest() {
	// var err error

	// err = suite.mockDb.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// err = suite.mockSqlxDb.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// Gak perlu teardown. Mungkin karena sql.DB gak dipake. Yang dipake hanya sqlx.DB
}

func (suite *MenuRepositoryTestSuite) TestGetAllMenu_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "image"})
	for _, dummy := range dummyMenus {
		rows.AddRow(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image)
	}

	suite.mockSql.ExpectQuery(utils.MENU_GET_ALL).WillReturnRows(rows)

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(actual))
	assert.Equal(suite.T(), "dummy id 1", actual[0].Id)
}

func (suite *MenuRepositoryTestSuite) TestGetAllMenu_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "image"})

	for _, dummy := range dummyMenus {
		rows.AddRow(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image)
	}

	suite.mockSql.ExpectQuery(utils.MENU_GET_ALL).WillReturnError(errors.New("failed to retrieve user list"))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.GetAll()

	assert.Nil(suite.T(), actual)
	assert.Equal(suite.T(), 0, len(actual))
	assert.Error(suite.T(), err)

}

func (suite *MenuRepositoryTestSuite) TestGetByIdMenu_Success() {
	dummy := dummyMenus[0]
	row := sqlmock.NewRows([]string{"id", "name", "price", "stock", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.MENU_GET_BY_ID)).WillReturnRows(row)

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), dummy, actual)
}

func (suite *MenuRepositoryTestSuite) TestGetByIdMenu_Failed() {
	dummy := dummyMenus[0]
	row := sqlmock.NewRows([]string{"id", "name", "price", "stock", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.MENU_GET_BY_ID)).WillReturnError(errors.New("failed to retrieve user"))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Menu{}, actual)
}

func (suite *MenuRepositoryTestSuite) TestGetByNameMenu_Success() {
	dummy := dummyMenus[0]
	row := sqlmock.NewRows([]string{"id", "name", "price", "stock", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.MENU_GET_BY_NAME)).WillReturnRows(row)

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.GetByName(dummy.Name)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), 1, len(actual))
}

func (suite *MenuRepositoryTestSuite) TestGetByNameMenu_Failed() {
	dummy := dummyMenus[0]
	row := sqlmock.NewRows([]string{"id", "name", "price", "stock", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(utils.MENU_GET_BY_NAME)).WillReturnError(errors.New("failed to retrieve user"))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.GetByName(dummy.Name)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(actual))
}

func (suite *MenuRepositoryTestSuite) TestInsertMenu_Success() {
	var dummy = dummyMenus[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.MENU_INSERT_TEST)).WithArgs(dummy.Id, dummy.Name, dummy.Price, dummy.Stock, dummy.Image).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.Insert(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, actual)
}

func (suite *MenuRepositoryTestSuite) TestInsertMenu_Failed() {
	var dummy = dummyMenus[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.MENU_INSERT_TEST)).WillReturnError(errors.New("insert failed"))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.Insert(&dummy)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Menu{}, actual)
}

func (suite *MenuRepositoryTestSuite) TestUpdateMenu_Success() {
	var dummy = dummyMenus[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.MENU_UPDATE_TEST)).WithArgs(dummy.Name, dummy.Price, dummy.Stock, dummy.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.Update(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, actual)
}

func (suite *MenuRepositoryTestSuite) TestUpdateMenu_Failed() {
	var dummy = dummyMenus[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.MENU_UPDATE_TEST)).WillReturnError(errors.New("update failed"))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	actual, err := repo.Update(&dummy)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.Menu{}, actual)
}

func (suite *MenuRepositoryTestSuite) TestDeleteMenu_Success() {
	var dummy = dummyMenus[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.MENU_DELETE)).WithArgs(dummy.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	err := repo.Delete(dummy.Id)

	assert.Nil(suite.T(), err)
}

func (suite *MenuRepositoryTestSuite) TestDeleteMenu_Failed() {
	var dummy = dummyMenus[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.MENU_DELETE)).WillReturnError(errors.New("delete failed"))

	repo := repository.NewMenuRepository(suite.mockSqlxDb)
	err := repo.Delete(dummy.Id)

	assert.NotNil(suite.T(), err)
}

func TestMenuRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MenuRepositoryTestSuite))
}
