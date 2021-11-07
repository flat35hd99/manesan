package model

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Name        string
	Experiments []Experiment
}
