package repository

import (
	"github.com/flat35hd99/manesan/model"
	"gorm.io/gorm"
)

type ExperimentRepositoryImpl struct {
	DB *gorm.DB
}

type ExperimentRepository interface {
	Create(note *model.Experiment) error
}

func (experimentRepo ExperimentRepositoryImpl) Create(note *model.Experiment) error {
	cx := experimentRepo.DB.Create(note)
	return cx.Error
}
