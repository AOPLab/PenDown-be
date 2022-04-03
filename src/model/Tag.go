package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID       int64  `gorm:"primary_key;auto_increment" json:"tag_id"`
	Tag_name string `gorm:"not null" json:"tag_name"`
}

type NoteTag struct {
	gorm.Model
}
