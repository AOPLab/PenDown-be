package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID          int64  `gorm:"primary_key;auto_increment" json:"course_id"`
	School_id   int64  `gorm:"not null" json:"school_id"`
	Course_name string `gorm:"not null" json:"course_name"`
	School      School `gorm:"foreignKey:School_id;constraint:OnDelete:CASCADE;"`
}
