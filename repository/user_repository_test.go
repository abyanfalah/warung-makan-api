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

var dummyUsers = []model.User{
	{
		Id:       "dummy id 1",
		Name:     "dummy name 1",
		Username: "dummy username 1",
		Image:    "dummy image path 1",
	},
	{
		Id:       "dummy id 2",
		Name:     "dummy name 2",
		Username: "dummy username 2",
		Image:    "dummy image path 2",
	},
}

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDb     *sql.DB
	mockSql    sqlmock.Sqlmock
	mockSqlxDb *sqlx.DB
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db, sql, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	suite.mockDb = db
	suite.mockSql = sql
	suite.mockSqlxDb = sqlx.NewDb(suite.mockDb, "postgres")
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
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

func (suite *UserRepositoryTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "username", "image"})
	for _, dummy := range dummyUsers {
		rows.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)
	}

	suite.mockSql.ExpectQuery(utils.USER_GET_ALL).WillReturnRows(rows)

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(actual))
	assert.Equal(suite.T(), "dummy id 1", actual[0].Id)
}

func (suite *UserRepositoryTestSuite) TestGetAll_Failed() {
	rows := sqlmock.NewRows([]string{"id failed", "name failed", "username failed", "image failed"})
	for _, dummy := range dummyUsers {
		rows.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)
	}

	suite.mockSql.ExpectQuery(utils.USER_GET_ALL).WillReturnError(errors.New("failed to retrieve user list"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetAll()

	assert.Nil(suite.T(), actual)
	assert.Error(suite.T(), err)

}

func (suite *UserRepositoryTestSuite) TestGetById_Success() {
	dummy := dummyUsers[0]
	row := sqlmock.NewRows([]string{"id", "name", "username", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)

	suite.mockSql.ExpectQuery(utils.USER_GET_BY_ID).WillReturnRows(row)

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), actual)
}

func (suite *UserRepositoryTestSuite) TestGetById_Failed() {
	dummy := dummyUsers[0]
	row := sqlmock.NewRows([]string{"id failed", "name failed", "username failed", "image failed"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)

	suite.mockSql.ExpectQuery(utils.USER_GET_BY_ID).WillReturnError(errors.New("failed to retrieve user"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

func (suite *UserRepositoryTestSuite) TestGetByName_Success() {
	dummy := dummyUsers[0]
	row := sqlmock.NewRows([]string{"id", "name", "username", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)

	suite.mockSql.ExpectQuery(utils.USER_GET_BY_NAME).WillReturnRows(row)

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), actual)
}

func (suite *UserRepositoryTestSuite) TestGetByName_Failed() {
	dummy := dummyUsers[0]
	row := sqlmock.NewRows([]string{"id failed", "name failed", "username failed", "image failed"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)

	suite.mockSql.ExpectQuery(utils.USER_GET_BY_NAME).WillReturnError(errors.New("failed to retrieve user"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

func (suite *UserRepositoryTestSuite) TestGetByCredential_Success() {
	dummy := dummyUsers[0]
	row := sqlmock.NewRows([]string{"id", "name", "username", "image"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)

	suite.mockSql.ExpectQuery(utils.USER_GET_BY_CREDENTIALS).WillReturnRows(row)

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), actual)
}

func (suite *UserRepositoryTestSuite) TestGetByCredential_Failed() {
	dummy := dummyUsers[0]
	row := sqlmock.NewRows([]string{"id failed", "name failed", "username failed", "image failed"})
	row.AddRow(dummy.Id, dummy.Name, dummy.Username, dummy.Image)

	suite.mockSql.ExpectQuery(utils.USER_GET_BY_CREDENTIALS).WillReturnError(errors.New("failed to retrieve user"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.GetById(dummy.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

func (suite *UserRepositoryTestSuite) TestInsert_Success() {
	var dummy = dummyUsers[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.USER_INSERT_TEST)).WithArgs(dummy.Id, dummy.Name, dummy.Username, dummy.Password, dummy.Image).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.Insert(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, actual)
}

func (suite *UserRepositoryTestSuite) TestInsert_Failed() {
	var dummy = dummyUsers[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.USER_INSERT_TEST)).WillReturnError(errors.New("insert failed"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.Insert(&dummy)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

func (suite *UserRepositoryTestSuite) TestUpdate_Success() {
	var dummy = dummyUsers[0]

	// suite.mockSql.ExpectExec(utils.USER_UPDATE_TEST).WithArgs(dummy.Id, dummy.Name, dummy.Username, dummy.Password, dummy.Image).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.USER_UPDATE_TEST)).WithArgs(dummy.Name, dummy.Username, dummy.Password, dummy.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.Update(&dummy)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummy, actual)
}

func (suite *UserRepositoryTestSuite) TestUpdate_Failed() {
	var dummy = dummyUsers[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.USER_UPDATE_TEST)).WillReturnError(errors.New("update failed"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	actual, err := repo.Update(&dummy)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

func (suite *UserRepositoryTestSuite) TestDelete_Success() {
	var dummy = dummyUsers[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.USER_DELETE)).WithArgs(dummy.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	err := repo.Delete(dummy.Id)

	assert.Nil(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestDelete_Failed() {
	var dummy = dummyUsers[0]

	suite.mockSql.ExpectExec(regexp.QuoteMeta(utils.USER_DELETE)).WillReturnError(errors.New("delete failed"))

	repo := repository.NewUserRepository(suite.mockSqlxDb)
	err := repo.Delete(dummy.Id)

	assert.NotNil(suite.T(), err)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
