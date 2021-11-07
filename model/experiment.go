package model

import "gorm.io/gorm"

type Experiment struct {
	gorm.Model
	Name   string
	NoteId uint
}
