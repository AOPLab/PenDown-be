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
	Tag_id  int64 `gorm:"not null;uniqueIndex:compositeindex;" json:"tag_id"`
	Note_id int64 `gorm:"not null;uniqueIndex:compositeindex;" json:"note_id"`
	Tag     Tag   `gorm:"foreignKey:Tag_id;constraint:OnDelete:CASCADE;"`
	Note    Note  `gorm:"foreignKey:Note_id;constraint:OnDelete:CASCADE;"`
}
