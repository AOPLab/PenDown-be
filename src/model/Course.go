package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID                int64     `gorm:"primary_key;auto_increment" json:"course_id"`
	School_id         int64     `gorm:"not null" json:"school_id"`
	Course_no         string    `gorm:"not null" json:"course_no"`
	Course_name       string    `gorm:"not null" json:"course_name"`
	View_cnt          int64     `gorm:"default:0" json:"view_cnt"`
	Note_cnt          int64     `gorm:"default:0" json:"note_cnt"`
	Last_updated_time time.Time `gorm:"default:null" json:"last_updated_time"`
	School            School    `gorm:"foreignKey:School_id;constraint:OnDelete:CASCADE;"`
}
