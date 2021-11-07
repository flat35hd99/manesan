package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NoteRepositoryTestSuite struct {
	suite.Suite
	noteRepository NoteRepositoryImpl
	mock           sqlmock.Sqlmock
}

func (suite *NoteRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	noteRepository := NoteRepositoryImpl{}
	// sqlite is not supported sqlite.New()
	noteRepository.DB, _ = gorm.Open(sqlite.New(sqlite.Config{
		Conn: db,
	}), &gorm.Config{})
	suite.noteRepository = noteRepository
}
