package repository

import (
	"github.com/flat35hd99/manesan/model"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	DB *gorm.DB
}

type NoteRepository interface {
	Create(note *model.Experiment) error
}

func (noteRepo NoteRepositoryImpl) Create(note *model.Note) error {
	cx := noteRepo.DB.Create(note)
	return cx.Error
}
