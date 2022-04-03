package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
}

type NoteTag struct {
	gorm.Model
}
